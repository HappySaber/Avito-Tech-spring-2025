package controllers

import (
	"PVZ/src/database"
	"PVZ/src/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreatePVZ_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	database.DB = db

	mock.ExpectExec("INSERT INTO pvz").
		WithArgs(sqlmock.AnyArg(), "Moscow").
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Настройка Gin
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	pvz := models.Pvz{City: "Moscow"}
	jsonValue, _ := json.Marshal(pvz)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonValue))
	c.Request.Header.Set("Content-Type", "application/json")

	CreatePVZ(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.Pvz
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "Moscow", response.City)
	assert.NotEmpty(t, response.ID)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreatePVZ_InvalidCity(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	pvz := models.Pvz{City: "InvalidCity"}
	jsonValue, _ := json.Marshal(pvz)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonValue))
	c.Request.Header.Set("Content-Type", "application/json")

	CreatePVZ(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, `{"error":"Now we can't create PVZ in this city"}`, w.Body.String())
}

func TestCreatePVZ_BindError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBufferString("invalid json"))
	c.Request.Header.Set("Content-Type", "application/json")

	CreatePVZ(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
