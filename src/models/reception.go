package models

import (
	"time"

	"github.com/google/uuid"
)

type Reception struct {
	ID        uuid.UUID `json:"id"`
	PvzID     uuid.UUID `json:"pvzid"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

var Statuses = []string{"in_progress", "close"}
