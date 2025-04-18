package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"Role"`
	CreatedAt time.Time `json:"created_at"`
}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Role string `json:"Role"`
	jwt.RegisteredClaims
}

// var Roles := {}
var Roles = []string{"client", "moderator"}
