package controllers

import (
	"backend/database"
	"backend/models"
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var request models.User

	if c.Bind(&request) != nil {
		c.JSON(422, gin.H{"error": "Fields are empty!"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(500, gin.H{"error": "Error hashing password"})
		return
	}

	user := models.User{
		FirstName:          request.FirstName,
		LastName:           request.LastName,
		DateOfBirth:        request.DateOfBirth,
		Email:              request.Email,
		Password:           string(hash),
		Gender:             request.Gender,
		Image:              request.Image,
		SecurityQuestionID: request.SecurityQuestionID,
		SecurityAnswer:     request.SecurityAnswer,
		Role:               request.Role,
	}

	result := database.DB.Create(&user)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(200, gin.H{})
}

func Login(c *gin.Context) {
	var request models.User

	if c.Bind(&request) != nil {
		c.JSON(422, gin.H{"error": "Fields are empty!"})
		return
	}

	var user models.User
	database.DB.First(&user, "email = ?", request.Email)

	if user.ID == 0 {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte("SECRET"))

	if err != nil {
		c.JSON(500, gin.H{"error": "Error signing token"})
		return
	}

	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("Authorizatrion", tokenString, 3600*72, "", "", false, true)

	c.JSON(200, gin.H{})
}

func CheckEmail(c *gin.Context) {
	var request struct {
		Email string `json:"email"`
	}

	if c.Bind(&request) != nil {
		c.JSON(422, gin.H{"error": "Fields are empty!"})
		return
	}

	var user models.User
	database.DB.First(&user, "email = ?", request.Email)

	if user.ID != 0 {
		c.JSON(200, gin.H{"message": false})
		return
	}

	c.JSON(200, gin.H{"message": true})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(200, gin.H{
		"message": user,
	})
}
