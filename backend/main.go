package main

import (
	"backend/controllers"
	"backend/database"
	"backend/middleware"
	"backend/seeders"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func refreshDB() {
	database.Migrate()
	seeders.Seed()
}

func main() {
	database.Connect()
	// refreshDB()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           12 * 3600,
	}))

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.POST("/api/check-email", controllers.CheckEmail)
	r.GET("/api/get-security-questions", controllers.FetchSecurityQuestions)

	r.Run(":8080")
}
