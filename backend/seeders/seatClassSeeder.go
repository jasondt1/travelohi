package seeders

import (
	"backend/database"
	"backend/models"
)

func SeatClassSeeder() {
	classes := []models.SeatClass{
		{Name: "Economy", Multiplier: 1.0},
		{Name: "Business", Multiplier: 2.0},
		{Name: "First", Multiplier: 3.0},
	}
	database.DB.Create(&classes)
}
