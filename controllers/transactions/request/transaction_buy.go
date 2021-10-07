package request

import "miniProject/business/transactions"

type TransactionBuy struct {
	UserID   uint   `json:"user_id"`
	Menu     string `json:"menu"`
	Quantity int    `json:"quantity"`
}

func (tr *TransactionBuy) ToDomain() transactions.Domain {
	return transactions.Domain{
		UserID:   tr.UserID,
		Menu:     tr.Menu,
		Quantity: tr.Quantity,
	}
}
