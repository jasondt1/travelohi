package seeders

import (
	"backend/database"
	"backend/models"
)

func FlightTransactionDetailSeeder() {
	trans := []models.FlightTransactionDetail{
		{
			FlightTransactionID: 1,
			SeatID:              1,
			IsLuggage:           false,
		},
		{
			FlightTransactionID: 1,
			SeatID:              2,
			IsLuggage:           false,
		},
		{
			FlightTransactionID: 1,
			SeatID:              3,
			IsLuggage:           false,
		},
	}
	database.DB.Create(&trans)
}
