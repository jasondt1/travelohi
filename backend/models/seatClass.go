package models

import "time"

type SeatClass struct {
	ID         uint    `json:"id"`
	Name       string  `json:"name"`
	Multiplier float64 `json:"multiplier"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
