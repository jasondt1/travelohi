package seeders

import (
	"backend/database"
	"backend/models"
)

func FlightPromoSeeder() {
	promos := []models.FlightPromo{
		{
			Code:        "PROMO1",
			Discount:    10,
			Image:       "https://www.pngitem.com/pimgs/m/31-312379_airplane-png-transparent-png.png",
			Description: "Promo 1",
		},
		{
			Code:        "PROMO2",
			Discount:    20,
			Image:       "https://www.pngitem.com/pimgs/m/31-312379_airplane-png-transparent-png.png",
			Description: "Promo 2",
		},
	}
	database.DB.Create(&promos)
}
