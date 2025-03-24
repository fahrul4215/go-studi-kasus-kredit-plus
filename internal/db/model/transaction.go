package model

import "time"

type Transaction struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	UserID uint `json:"user_id" gorm:"not null"`

	ContractNumber    string `json:"contract_number" gorm:"unique;not null"`
	AssetName         string `json:"asset_name" gorm:"not null"`
	OTR               int64  `json:"otr" gorm:"type:bigint;not null"`
	AdminFee          int64  `json:"admin_fee" gorm:"type:bigint;not null"`
	InterestAmount    int64  `json:"interest_amount" gorm:"type:bigint;not null"`
	InstallmentAmount int64  `json:"installment_amount" gorm:"type:bigint;not null"`

	LimitAmount int64 `json:"limit_amount" gorm:"type:bigint;not null"`
	Tenor       int   `json:"tenor" gorm:"type:smallint;not null"`

	Status string `json:"status" gorm:"default:'null'"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	User     User      `json:"user" gorm:"foreignKey:UserID;references:ID"`
	Payments []Payment `json:"payments" gorm:"foreignKey:TransactionID;references:ID"`
}

func (Transaction) TableName() string {
	return "transactions"
}
