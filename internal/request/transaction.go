package request

import "go-studi-kasus-kredit-plus/internal/pkg/pagination"

type GetTransaction struct {
	pagination.Pages

	ContractNumber    string `form:"contract_number" json:"contract_number,omitempty"`
	AssetName         string `form:"asset_name" json:"asset_name,omitempty"`
	OTR               int64  `form:"otr" json:"otr,omitempty"`
	AdminFee          int64  `form:"admin_fee" json:"admin_fee,omitempty"`
	InstallmentAmount int64  `form:"installment_amount" json:"installment_amount,omitempty"`
	InterestAmount    int64  `form:"interest_amount" json:"interest_amount,omitempty"`
	UserID            uint   `form:"user_id" json:"user_id,omitempty"`
	CreatedAt         string `form:"created_at" json:"created_at,omitempty"`
}

type CreateTransaction struct {
	ContractNumber    string `json:"contract_number" binding:"required"`
	AssetName         string `json:"asset_name" binding:"required"`
	OTR               int64  `json:"otr" binding:"required"`
	AdminFee          int64  `json:"admin_fee" binding:"required"`
	InstallmentAmount int64  `json:"installment_amount" binding:"required"`
	InterestAmount    int64  `json:"interest_amount" binding:"required"`
	UserID            uint   `json:"user_id" binding:"required"`

	Tenor int `json:"tenor" binding:"required"`
}
