package controllers

import (
	"backend/database"
	"backend/models"
)

func CreateSeat(id uint) {
	query := `
    SELECT *
    FROM flights
    JOIN airplanes ON airplanes.id = flights.airplane_id
    WHERE flights.id = ?
    `
	var result struct {
		ID       uint
		Capacity int
	}

	database.DB.Raw(query, id).Scan(&result)

	var seats []models.Seat

	for i := 1; i <= result.Capacity; i++ {
		var SeatClassID uint
		if i <= 10 {
			SeatClassID = 1
		} else if i <= 30 {
			SeatClassID = 2
		} else {
			SeatClassID = 3
		}
		seat := models.Seat{
			SeatNumber:  i,
			IsAvailable: true,
			FlightID:    id,
			SeatClassID: SeatClassID,
		}
		seats = append(seats, seat)
	}
	database.DB.Create(&seats)
}
