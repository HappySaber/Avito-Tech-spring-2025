package controllers

import (
	"PVZ/src/database"
	"PVZ/src/models"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreatePVZ(c *gin.Context) {
	var pvz models.Pvz

	if err := c.ShouldBindBodyWithJSON(&pvz); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	flag := false

	for i := range models.Cities {
		if models.Cities[i] == pvz.City {
			flag = true
		}
	}
	if !flag {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Now we can't create PVZ in this city"})
		return
	}
	pvz.ID = uuid.New().String()
	query := "INSERT INTO pvz (id, city, created_at) VALUES ($1,$2, NOW())"
	if _, err := database.DB.Exec(query, pvz.ID, pvz.City); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pvz)
}

func GetPVZDataHandler(c *gin.Context) {

	page := c.Query("page")
	limit := c.Query("limit")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	pageNum, err := strconv.Atoi(page)
	if err != nil || pageNum < 1 {
		pageNum = 1
	}
	limitNum, err := strconv.Atoi(limit)
	if err != nil || limitNum < 1 {
		limitNum = 10
	}

	var startTime, endTime time.Time
	if startDate != "" {
		startTime, err = time.Parse("2006-01-02", startDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format", "details": err.Error()})
			return
		}
	} else {
		startTime = time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	}

	if endDate != "" {
		endTime, err = time.Parse("2006-01-02", endDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date format", "details": err.Error()})
			return
		}
		endTime = endTime.Add(24 * time.Hour)
	} else {
		endTime = time.Now()
	}

	pvzsWithReceptions, total, err := getPVZsWithReceptionsAndProducts(startTime, endTime, pageNum, limitNum)
	if err != nil {
		log.Printf("Error in getPVZsWithReceptionsAndProducts: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error retrieving PVZ data",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"page":  pageNum,
		"limit": limitNum,
		"data":  pvzsWithReceptions,
	})
}

func getPVZsWithReceptionsAndProducts(startDate, endDate time.Time, page, limit int) ([]models.PvzWithReceptions, int, error) {
	pvzs, total, err := getPaginatedPVZs(startDate, endDate, page, limit)
	if err != nil {
		return nil, 0, fmt.Errorf("getPaginatedPVZs error: %v", err)
	}

	var result []models.PvzWithReceptions
	for _, pvz := range pvzs {
		receptions, err := getReceptionsWithProducts(pvz.ID, startDate, endDate)
		if err != nil {
			return nil, 0, fmt.Errorf("getReceptionsWithProducts for pvz %s error: %v", pvz.ID, err)
		}

		result = append(result, models.PvzWithReceptions{
			Pvz:        pvz,
			Receptions: receptions,
		})
	}

	return result, total, nil
}

func getPaginatedPVZs(startDate, endDate time.Time, page, limit int) ([]models.Pvz, int, error) {
	var pvzs []models.Pvz
	var total int

	countQuery := `SELECT COUNT(DISTINCT p.id) FROM pvz p
                   JOIN receptions r ON p.id = r.pvz_id 
                   WHERE r.created_at BETWEEN $1 AND $2`

	err := database.DB.QueryRow(countQuery, startDate, endDate).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("count query error: %v", err)
	}

	query := `SELECT DISTINCT p.id, p.city, p.created_at FROM pvz p
              JOIN receptions r ON p.id = r.pvz_id 
              WHERE r.created_at BETWEEN $1 AND $2 
              ORDER BY p.id
              LIMIT $3 OFFSET $4`

	rows, err := database.DB.Query(query, startDate, endDate, limit, (page-1)*limit)
	if err != nil {
		return nil, 0, fmt.Errorf("pvz query error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var pvz models.Pvz
		if err := rows.Scan(&pvz.ID, &pvz.City, &pvz.CreatedAt); err != nil {
			return nil, 0, fmt.Errorf("scan pvz error: %v", err)
		}
		pvzs = append(pvzs, pvz)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("rows error: %v", err)
	}

	return pvzs, total, nil
}

func getReceptionsWithProducts(pvzID string, startDate, endDate time.Time) ([]models.ReceptionInfo, error) {
	var receptions []models.ReceptionInfo

	query := `SELECT id, pvz_id, status, created_at FROM receptions 
              WHERE pvz_id = $1 AND created_at BETWEEN $2 AND $3 
              ORDER BY created_at`

	rows, err := database.DB.Query(query, pvzID, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("receptions query error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var reception models.Reception
		if err := rows.Scan(&reception.ID, &reception.PvzID, &reception.Status, &reception.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan reception error: %v", err)
		}

		products, err := getProductsForReception(reception.ID)
		if err != nil {
			return nil, fmt.Errorf("getProductsForReception error: %v", err)
		}

		receptions = append(receptions, models.ReceptionInfo{
			Reception: reception,
			Products:  products,
		})
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}

	return receptions, nil
}

func getProductsForReception(receptionID string) ([]models.Product, error) {
	var products []models.Product

	query := `SELECT id, reception_id, created_at, type FROM products 
              WHERE reception_id = $1 
              ORDER BY created_at`

	rows, err := database.DB.Query(query, receptionID)
	if err != nil {
		return nil, fmt.Errorf("products query error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.ReceptionID, &product.CreatedAt, &product.Type); err != nil {
			return nil, fmt.Errorf("scan product error: %v", err)
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}

	return products, nil
}
