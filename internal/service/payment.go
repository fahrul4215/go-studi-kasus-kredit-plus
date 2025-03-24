package service

import (
	"errors"
	"fmt"
	"go-studi-kasus-kredit-plus/internal/db"
	"go-studi-kasus-kredit-plus/internal/db/model"
	"go-studi-kasus-kredit-plus/internal/request"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreatePayment(req request.CreatePayment) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		var transaction model.Transaction

		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&transaction, "id = ?", req.TransactionID).Error; err != nil {
			return err
		}

		if transaction.ID == 0 {
			return errors.New("transaction not found")
		}

		if transaction.Status == "PAID" {
			return errors.New("transaction already paid")
		}

		if transaction.InstallmentAmount < req.AmountPaid {
			return errors.New("amount paid exceeds the installment amount")
		}

		payment := model.Payment{
			TransactionID: transaction.ID,
			AmountPaid:    req.AmountPaid,
			PaymentDate:   time.Now(),
		}
		if err := tx.Create(&payment).Error; err != nil {
			return err
		}

		transaction.InstallmentAmount -= req.AmountPaid
		if err := tx.Save(&transaction).Error; err != nil {
			return err
		}

		if transaction.InstallmentAmount == 0 {
			transaction.Status = "PAID"
			if err := tx.Save(&transaction).Error; err != nil {
				return err
			}

			// update limit amount
			var limit model.Limit
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
				First(&limit, "user_id = ? AND tenor = ?", transaction.UserID, transaction.Tenor).Error; err != nil {
				return err
			}
			limit.Amount = transaction.LimitAmount
			if err := tx.Save(&limit).Error; err != nil {
				return err
			}

			// Create log
			log := model.Logs{
				ServiceName: "CreatePayment",
				LogMessage:  fmt.Sprintf("Payment for transaction %s completed", transaction.ContractNumber),
			}
			if err := tx.Create(&log).Error; err != nil {
				return err
			}
		}

		// Create log
		log := model.Logs{
			ServiceName: "CreatePayment",
			LogMessage:  fmt.Sprintf("Payment for transaction %s", transaction.ContractNumber),
		}
		if err := tx.Create(&log).Error; err != nil {
			return err
		}

		return nil
	})
}
