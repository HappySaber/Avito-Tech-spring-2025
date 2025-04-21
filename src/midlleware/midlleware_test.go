package midlleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestIsAuthorized_NoCookie(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodGet, "/", nil)

	IsAuthorized()(c)

	assert.True(t, c.IsAborted())
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.JSONEq(t, `{"error": "Couldn't get cookie 'token'"}`, w.Body.String())
}

func TestIsModerator_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("role", "moderator")

	IsModerator()(c)

	assert.False(t, c.IsAborted())
}

func TestIsModerator_AccessDenied(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("role", "PVZemployee")

	IsModerator()(c)

	assert.True(t, c.IsAborted())
	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.JSONEq(t, `{"error": "Access denied, only moderator can do this"}`, w.Body.String())
}

func TestIsPVZemployee_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("role", "PVZemployee")

	IsPVZemployee()(c)

	assert.False(t, c.IsAborted())
}

func TestIsPVZemployee_AccessDenied(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("role", "moderator")

	IsPVZemployee()(c)

	assert.True(t, c.IsAborted())
	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.JSONEq(t, `{"error": "Access denied, only PVZemployee can do this"}`, w.Body.String())
}
