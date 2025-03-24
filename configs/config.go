package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	DB_DSN     string
	ServerPort string
	JWTSecret  string
	OpenAccess bool
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("Error reading config file: %v", err)
	}

	AppConfig = Config{
		DB_DSN:     viper.GetString("DB_DSN"),
		ServerPort: viper.GetString("SERVER_PORT"),
		JWTSecret:  viper.GetString("JWT_SECRET"),
		OpenAccess: viper.GetBool("OPEN_ACCESS"),
	}
	logrus.Info("Configuration loaded successfully")
}
