package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func Init() {
	Log = logrus.New()

	// Set log level
	Log.SetLevel(logrus.InfoLevel)

	// Set log format to JSON (optional)
	Log.SetFormatter(&logrus.JSONFormatter{})

	// Optional: log to a file
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Log.SetOutput(file)
	} else {
		Log.SetOutput(os.Stdout)
	}

	// Log to standard output
	Log.SetOutput(os.Stdout)
}
