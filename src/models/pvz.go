package models

import (
	"time"
)

type Pvz struct {
	ID        string    `json:"id"`
	City      string    `json:"city"`
	CreatedAt time.Time `json:"created_at"`
}

var Cities = []string{"Saint-Peterburg", "Moscow", "Kazan"}
