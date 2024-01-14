package models

import "time"

type Seat struct {
	ID          uint `json:"id"`
	SeatNumber  int  `json:"seatNumber"`
	IsAvailable bool `json:"isAvailable"`
	SeatClassID uint `json:"seatClassID"`
	FlightID    uint `json:"flightID"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	SeatClass SeatClass `gorm:"foreignKey:SeatClassID"`
	Flight    Flight    `gorm:"foreignKey:FlightID"`
}
