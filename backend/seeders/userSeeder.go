package seeders

import (
	"backend/database"
	"backend/models"
	"time"
)

func UserSeeder() {
	layout := "2006-01-02 15:04:05"
	date, _ := time.Parse(layout, "2021-01-01 08:00:00")

	users := []models.User{
		{
			FirstName:          "John",
			LastName:           "Doe",
			DateOfBirth:        date,
			Email:              "john.doe@example.com",
			Password:           "hashedpassword123",
			Gender:             "male",
			Image:              "https://example.com/john_doe_picture.jpg",
			SecurityQuestionID: 2,
			SecurityAnswer:     "Blue",
			Role:               "user",
			Address:            "Jl. Janur Raya",
			IsSubscribed:       false,
			Phone:              "081234567890",
			OTP:                "",
			OTPExpiredAt:       date,
			IsBanned:           false,
			Wallet:             0,
		},
	}
	database.DB.Create(&users)
}
