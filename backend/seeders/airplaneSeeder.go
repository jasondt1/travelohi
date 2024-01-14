package seeders

import (
	"backend/database"
	"backend/models"
)

func AirplaneSeeder() {
	airplanes := []models.Airplane{
		{
			Code:      "GA-001",
			Model:     "Airbus A330-300",
			Capacity:  222,
			AirlineID: 1,
		},
		{
			Code:      "QZ-001",
			Model:     "Airbus A320-200",
			Capacity:  180,
			AirlineID: 2,
		},
		{
			Code:      "CX-001",
			Model:     "Boeing 777-300ER",
			Capacity:  368,
			AirlineID: 3,
		},
	}
	database.DB.Create(&airplanes)
}
