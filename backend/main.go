package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	var input struct {
		Email string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}

	user := User{
		Email: input.Email,
		PasswordHash: string(hashedPassword),
	}

	if result := db.Create(&user); result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(201, gin.H{"message": "User created successfully"})
}

func login(c *gin.Context) {
	// Implement user login logic
}

func subscribe(c *gin.Context) {
	// Implement subscription logic using Paddle
}