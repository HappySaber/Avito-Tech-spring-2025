package controllers

import (
	"PVZ/src/database"
	"PVZ/src/models"
	"PVZ/src/utils"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var jwtKey = []byte("jwtSecret")

// {
//     "email":"example@mail.ru",
//     "password":"securePassword",
//     "role": "client"
// }

// {
//     "email":"example@yandex.ru",
//     "password":"securePassword",
//     "role": "moderator"
// }

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
	var user models.UserRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User

	query := "SELECT email, password, role FROM users WHERE email = $1"

	err := database.DB.QueryRow(query, user.Email).Scan(&existingUser.Email, &existingUser.Password, &existingUser.Role)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User doesn't exist"})
		return
	}

	//deleting spaces from passwords, if they some way managed to be
	user.Password = strings.TrimSpace(user.Password)
	existingUser.Password = strings.TrimSpace(existingUser.Password)

	errHash := utils.CompareHashPassword(user.Password, existingUser.Password)

	if !errHash {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}

	expirationTime := time.Now().Add(30 * time.Minute)

	claims := &models.Claims{
		Role: existingUser.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   existingUser.ID,
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
	c.JSON(http.StatusOK, gin.H{"succes": "user logged in"})
}

func Signup(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !govalidator.IsEmail(user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email address"})
		return
	}

	flag := false
	for i := range models.Roles {
		if models.Roles[i] == user.Role {
			flag = true
		}
	}

	if !flag {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Such role doesn't exist"})
		return
	}

	query := `SELECT id, email, password, role, created_at FROM users WHERE email = $1`
	rows, err := database.DB.Query(query, user.Email)

	if err != nil {
		log.Println("Error executing query:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	defer rows.Close()

	if rows.Next() {
		var id uuid.UUID
		var username, email, password, createdAt string
		if err := rows.Scan(&id, &username, &email, &password, &createdAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning user data"})
			return
		}
		if id != uuid.Nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
			return
		}
	}

	var errHash error

	user.Password, errHash = utils.GenerateHashPassword(user.Password)

	if errHash != nil {
		c.JSON(500, gin.H{"error": "could not generate password hash"})
		return
	}

	query = "INSERT INTO users (email, password, role, created_at) VALUES ($1, $2, $3, NOW())"
	if _, err := database.DB.Exec(query, user.Email, user.Password, user.Role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "user created"})
}

func Home(c *gin.Context) {

}

func Logout(c *gin.Context) {

}
