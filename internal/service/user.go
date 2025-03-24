package service

import (
	"go-studi-kasus-kredit-plus/internal/db"
	"go-studi-kasus-kredit-plus/internal/db/model"
)

func AssignRoleToUser(userID uint, roleName string) error {
	// Find the role by name
	var role model.Role
	if err := db.DB.Where("name = ?", roleName).First(&role).Error; err != nil {
		return err
	}

	// Assign the role to the user
	userRole := model.UserRole{
		UserID: userID,
		RoleID: role.ID,
	}
	return db.DB.Create(&userRole).Error
}

func GetUserRoles(userID uint) ([]model.Role, error) {
	var roles []model.Role
	err := db.DB.Model(&model.User{}).
		Where("users.id = ?", userID).
		Joins("JOIN user_roles ON user_roles.user_id = users.id").
		Joins("JOIN roles ON roles.id = user_roles.role_id").
		Find(&roles).Error

	return roles, err
}
