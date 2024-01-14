package models

import "time"

type Transit struct {
	ID        uint      `json:"id"`
	AirportID uint      `json:"airportID"`
	FlightID  uint      `json:"flightID"`
	Duration  time.Time `json:"duration"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Airport Airport `gorm:"foreignKey:AirportID"`
	Flight  Flight  `gorm:"foreignKey:FlightID"`
}
