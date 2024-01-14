package models

import "time"

type HotelPromo struct {
	ID          uint   `json:"id"`
	Code        string `json:"code"`
	Discount    int    `json:"discount"`
	Image       string `json:"image"`
	Description string `json:"description"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
