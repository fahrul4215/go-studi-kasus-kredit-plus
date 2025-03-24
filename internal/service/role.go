package service

import (
	"go-studi-kasus-kredit-plus/internal/db"
	"go-studi-kasus-kredit-plus/internal/db/model"
)

func GetRoles() ([]model.Role, error) { // TODO: add auth middleware
	var roles []model.Role
	db.DB.Find(&roles)
	return roles, nil
}
