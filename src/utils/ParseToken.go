package utils

import (
	"PVZ/src/models"
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

func ParseToken(tokenString string) (*models.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("sosal?da!"), nil
	})

	if err != nil {
		log.Printf("Error: %s", err)
		return nil, err
	}

	claims, ok := token.Claims.(*models.Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claims or token is not valid")
	}

	return claims, nil
}
