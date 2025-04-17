package models

import (
	"time"

	"github.com/google/uuid"
)

type Pvz struct {
	ID        uuid.UUID `json:"id"`
	City      string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}
