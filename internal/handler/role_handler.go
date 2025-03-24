package handler

import (
	"go-studi-kasus-kredit-plus/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRoles(c *gin.Context) { // TODO: add auth middleware
	data, err := service.GetRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}
