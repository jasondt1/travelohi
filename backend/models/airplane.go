package models

import "time"

type Airplane struct {
	ID        uint   `json:"id"`
	Code      string `json:"code"`
	Model     string `json:"model"`
	Capacity  int    `json:"capacity"`
	AirlineID uint   `json:"airlineId"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Airline Airline `gorm:"foreignKey:AirlineID"`
}
