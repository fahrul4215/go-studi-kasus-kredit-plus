package request

type CreatePayment struct {
	TransactionID uint  `json:"transaction_id" binding:"required"`
	AmountPaid    int64 `json:"amount_paid" binding:"required"`
}
