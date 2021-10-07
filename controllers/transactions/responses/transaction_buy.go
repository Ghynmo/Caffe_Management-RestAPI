package responses

import (
	"miniProject/business/transactions"
)

type TransactionBuy struct {
	UserID     uint   `json:"user_id"`
	Menu       string `json:"menu"`
	Quantity   int    `json:"quantity"`
	TotalPrice int    `json:"total_price"`
	Status     bool   `json:"status"`
}

func CreateFromDomain(domain transactions.Domain) TransactionBuy {
	return TransactionBuy{
		UserID:     domain.UserID,
		Menu:       domain.Menu,
		Quantity:   domain.Quantity,
		TotalPrice: domain.Payment,
		Status:     domain.Status,
	}
}
