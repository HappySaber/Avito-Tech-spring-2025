package models

import (
	"time"
)

type Product struct {
	ID          string    `json:"id"`
	ReceptionID string    `json:"receptionid"`
	CreatedAt   time.Time `json:"created_at"`
	Type        string    `json:"type"`
}

var Products = []Product{}

// электроника, одежда, обувь
var Types = []string{"electronics", "clothes", "shoes"}
