package models

import "time"

type User struct {
	ID                 uint      `json:"id"`
	FirstName          string    `json:"firstName"`
	LastName           string    `json:"lastName"`
	DateOfBirth        time.Time `json:"dateOfBirth"`
	Email              string    `json:"email"`
	Password           string    `json:"password"`
	Gender             string    `json:"gender"`
	Image              string    `json:"image"`
	SecurityQuestionID uint      `json:"securityQuestionID"`
	SecurityAnswer     string    `json:"securityAnswer"`
	Role               string    `json:"role"`
	Address            string    `json:"address"`
	Phone              string    `json:"phone"`
	IsSubscribed       bool      `json:"isSubscribed"`
	OTP                string    `json:"otp"`
	IsBanned           bool      `json:"isBanned"`
	Wallet             float64   `json:"wallet"`

	OTPExpiredAt time.Time `json:"otpExpiredAt"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`

	SecurityQuestion SecurityQuestion `gorm:"foreignKey:SecurityQuestionID"`
}
