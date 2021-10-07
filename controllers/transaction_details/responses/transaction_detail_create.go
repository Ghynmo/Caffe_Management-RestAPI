package responses

import "miniProject/business/transaction_details"

type TransactionDetailInsert struct {
	TransactionID uint `json:"transaction_id"`
	MenuID        uint `json:"menu_id"`
	Quantity      int  `json:"quantity"`
}

func CreateFromDomain(domain transaction_details.Domain) TransactionDetailInsert {
	return TransactionDetailInsert{
		TransactionID: domain.TransactionID,
		MenuID:        domain.MenuID,
		Quantity:      domain.Quantity,
	}
}
