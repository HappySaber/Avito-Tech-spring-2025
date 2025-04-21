package controllers

import (
	"PVZ/src/database"
	"PVZ/src/models"
	"bytes"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetActiveReceptionIDByPVZ_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected", err)
	}
	defer db.Close()
	database.DB = db

	rows := sqlmock.NewRows([]string{"id"}).AddRow("rec123")
	mock.ExpectQuery(`SELECT id FROM receptions WHERE pvz_id = \$1`).
		WithArgs("pvz123").
		WillReturnRows(rows)

	id, err := getActiveReceptionIDByPVZ("pvz123")

	assert.NoError(t, err)
	assert.Equal(t, "rec123", id)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetActiveReceptionIDByPVZ_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected", err)
	}
	defer db.Close()
	database.DB = db

	mock.ExpectQuery(`SELECT id FROM receptions WHERE pvz_id = \$1`).
		WithArgs("pvz123").
		WillReturnError(sql.ErrNoRows)

	id, err := getActiveReceptionIDByPVZ("pvz123")

	assert.Error(t, err)
	assert.Equal(t, "", id)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSaveProduct_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected", err)
	}
	defer db.Close()
	database.DB = db

	mock.ExpectExec(`INSERT INTO products`).
		WithArgs("rec123", "electronics").
		WillReturnResult(sqlmock.NewResult(1, 1))

	product := models.Product{
		ReceptionID: "rec123",
		Type:        "electronics",
	}

	err = saveProduct(product)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSaveProduct_DBError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected", err)
	}
	defer db.Close()
	database.DB = db

	mock.ExpectExec(`INSERT INTO products`).
		WithArgs("rec123", "electronics").
		WillReturnError(sql.ErrConnDone)

	product := models.Product{
		ReceptionID: "rec123",
		Type:        "electronics",
	}

	err = saveProduct(product)

	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
func TestDeleteProduct_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected", err)
	}
	defer db.Close()
	database.DB = db

	mock.ExpectExec(`DELETE FROM products`).
		WithArgs("rec123").
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = deleteProduct("rec123")

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteProduct_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected", err)
	}
	defer db.Close()
	database.DB = db

	mock.ExpectExec(`DELETE FROM products`).
		WithArgs("rec123").
		WillReturnResult(sqlmock.NewResult(0, 0))

	err = deleteProduct("rec123")

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
func TestAddProductHandler_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected", err)
	}
	defer db.Close()
	database.DB = db

	// Моки для isReceptionActive
	rows := sqlmock.NewRows([]string{"status"}).AddRow("in_progress")
	mock.ExpectQuery(`SELECT status FROM receptions`).
		WithArgs("pvz123").
		WillReturnRows(rows)

	// Моки для getActiveReceptionIDByPVZ
	rows = sqlmock.NewRows([]string{"id"}).AddRow("rec123")
	mock.ExpectQuery(`SELECT id FROM receptions WHERE pvz_id = \$1`).
		WithArgs("pvz123").
		WillReturnRows(rows)

	// Моки для saveProduct
	mock.ExpectExec(`INSERT INTO products`).
		WithArgs("rec123", "electronics").
		WillReturnResult(sqlmock.NewResult(1, 1))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "pvzid", Value: "pvz123"}}
	c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(`{"type":"electronics"}`))

	AddProductHandler(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestAddProductHandler_ClosedReception(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected", err)
	}
	defer db.Close()
	database.DB = db

	mock.ExpectQuery(`SELECT status FROM receptions`).
		WithArgs("pvz123").
		WillReturnError(sql.ErrNoRows)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "pvzid", Value: "pvz123"}}
	c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(`{"type":"electronics"}`))

	AddProductHandler(c)

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.JSONEq(t, `{"error":"The reception is closed"}`, w.Body.String())
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteLastProduct_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected", err)
	}
	defer db.Close()
	database.DB = db

	rows := sqlmock.NewRows([]string{"status"}).AddRow("in_progress")
	mock.ExpectQuery(`SELECT status FROM receptions`).
		WithArgs("pvz123").
		WillReturnRows(rows)

	rows = sqlmock.NewRows([]string{"id"}).AddRow("rec123")
	mock.ExpectQuery(`SELECT id FROM receptions WHERE pvz_id = \$1`).
		WithArgs("pvz123").
		WillReturnRows(rows)

	mock.ExpectExec(`DELETE FROM products`).
		WithArgs("rec123").
		WillReturnResult(sqlmock.NewResult(0, 1))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "pvzid", Value: "pvz123"}}

	DeleteLastProduct(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message":"Товар успешно удалён"}`, w.Body.String())
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteLastProduct_ClosedReception(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected", err)
	}
	defer db.Close()
	database.DB = db

	mock.ExpectQuery(`SELECT status FROM receptions`).
		WithArgs("pvz123").
		WillReturnError(sql.ErrNoRows)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "pvzid", Value: "pvz123"}}

	DeleteLastProduct(c)

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.JSONEq(t, `{"error":"The previous reception of the products was closed"}`, w.Body.String())
	assert.NoError(t, mock.ExpectationsWereMet())
}
