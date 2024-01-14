package seeders

import (
	"backend/database"
	"backend/models"
)

func SecurityQuestionSeeder() {
	securityQuestions := []models.SecurityQuestion{
		{Question: "What is your favorite childhood pet's name?"},
		{Question: "In which city where you born?"},
		{Question: "What is the name of your favorite book or movie?"},
		{Question: "What is the name of the elementary school you attended?"},
		{Question: "What is the model of your first car?"},
	}
	database.DB.Create(&securityQuestions)
}
