package model

import (
	"time"
)

type User struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Username string `json:"username" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`

	FullName   string    `json:"full_name" gorm:"not null"`
	LegalName  string    `json:"legal_name" gorm:"not null"`
	NIK        string    `json:"nik" gorm:"unique;not null"`
	BirthPlace string    `json:"birth_place" gorm:"not null"`
	BirthDate  time.Time `json:"birth_date" gorm:"type:date;not null"`
	Salary     int64     `json:"salary" gorm:"type:bigint;not null"`
	PhotoKtp   string    `json:"photo_ktp" gorm:"not null"`
	PhotoSelf  string    `json:"photo_self" gorm:"not null"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Roles        []Role        `json:"roles,omitempty" gorm:"many2many:user_roles;"`
	Limits       []Limit       `json:"limits,omitempty" gorm:"foreignKey:UserID;references:ID"`
	Transactions []Transaction `json:"transactions,omitempty" gorm:"foreignKey:UserID;references:ID"`
}

func (User) TableName() string {
	return "users"
}
