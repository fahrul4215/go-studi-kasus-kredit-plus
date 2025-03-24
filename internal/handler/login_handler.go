package handler

import (
	"go-studi-kasus-kredit-plus/internal/auth"
	"go-studi-kasus-kredit-plus/internal/db"
	"go-studi-kasus-kredit-plus/internal/db/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// LoginRequest represents the structure of the login request payload.
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse represents the structure of the login response payload.
type LoginResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

// Login handles user login and JWT generation.
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fetch user and role
	var user model.User
	if err := db.DB.Preload("Roles").Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check password
	if err := auth.CheckPassword(user.Password, req.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}

	// Generate JWT token with role
	token, err := auth.GenerateTokenWithRole(int64(user.ID), user.Roles)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{Token: token, ExpiresAt: time.Now().Add(24 * time.Hour)})
}
