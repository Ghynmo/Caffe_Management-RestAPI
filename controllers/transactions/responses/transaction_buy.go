package responses

import (
	"miniProject/business/transactions"
)

type TransactionBuy struct {
	UserID     uint `json:"user_id"`
	TotalPrice int  `json:"total_price"`
	Status     bool `json:"status"`
}

func CreateFromDomain(domain transactions.Domain) TransactionBuy {
	return TransactionBuy{
		UserID:     domain.UserID,
		TotalPrice: domain.Payment,
		Status:     domain.Status,
	}
}
