package main

import (
	config "go-studi-kasus-kredit-plus/configs"
	"go-studi-kasus-kredit-plus/internal/db"
	"go-studi-kasus-kredit-plus/internal/logger"
	"go-studi-kasus-kredit-plus/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	// Initialize logger
	logger.Init()

	// Load configuration
	config.LoadConfig()

	// Connect to the database
	db.Connect(config.AppConfig.DB_DSN)
	// db.SchemaMigrate()
	db.Migrate()
	db.Seed() // Seeding test data.

	// Initialize Gin router
	r := gin.Default()

	// Register routes
	routes.RegisterRoutes(r)

	// Start the server
	logrus.Infof("Starting server on port %s", config.AppConfig.ServerPort)
	r.Run(":" + config.AppConfig.ServerPort)
}
