package seeders

import (
	"backend/database"
	"backend/models"
	"time"
)

func FlightTransactionSeeder() {
	layout := "2006-01-02 15:04:05"
	date, _ := time.Parse(layout, "2021-01-01 08:00:00")

	trans := []models.FlightTransaction{
		{
			Date:          date,
			FlightPromoID: 1,
			PaymentID:     1,
			UserID:        1,
		},
	}
	database.DB.Create(&trans)
}
