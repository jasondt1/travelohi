package controllers

import (
	"backend/database"
	"backend/models"

	"github.com/gin-gonic/gin"
)

func FetchSecurityQuestions(c *gin.Context) {
	var questions []models.SecurityQuestion
	result := database.DB.Find(&questions)

	if result.Error != nil {
		c.JSON(200, gin.H{})
		return
	}

	c.JSON(200, gin.H{"questions": questions})
}
