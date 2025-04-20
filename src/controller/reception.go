package controllers

import (
	"PVZ/src/database"
	"PVZ/src/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitiateReceivingHandler(c *gin.Context) {
	pvzId := c.Param("pvzid")
	var reception models.Reception

	if err := c.ShouldBindJSON(&reception); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incrorrect data"})
		return
	}

	// Проверка на наличие активной приёмки для данного ПВЗ
	check, err := isReceptionActive(pvzId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if check {
		c.JSON(http.StatusConflict, gin.H{"error": "The previous reception of the products was not closed"})
		return
	}

	reception.PvzID = pvzId
	reception.Status = models.Statuses[0]

	if err := saveReception(reception); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while saving data"})
		return
	}

	c.JSON(http.StatusOK, reception)
}

func isReceptionActive(pvz string) (bool, error) {
	var reception models.Reception

	query := `SELECT status FROM receptions WHERE pvz_id = $1 ORDER BY created_at DESC LIMIT 1`
	err := database.DB.QueryRow(query, pvz).Scan(&reception.Status)

	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}

	//models.Statuses[0] == "in_progress"
	if reception.Status == models.Statuses[0] {
		return true, nil
	}
	return false, nil
}

func saveReception(reception models.Reception) error {
	query := `INSERT INTO receptions (pvz_id, status, created_at) VALUES ($1, $2, NOW())`

	_, err := database.DB.Exec(query, reception.PvzID, reception.Status)
	if err != nil {
		return err
	}
	return nil
}

func CloseReception(c *gin.Context) {
	pvzId := c.Param("pvzid")
	var reception models.Reception
	// Проверка на наличие активной приёмки для данного ПВЗ
	check, err := isReceptionActive(pvzId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !check {
		c.JSON(http.StatusConflict, gin.H{"error": "No active reception in this PVZ"})
		return
	}

	reception.PvzID = pvzId
	reception.Status = models.Statuses[1] //closed

	if err := saveReception(reception); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while saving data"})
		return
	}

	c.JSON(http.StatusOK, reception)
}
