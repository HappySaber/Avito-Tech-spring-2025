package models

import (
	"time"
)

type Pvz struct {
	ID        string    `json:"id"`
	City      string    `json:"city"`
	CreatedAt time.Time `json:"created_at"`
}

type PvzWithReceptions struct {
	Pvz        Pvz             `json:"pvz"`
	Receptions []ReceptionInfo `json:"receptions"`
}

var Cities = []string{"Saint-Peterburg", "Moscow", "Kazan"}
