package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `json:"id"`
	ReceptionID uuid.UUID `json:"receptionid"`
	CreatedAt   time.Time `json:"created_at"`
	Type        string    `json:"type"`
}

var Products = []Product{}

// электроника, одежда, обувь
var Types = []string{"electronics", "clothes", "shoes"}
