package request

import (
	"miniProject/business/transaction_details"
)

type TransactionBuyArr struct {
	UserID        uint                         `json:"user_id"`
	TransacDetail []transaction_details.Domain `json:"menuand_q"`
}

// func (tr *TransactionBuy) ToDomain() transactions.Domain {
// 	return transactions.Domain{
// 		UserID:   tr.UserID,
// 		Menu:     tr.Menu,
// 		Quantity: tr.Quantity,
// 	}
// }
