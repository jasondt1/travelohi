package models

import "time"

type HotelCart struct {
	ID          uint `json:"id"`
	UserID      uint `json:"userID"`
	HotelRoomID uint `json:"hotelRoomID"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	User      User      `gorm:"foreignKey:UserID"`
	HotelRoom HotelRoom `gorm:"foreignKey:HotelRoomID"`
}
