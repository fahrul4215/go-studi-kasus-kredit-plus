package handler

import (
	"go-studi-kasus-kredit-plus/internal/auth"
	"go-studi-kasus-kredit-plus/internal/db"
	"go-studi-kasus-kredit-plus/internal/db/model"
	"go-studi-kasus-kredit-plus/internal/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Username string       `json:"username" binding:"required"`
	Password string       `json:"password" binding:"required"`
	Roles    []model.Role `json:"roles" binding:"required"`
}

// RegisterUser handles user registration.
func RegisterUser(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Roles = []model.Role{
		{Name: "user"},
	}
	// Fetch the role by name
	var roles []model.Role
	for _, roleName := range req.Roles {
		var role model.Role
		if err := db.DB.Where("name = ?", roleName).First(&role).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid role name"})
			return
		}
	}

	// Hash the password
	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	// Create a new user record
	user := model.User{
		Username: req.Username,
		Password: hashedPassword,
		Roles:    roles,
	}

	if err := db.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user registered successfully"})
}

func GetUsers(c *gin.Context) {
	var users []model.User
	if err := db.DB.Preload("Limits").Preload("Transactions").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.Success{
		Message: "Success",
		Data:    users,
	})
}
