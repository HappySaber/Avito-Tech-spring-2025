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

func TestIsReceptionActive_NoRows(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected", err)
	}
	defer db.Close()

	database.DB = db

	mock.ExpectQuery(`SELECT status FROM receptions WHERE pvz_id = \$1`).
		WithArgs("pvz123").
		WillReturnError(sql.ErrNoRows)

	active, err := isReceptionActive("pvz123")

	assert.NoError(t, err)
	assert.False(t, active)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestIsReceptionActive_ActiveReception(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected", err)
	}
	defer db.Close()

	database.DB = db

	rows := sqlmock.NewRows([]string{"status"}).AddRow(models.Statuses[0])
	mock.ExpectQuery(`SELECT status FROM receptions WHERE pvz_id = \$1`).
		WithArgs("pvz123").
		WillReturnRows(rows)

	active, err := isReceptionActive("pvz123")

	assert.NoError(t, err)
	assert.True(t, active)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestIsReceptionActive_DBError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected", err)
	}
	defer db.Close()

	database.DB = db

	mock.ExpectQuery(`SELECT status FROM receptions WHERE pvz_id = \$1`).
		WithArgs("pvz123").
		WillReturnError(sql.ErrConnDone)

	active, err := isReceptionActive("pvz123")

	assert.Error(t, err)
	assert.False(t, active)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSaveReception_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected", err)
	}
	defer db.Close()

	database.DB = db

	mock.ExpectExec(`INSERT INTO receptions`).
		WithArgs("pvz123", models.Statuses[0]).
		WillReturnResult(sqlmock.NewResult(1, 1))

	reception := models.Reception{
		PvzID:  "pvz123",
		Status: models.Statuses[0],
	}

	err = saveReception(reception)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSaveReception_DBError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected", err)
	}
	defer db.Close()

	database.DB = db

	mock.ExpectExec(`INSERT INTO receptions`).
		WithArgs("pvz123", models.Statuses[0]).
		WillReturnError(sql.ErrConnDone)

	reception := models.Reception{
		PvzID:  "pvz123",
		Status: models.Statuses[0],
	}

	err = saveReception(reception)

	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestInitiateReceivingHandler_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected", err)
	}
	defer db.Close()

	database.DB = db

	mock.ExpectQuery(`SELECT status FROM receptions`).
		WithArgs("pvz123").
		WillReturnError(sql.ErrNoRows)

	mock.ExpectExec(`INSERT INTO receptions`).
		WithArgs("pvz123", models.Statuses[0]).
		WillReturnResult(sqlmock.NewResult(1, 1))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "pvzid", Value: "pvz123"}}
	c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(`{}`))

	InitiateReceivingHandler(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestInitiateReceivingHandler_ActiveReceptionExists(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected", err)
	}
	defer db.Close()

	database.DB = db

	rows := sqlmock.NewRows([]string{"status"}).AddRow(models.Statuses[0])
	mock.ExpectQuery(`SELECT status FROM receptions`).
		WithArgs("pvz123").
		WillReturnRows(rows)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "pvzid", Value: "pvz123"}}
	c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(`{}`))

	InitiateReceivingHandler(c)

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.JSONEq(t, `{"error":"The previous reception of the products was not closed"}`, w.Body.String())
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCloseReception_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected", err)
	}
	defer db.Close()

	database.DB = db

	rows := sqlmock.NewRows([]string{"status"}).AddRow(models.Statuses[0])
	mock.ExpectQuery(`SELECT status FROM receptions`).
		WithArgs("pvz123").
		WillReturnRows(rows)

	mock.ExpectExec(`INSERT INTO receptions`).
		WithArgs("pvz123", models.Statuses[1]).
		WillReturnResult(sqlmock.NewResult(1, 1))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "pvzid", Value: "pvz123"}}

	CloseReception(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCloseReception_NoActiveReception(t *testing.T) {
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

	CloseReception(c)

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.JSONEq(t, `{"error":"No active reception in this PVZ"}`, w.Body.String())
	assert.NoError(t, mock.ExpectationsWereMet())
}
