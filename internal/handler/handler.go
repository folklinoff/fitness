package handler

import (
	"net/http"

	"github.com/folklinoff/fitness-app/internal/domain"
	middleware "github.com/folklinoff/fitness-app/internal/middleware/auth"
	"github.com/gin-gonic/gin"
	"golang.org/x/xerrors"
)

var users []domain.User
var id = 1

// Function for logging in
func Login(c *gin.Context) {
	t := c.Param("user_type")

	var user domain.User

	// Check user credentials and generate a JWT token
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Check if credentials are valid (replace this logic with real authentication)
	userInDB, err := getUser(user.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "no user in db"})
	}

	if user.Password == userInDB.Password {
		// Generate a JWT token
		token, err := middleware.GenerateToken(uint(user.ID), t)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}

func getUser(username string) (*domain.User, error) {
	for _, user := range users {
		if user.Username == username {
			return &user, nil
		}
	}
	return nil, xerrors.Errorf("not found")
}

// Function for registering a new user (for demonstration purposes)
func Register(c *gin.Context) {
	var user domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Remember to securely hash passwords before storing them
	user.ID = id // Just for demonstration purposes
	id++
	users = append(users, user)
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
