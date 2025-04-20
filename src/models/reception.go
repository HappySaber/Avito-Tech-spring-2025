package models

import (
	"time"
)

type Reception struct {
	ID        string    `json:"id"`
	PvzID     string    `json:"pvzid"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type ReceptionInfo struct {
	Reception Reception
	Products  []Product
}

var Statuses = []string{"in_progress", "close"}
