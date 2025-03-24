package model

import (
	"time"
)

type Limit struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	UserID uint `json:"user_id" gorm:"not null"`

	Tenor  int   `json:"tenor" gorm:"type:smallint;not null"`
	Amount int64 `json:"amount" gorm:"type:bigint;not null"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	User User `json:"-" gorm:"foreignKey:UserID;references:ID"`
}

func (Limit) TableName() string {
	return "limits"
}
