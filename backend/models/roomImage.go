package models

import "time"

type RoomImage struct {
	ID          uint   `json:"id"`
	HotelRoomID uint   `json:"roomId"`
	Image       string `json:"image"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	HotelRoom HotelRoom `gorm:"foreignKey:HotelRoomID"`
}
