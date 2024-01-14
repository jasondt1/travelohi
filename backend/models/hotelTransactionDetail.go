package models

import "time"

type HotelTransactionDetail struct {
	ID                 uint      `json:"id"`
	HotelTransactionID uint      `json:"hotelTransactionId"`
	HotelRoomID        uint      `json:"hotelRoomId"`
	CheckInDate        time.Time `json:"checkInDate"`
	CheckOutDate       time.Time `json:"checkOutDate"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	HotelTransaction HotelTransaction `gorm:"foreignKey:HotelTransactionID"`
	HotelRoom        HotelRoom        `gorm:"foreignKey:HotelRoomID"`
}
