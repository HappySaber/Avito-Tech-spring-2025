package utils

import (
	"PVZ/src/models"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestCompareHashPassword_MatchingPassword(t *testing.T) {
	password := "securePassword"
	hash, err := GenerateHashPassword(password)
	if err != nil {
		t.Fatalf("Failed to generate hash: %v", err)
	}

	result := CompareHashPassword(password, string(hash))
	assert.True(t, result, "Expected passwords to match")
}

func TestCompareHashPassword_NonMatchingPassword(t *testing.T) {
	password := "securePassword"
	hash, err := GenerateHashPassword(password)
	if err != nil {
		t.Fatalf("Failed to generate hash: %v", err)
	}

	result := CompareHashPassword("wrongPassword", string(hash))
	assert.False(t, result, "Expected passwords to not match")
}

func createTestToken(role string) string {
	claims := &models.Claims{
		Role: role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte("jwtSecret"))
	return tokenString
}

func TestParseToken_ValidToken(t *testing.T) {
	tokenString := createTestToken("moderator")
	claims, err := ParseToken(tokenString)

	assert.NoError(t, err)
	assert.Equal(t, "moderator", claims.Role)
}

func TestParseToken_InvalidToken(t *testing.T) {
	invalidToken := "invalid.token.string"
	claims, err := ParseToken(invalidToken)

	assert.Error(t, err)
	assert.Nil(t, claims)
}

func TestParseToken_InvalidClaims(t *testing.T) {
	claims := &models.Claims{
		Role: "moderator",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte("wrongSecret"))

	claims, err := ParseToken(tokenString)

	assert.Error(t, err)
	assert.Nil(t, claims)
}

func TestGenerateHashPassword_ValidPassword(t *testing.T) {
	password := "mySecurePassword"
	hash, err := GenerateHashPassword(password)

	assert.NoError(t, err)
	assert.NotEmpty(t, hash)

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	assert.NoError(t, err)
}

func TestGenerateHashPassword_EmptyPassword(t *testing.T) {
	password := ""
	hash, err := GenerateHashPassword(password)

	assert.NoError(t, err)
	assert.NotEmpty(t, hash)

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	assert.NoError(t, err)
}
