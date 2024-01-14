package models

import "time"

type HotelFacility struct {
	ID          uint   `json:"id"`
	HotelID     uint   `json:"hotelId"`
	Name        string `json:"name"`
	Description string `json:"description"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Hotel Hotel `gorm:"foreignKey:HotelID"`
}
