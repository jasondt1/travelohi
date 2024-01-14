package seeders

import (
	"backend/database"
	"backend/models"
)

func CitySeeder() {
	cities := []models.City{
		{Name: "Jakarta"},
		{Name: "Bandung"},
		{Name: "Surabaya"},
		{Name: "Yogyakarta"},
		{Name: "Denpasar"},
	}
	database.DB.Create(&cities)
}
