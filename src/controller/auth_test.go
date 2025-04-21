package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	return router
}

func TestDummyLogin_ValidRequest(t *testing.T) {
	router := setupTestRouter()
	router.POST("/dummy-login", DummyLogin)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/dummy-login", strings.NewReader(`{"email":"test@example.com","password":"password","role":"client"}`))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"succes":"user logged in by dummyLogin"}`, w.Body.String())

	cookies := w.Result().Cookies()
	assert.NotEmpty(t, cookies)
	assert.Equal(t, "token", cookies[0].Name)
}

func TestDummyLogin_InvalidRole(t *testing.T) {
	router := setupTestRouter()
	router.POST("/dummy-login", DummyLogin)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/dummy-login", strings.NewReader(`{"email":"test@example.com","password":"password","role":"invalid"}`))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, `{"error":"Wrong role"}`, w.Body.String())
}

func TestDummyLogin_InvalidJSON(t *testing.T) {
	router := setupTestRouter()
	router.POST("/dummy-login", DummyLogin)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/dummy-login", strings.NewReader(`invalid json`))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, `{"error":"invalid character 'i' looking for beginning of value"}`, w.Body.String())
}
