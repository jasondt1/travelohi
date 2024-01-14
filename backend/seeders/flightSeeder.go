package seeders

import (
	"backend/controllers"
	"backend/database"
	"backend/models"
	"time"
)

func FlightSeeder() {
	layout := "2006-01-02 15:04:05"
	departureTime, _ := time.Parse(layout, "2021-01-01 08:00:00")
	arrivalTime, _ := time.Parse(layout, "2021-01-01 09:00:00")

	flights := []models.Flight{
		{
			AirplaneID:           1,
			OriginAirportID:      1,
			DestinationAirportID: 2,
			DepartureTime:        departureTime,
			ArrivalTime:          arrivalTime,
			Price:                1000000,
		},
		{
			AirplaneID:           2,
			OriginAirportID:      1,
			DestinationAirportID: 3,
			DepartureTime:        departureTime,
			ArrivalTime:          arrivalTime,
			Price:                2000000,
		},
		{
			AirplaneID:           3,
			OriginAirportID:      1,
			DestinationAirportID: 4,
			DepartureTime:        departureTime,
			ArrivalTime:          arrivalTime,
			Price:                3000000,
		},
	}
	database.DB.Create(&flights)
	for i := 1; i <= 3; i++ {
		controllers.CreateSeat(uint(i))
	}
}
