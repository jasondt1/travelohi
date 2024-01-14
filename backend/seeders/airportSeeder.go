package seeders

import (
	"backend/database"
	"backend/models"
)

func AirportSeeder() {
	airports := []models.Airport{
		{
			Name:   "Soekarno-Hatta International Airport",
			Code:   "CGK",
			CityID: 1,
		},
		{
			Name:   "Husein Sastranegara International Airport",
			Code:   "BDO",
			CityID: 2,
		},
		{
			Name:   "Juanda International Airport",
			Code:   "SUB",
			CityID: 3,
		},
		{
			Name:   "Adisutjipto International Airport",
			Code:   "JOG",
			CityID: 4,
		},
		{
			Name:   "Ngurah Rai International Airport",
			Code:   "DPS",
			CityID: 5,
		},
	}
	database.DB.Create(&airports)
}
