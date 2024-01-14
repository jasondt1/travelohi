package seeders

import (
	"backend/database"
	"backend/models"
)

func AirlineSeeder() {
	airlines := []models.Airline{
		{
			Name:  "Garuda Indonesia",
			Image: "https://upload.wikimedia.org/wikipedia/id/f/fe/Garuda_Indonesia_Logo.svg",
		},
		{
			Name:  "AirAsia",
			Image: "https://id.wikipedia.org/wiki/AirAsia#/media/Berkas:AirAsia_New_Logo_(2020).svg",
		},
		{
			Name:  "Cathay Pacific",
			Image: "https://en.wikipedia.org/wiki/Cathay_Pacific#/media/File:Cathay_Pacific_logo.svg",
		},
	}
	database.DB.Create(&airlines)
}
