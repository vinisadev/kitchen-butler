package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string `gorm:"unique"`
	PasswordHash string
	SubscriptionStatus string
	PaddleSubscriptionID string
}

var db *gorm.DB

func main() {
	// Database connection
	dsn := os.Getenv("DATABASE_URL")
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Auto migrate the schema
	db.AutoMigrate(&User{})

	// Setup Gin router
	r := gin.Default()

	// Define routes
	r.POST("/signup", signUp)
	r.POST("/login", login)
	r.POST("/subscribe", subscribe)

	// Run the server
	r.Run(":8080")
}

func signUp(c *gin.Context) {
	// Implement user signup logic
}

func login(c *gin.Context) {
	// Implement user login logic
}

func subscribe(c *gin.Context) {
	// Implement subscription logic using Paddle
}