package models

import (
	"time"

	"github.com/google/uuid"
)

type Pvz struct {
	ID        uuid.UUID `json:"id"`
	City      string    `json:"city"`
	CreatedAt time.Time `json:"created_at"`
}

var Cities = []string{"Saint-Peterburg", "Moscow", "Kazan"}
