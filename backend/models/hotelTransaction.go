package models

import "time"

type HotelTransaction struct {
	ID           uint      `json:"id"`
	Date         time.Time `json:"date"`
	HotelPromoID uint      `json:"hotelPromoId"`
	PaymentID    uint      `json:"paymentId"`
	UserID       uint      `json:"userId"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	HotelPromo HotelPromo `gorm:"foreignKey:HotelPromoID"`
	Payment    Payment    `gorm:"foreignKey:PaymentID"`
	User       User       `gorm:"foreignKey:UserID"`
}
