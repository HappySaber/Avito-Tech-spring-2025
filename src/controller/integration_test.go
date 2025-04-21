package controllers

import (
	"PVZ/src/database"
	"PVZ/src/models"
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestPVZWorkflow(t *testing.T) {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatalf("Error while loading .env file: %v", err)
	}

	database.Init()

	cleanDatabase(t, database.DB)

	router := gin.Default()
	router.POST("/pvz", CreatePVZ)
	router.POST("/pvz/:pvzid/receptions", InitiateReceivingHandler)
	router.POST("/pvz/:pvzid/products", AddProductHandler)
	router.PUT("/pvz/:pvzid/receptions/close", CloseReception)

	var pvzID string
	t.Run("Create PVZ", func(t *testing.T) {
		pvz := models.Pvz{City: "Moscow"}
		jsonValue, _ := json.Marshal(pvz)
		t.Logf("Request JSON: %s", jsonValue)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/pvz", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response models.Pvz
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Moscow", response.City, "City should be 'Moscow'")
		assert.NotEmpty(t, response.ID, "PVZ ID should not be empty")

		pvzID = response.ID
		t.Logf("Created PVZ with ID: %s", pvzID)
	})

	var receptionID string

	assert.NoError(t, err)
	t.Run("Create Reception", func(t *testing.T) {
		receptionInput := map[string]interface{}{
			"status": "in_progress",
		}
		jsonValue, _ := json.Marshal(receptionInput)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/pvz/"+pvzID+"/receptions", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		t.Logf("Response: %s", w.Body.String())
		assert.Equal(t, http.StatusOK, w.Code)

		var response models.Reception
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, pvzID, response.PvzID)
		assert.Equal(t, "in_progress", response.Status)
		assert.NotEmpty(t, response.ID)

		receptionID = response.ID
		t.Logf("Created Reception with ID: %s", receptionID)
	})

	t.Run("Add Products", func(t *testing.T) {
		productTypes := []string{"electronics", "clothes", "shoes"}
		for _, productType := range productTypes {
			product := models.Product{
				ReceptionID: receptionID,
				Type:        productType,
			}
			jsonValue, _ := json.Marshal(product)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/pvz/"+pvzID+"/products", bytes.NewBuffer(jsonValue))
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code)

			var response models.Product
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, productType, response.Type)
		}
	})

	t.Run("Close Reception", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", "/pvz/"+pvzID+"/receptions/close", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response models.Reception
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "closed", response.Status)
	})

	t.Run("Add Product to Closed Reception", func(t *testing.T) {
		product := models.Product{
			ReceptionID: receptionID,
			Type:        "electronics",
		}
		jsonValue, _ := json.Marshal(product)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/pvz/"+pvzID+"/products", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusConflict, w.Code)
	})
}

func cleanDatabase(t *testing.T, db *sql.DB) {
	if db == nil {
		t.Fatal("Database connection is nil")
	}

	tables := []string{"products", "receptions", "pvz"}
	for _, table := range tables {
		_, err := db.Exec("TRUNCATE TABLE " + table + " CASCADE")
		if err != nil {
			t.Logf("Failed to truncate table %s: %v", table, err)
		}
	}
}

func init() {
	models.Cities = []string{"Moscow", "Saint-Petersburg", "Kazan"}
	models.Types = []string{"electronics", "clothes", "shoes"}
	models.Statuses = []string{"in_progress", "closed"}

	gin.SetMode(gin.TestMode)
}
