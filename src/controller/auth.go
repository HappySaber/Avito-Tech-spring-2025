package controllers

import (
	"PVZ/src/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("sosal?da!")

func DummyLogin(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	flag := false
	for i := range models.Roles {
		if models.Roles[i] == user.Role {
			flag = true
		}
	}

	if !flag {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong role"})
		return
	}

	expirationTime := time.Now().Add(30 * time.Minute)

	claims := &models.Claims{
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   "dummyLogin",
			ExpiresAt: &jwt.NumericDate{Time: expirationTime},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		c.JSON(500, gin.H{"error": "could not create token"})
		return
	}

	c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"succes": "user logged in by dummyLogin"})
}

func Login(c *gin.Context) {
}

func Signup(c *gin.Context) {

}

func Home(c *gin.Context) {

}

func Logout(c *gin.Context) {
}
