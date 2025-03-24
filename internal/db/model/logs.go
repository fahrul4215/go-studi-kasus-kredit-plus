package model

import "time"

type Logs struct {
	ID uint `json:"id" gorm:"primaryKey"`

	ServiceName string `json:"service_name" gorm:"not null"`
	LogMessage  string `json:"log_message" gorm:"not null"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Logs) TableName() string {
	return "logs"
}
