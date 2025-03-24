package handler

import (
	"go-studi-kasus-kredit-plus/internal/errors"
	"go-studi-kasus-kredit-plus/internal/request"
	"go-studi-kasus-kredit-plus/internal/response"
	"go-studi-kasus-kredit-plus/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePayment(c *gin.Context) {
	var req request.CreatePayment
	if err := c.ShouldBindJSON(&req); err != nil {
		parserError := errors.ParseErrorValidation(err)
		parseErrorResponse := errors.BadRequest(parserError)
		c.JSON(parseErrorResponse.StatusCode(), parseErrorResponse)
		return
	}

	if err := service.CreatePayment(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response.Success{
		Message: "Success",
	})
}
