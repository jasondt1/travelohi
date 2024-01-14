package models

import "time"

type FlightTransaction struct {
	ID            uint      `json:"id"`
	Date          time.Time `json:"date"`
	FlightPromoID uint      `json:"flightPromoId"`
	PaymentID     uint      `json:"paymentId"`
	UserID        uint      `json:"userId"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	FlightPromo FlightPromo `gorm:"foreignKey:FlightPromoID"`
	Payment     Payment     `gorm:"foreignKey:PaymentID"`
	User        User        `gorm:"foreignKey:UserID"`
}
