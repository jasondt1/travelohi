package seeders

import (
	"backend/database"
	"backend/models"
)

func PaymentSeeder() {
	payments := []models.Payment{
		{
			Name:  "BCA",
			Image: "https://id.wikipedia.org/wiki/Bank_Central_Asia#/media/Berkas:Bank_Central_Asia.svg",
		},
		{
			Name:  "OVO",
			Image: "https://id.wikipedia.org/wiki/OVO_(pembayaran)#/media/Berkas:Logo_ovo_purple.svg",
		},
	}
	database.DB.Create(&payments)
}
