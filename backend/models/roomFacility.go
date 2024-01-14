package models

import "time"

type RoomFacility struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	HotelRoomID uint   `json:"roomId"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	HotelRoom HotelRoom `gorm:"foreignKey:HotelRoomID"`
}
