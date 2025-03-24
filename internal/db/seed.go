package db

import (
	"go-studi-kasus-kredit-plus/internal/auth"
	"go-studi-kasus-kredit-plus/internal/db/model"
	"time"

	"github.com/sirupsen/logrus"
)

func Seed() {
	roles := []model.Role{
		// {Name: "admin"},
		{Name: "user"},
	}

	for _, role := range roles {
		if err := DB.FirstOrCreate(&role, model.Role{Name: role.Name}).Error; err != nil {
			panic(err)
		}
	}
	logrus.Info("Roles seeded successfully")

	DB.Find(&roles)

	hash, _ := auth.HashPassword("budi")
	DB.Create(&model.User{
		Email:      "budi@example.com",
		Username:   "budi",
		Password:   hash,
		FullName:   "Budi",
		LegalName:  "Budi",
		NIK:        "1234567890",
		BirthPlace: "Jakarta",
		BirthDate:  time.Date(1998, 3, 3, 0, 0, 0, 0, time.UTC),
		Salary:     10000000,
		PhotoKtp:   "ktp.jpg",
		PhotoSelf:  "self.jpg",
		Roles:      roles,
		Limits: []model.Limit{
			{Tenor: 1, Amount: 100000},
			{Tenor: 2, Amount: 200000},
			{Tenor: 3, Amount: 500000},
			{Tenor: 4, Amount: 700000},
		},
	}).Association("Roles")

	hash, _ = auth.HashPassword("annisa")
	DB.Create(&model.User{
		Email:      "annisa@example.com",
		Username:   "annisa",
		Password:   hash,
		FullName:   "Annisa",
		LegalName:  "Annisa",
		NIK:        "0987654321",
		BirthPlace: "Jakarta",
		BirthDate:  time.Date(1994, 1, 27, 0, 0, 0, 0, time.UTC),
		Salary:     20000000,
		PhotoKtp:   "ktp.jpg",
		PhotoSelf:  "self.jpg",
		Roles:      roles,
		Limits: []model.Limit{
			{Tenor: 1, Amount: 1000000},
			{Tenor: 2, Amount: 1200000},
			{Tenor: 3, Amount: 1500000},
			{Tenor: 4, Amount: 2000000},
		},
	}).Association("Roles")

	logrus.Info("User seeded successfully")
}
