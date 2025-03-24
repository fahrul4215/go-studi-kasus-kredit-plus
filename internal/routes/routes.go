package routes

import (
	"go-studi-kasus-kredit-plus/internal/handler"
	"go-studi-kasus-kredit-plus/internal/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/login", func(c *gin.Context) {
		handler.Login(c)
	})

	apiV1 := r.Group("/api/v1")

	apiV1.Use(middleware.AuthMiddleware())
	{
		apiV1.Use(middleware.RoleMiddleware([]string{"user", "admin"}))
		{
			apiV1.GET("/users", handler.GetUsers)
			apiV1.GET("/transactions", handler.GetTransactions)

			apiV1.POST("/transactions", handler.CreateTransaction)
			apiV1.POST("/payments", handler.CreatePayment)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
