package models

import "time"

type FlightTransactionDetail struct {
	ID                  uint `json:"id"`
	FlightTransactionID uint `json:"flightTransactionId"`
	SeatID              uint `json:"seatID"`
	IsLuggage           bool `json:"isLuggage"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	FlightTransaction FlightTransaction `gorm:"foreignKey:FlightTransactionID"`
	Seat              Seat              `gorm:"foreignKey:SeatID"`
}
