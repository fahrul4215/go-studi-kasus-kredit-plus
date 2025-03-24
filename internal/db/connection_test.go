package db

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestConnect_Success(t *testing.T) {
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	Connect(dsn)
	assert.NotNil(t, DB)
	assert.NoError(t, DB.Error)
}

func TestConnect_Failure(t *testing.T) {
	// Invalid DSN to force a connection failure
	dsn := "invalid_dsn"
	hook := test.NewLocal(logrus.StandardLogger())
	Connect(dsn)
	assert.Nil(t, DB)
	assert.Contains(t, hook.LastEntry().Message, "Failed to connect to database")
}
