package db

import (
	"go-studi-kasus-kredit-plus/internal/db/model"

	"github.com/sirupsen/logrus"
)

func Migrate() {
	logrus.Info("Running migrations...")
	if err := DB.AutoMigrate(
		&model.Role{},
		&model.User{},
		&model.Limit{},
		&model.Transaction{},
		&model.Payment{},
		&model.Logs{},
	); err != nil {
		logrus.Fatalf("Migration failed: %v", err)
	}
	logrus.Info("Migrations completed.")
}

func SchemaMigrate() error {
	schemas := []string{"custom"}

	for _, schema := range schemas {
		createSchemaSQL := `CREATE SCHEMA IF NOT EXISTS ` + schema
		if err := DB.Exec(createSchemaSQL).Error; err != nil {
			return err
		}
	}
	return nil
}
