package models

import "time"

type FlightCart struct {
	ID     uint `json:"id"`
	UserID uint `json:"userID"`
	SeatID uint `json:"seatID"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	User User `gorm:"foreignKey:UserID"`
	Seat Seat `gorm:"foreignKey:SeatID"`
}
