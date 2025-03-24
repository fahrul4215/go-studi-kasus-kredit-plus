package model

import "time"

type Payment struct {
	ID            uint `json:"id" gorm:"primaryKey"`
	TransactionID uint `json:"transaction_id" gorm:"not null"`

	AmountPaid  int64     `json:"amount_paid" gorm:"type:bigint;not null"`
	PaymentDate time.Time `json:"payment_date" gorm:"type:date;not null"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Transaction Transaction `json:"transaction" gorm:"foreignKey:TransactionID;references:ID"`
}

func (Payment) TableName() string {
	return "payments"
}
