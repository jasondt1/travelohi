package models

import "time"

type Flight struct {
	ID                   uint      `json:"id"`
	AirplaneID           uint      `json:"airplaneID"`
	OriginAirportID      uint      `json:"originAirportID"`
	DestinationAirportID uint      `json:"destinationAirportID"`
	DepartureTime        time.Time `json:"departureTime"`
	ArrivalTime          time.Time `json:"arrivalTime"`
	Price                float64   `json:"price"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Airplane           Airplane `gorm:"foreignKey:AirplaneID"`
	OriginAirport      Airport  `gorm:"foreignKey:OriginAirportID"`
	DestinationAirport Airport  `gorm:"foreignKey:DestinationAirportID"`
}
