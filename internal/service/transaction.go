package service

import (
	"fmt"
	"go-studi-kasus-kredit-plus/internal/db"
	"go-studi-kasus-kredit-plus/internal/db/model"
	"go-studi-kasus-kredit-plus/internal/request"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetTransactions(req request.GetTransaction) (any, int64, error) {
	var transactions []model.Transaction
	var total int64

	query := db.DB.Model(&model.Transaction{}).
		Preload("User").
		Preload("Payments")

	query = buildTransactionQuery(query, req)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.
		Offset(req.Offset()).
		Limit(req.GetLimit()).
		Order(req.OrderDB()).
		Find(&transactions).Error; err != nil {
		return nil, 0, err
	}

	var data []any
	for _, transaction := range transactions {
		data = append(data, map[string]any{
			"id":                 transaction.ID,
			"contract_number":    transaction.ContractNumber,
			"asset_name":         transaction.AssetName,
			"otr":                transaction.OTR,
			"admin_fee":          transaction.AdminFee,
			"installment_amount": transaction.InstallmentAmount,
			"interest_amount":    transaction.InterestAmount,
			"created_at":         transaction.CreatedAt,
			"user":               transaction.User,
			"payments":           transaction.Payments,
		})
	}

	return data, total, nil
}
func buildTransactionQuery(query *gorm.DB, req request.GetTransaction) *gorm.DB {
	if req.ContractNumber != "" {
		query = query.Where("contract_number = ?", req.ContractNumber)
	}
	if req.AssetName != "" {
		query = query.Where("asset_name = ?", req.AssetName)
	}
	if req.OTR != 0 {
		query = query.Where("otr = ?", req.OTR)
	}
	if req.AdminFee != 0 {
		query = query.Where("admin_fee = ?", req.AdminFee)
	}
	if req.InstallmentAmount != 0 {
		query = query.Where("installment_amount = ?", req.InstallmentAmount)
	}
	if req.InterestAmount != 0 {
		query = query.Where("interest_amount = ?", req.InterestAmount)
	}
	if req.UserID != 0 {
		query = query.Where("user_id = ?", req.UserID)
	}
	if req.CreatedAt != "" {
		query = query.Where("created_at = ?", req.CreatedAt)
	}

	return query
}

func CreateTransaction(req request.CreateTransaction) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		var limit model.Limit

		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&limit, "user_id = ? AND tenor = ?", req.UserID, req.Tenor).Error; err != nil {
			return err
		}

		if req.InstallmentAmount > limit.Amount {
			return fmt.Errorf("transaction amount exceeds the available limit")
		}

		transaction := model.Transaction{
			UserID:            req.UserID,
			ContractNumber:    req.ContractNumber,
			AssetName:         req.AssetName,
			OTR:               req.OTR,
			AdminFee:          req.AdminFee,
			InterestAmount:    req.InterestAmount,
			InstallmentAmount: req.InstallmentAmount,
			LimitAmount:       limit.Amount,
			Tenor:             limit.Tenor,
		}

		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}

		limit.Amount -= req.InstallmentAmount
		if err := tx.Save(&limit).Error; err != nil {
			return err
		}

		// Create log
		log := model.Logs{
			ServiceName: "CreateTransaction",
			LogMessage:  fmt.Sprintf("Transaction %s created from user %d", transaction.ContractNumber, transaction.UserID),
		}
		if err := tx.Create(&log).Error; err != nil {
			return err
		}

		return nil
	})
}
