package models

import "time"

type HotelRoom struct {
	ID         uint   `json:"id"`
	RoomNumber string `json:"roomNumber"`
	HotelID    uint   `json:"hotelId"`
	RoomTypeID uint   `json:"roomTypeId"`
	Price      uint   `json:"price"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Hotel    Hotel    `gorm:"foreignKey:HotelID"`
	RoomType RoomType `gorm:"foreignKey:RoomTypeID"`
}
