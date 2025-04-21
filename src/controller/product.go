package controllers

import (
	"PVZ/src/database"
	"PVZ/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AddProductHandler(c *gin.Context) {
	pvzID := c.Param("pvzid")

	check, err := isReceptionActive(pvzID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !check {
		c.JSON(http.StatusConflict, gin.H{"error": "The reception is closed"})
		return
	}

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong values"})
		return
	}

	flag := false
	for i := range models.Types {
		if product.Type == models.Types[i] {
			flag = true
		}
	}

	if !flag {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong type of product"})
		return
	}
	// Получаем ID последней незакрытой приёмки для данного ПВЗ
	receptionID, err := getActiveReceptionIDByPVZ(pvzID) // Передайте ID ПВЗ
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "No active reception in this PVZ"})
		return
	}
	product.ReceptionID = receptionID

	// Сохранение товара в хранилище данных
	product.ID = uuid.New().String()
	if err := saveProduct(product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while saving the product: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

func DeleteLastProduct(c *gin.Context) {
	pvzID := c.Param("pvzid")

	check, err := isReceptionActive(pvzID)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	if !check {
		c.JSON(http.StatusConflict, gin.H{"error": "The previous reception of the products was closed"})
		return
	}

	// Получаем ID последней незакрытой приёмки для данного ПВЗ
	receptionID, err := getActiveReceptionIDByPVZ(pvzID)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "No active reception in this PVZ"})
		return
	}

	// Удаляем товар
	if err := deleteProduct(receptionID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while deleting product: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Товар успешно удалён"})
}

func getActiveReceptionIDByPVZ(pvzID string) (string, error) {
	var receptionID string
	query := `SELECT id FROM receptions WHERE pvz_id = $1 ORDER BY created_at DESC LIMIT 1`
	err := database.DB.QueryRow(query, pvzID).Scan(&receptionID)
	if err != nil {
		return "", err // Возвращаем nil, если нет активной приёмки
	}
	return receptionID, nil
}

func saveProduct(product models.Product) error {

	query := `INSERT INTO products (reception_id, type, created_at) VALUES ($1,$2, NOW())`
	_, err := database.DB.Exec(query, product.ReceptionID, product.Type)
	if err != nil {
		return err
	}
	return nil
}

func deleteProduct(receptionID string) error {
	query := "DELETE FROM products WHERE id = (SELECT id FROM products WHERE reception_id = $1 ORDER BY created_at DESC LIMIT 1)"
	if _, err := database.DB.Exec(query, receptionID); err != nil {
		return err
	}
	return nil
}
