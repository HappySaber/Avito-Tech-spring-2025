package controllers

import (
	"PVZ/src/database"
	"PVZ/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
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

	query := "INSERT INTO pvz (city, created_at) VALUES ($1, NOW())"
	if _, err := database.DB.Exec(query, pvz.City); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "PVZ created"})
}
