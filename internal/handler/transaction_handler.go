package handler

import (
	"go-studi-kasus-kredit-plus/internal/errors"
	"go-studi-kasus-kredit-plus/internal/pkg/pagination"
	"go-studi-kasus-kredit-plus/internal/request"
	"go-studi-kasus-kredit-plus/internal/response"
	"go-studi-kasus-kredit-plus/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTransactions(c *gin.Context) {
	var req request.GetTransaction
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, total, err := service.GetTransactions(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if total == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "data transactions empty"})
		return
	}

	c.JSON(http.StatusOK, response.Success{
		Message: "Success",
		Data:    data,
		Pagination: pagination.Pages{
			Page:       req.Page,
			Limit:      req.GetLimit(),
			Sort:       req.Sort,
			Order:      req.OrderDB(),
			PageCount:  req.PageCount,
			TotalCount: int(total),
			Keyword:    req.Keyword,
		},
	})
}

func CreateTransaction(c *gin.Context) {
	var req request.CreateTransaction
	if err := c.ShouldBindJSON(&req); err != nil {
		parserError := errors.ParseErrorValidation(err)
		parseErrorResponse := errors.BadRequest(parserError)
		c.JSON(parseErrorResponse.StatusCode(), parseErrorResponse)
		return
	}

	if err := service.CreateTransaction(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response.Success{
		Message: "Transaction created successfully",
	})
}
