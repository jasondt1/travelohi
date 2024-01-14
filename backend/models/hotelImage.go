package models

import "time"

type HotelImage struct {
	ID      uint   `json:"id"`
	HotelID uint   `json:"hotelId"`
	Image   string `json:"image"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Hotel Hotel `gorm:"foreignKey:HotelID"`
}
