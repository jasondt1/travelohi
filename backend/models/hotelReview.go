package models

import "time"

type HotelReview struct {
	ID      uint   `json:"id"`
	HotelID uint   `json:"hotelId"`
	UserID  uint   `json:"userId"`
	Review  string `json:"review"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Hotel Hotel `gorm:"foreignKey:HotelID"`
	User  User  `gorm:"foreignKey:UserID"`
}
