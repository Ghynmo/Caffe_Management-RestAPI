package request

import "miniProject/business/transaction_details"

type TransactionDetailCreate struct {
	TransactionID int `json:"transaction_id"`
	MenuID        int `json:"menu_id"`
	Quantity      int `json:"quantity"`
}

func (transaction_detail *TransactionDetailCreate) ToDomain() transaction_details.Domain {
	return transaction_details.Domain{
		TransactionID: transaction_detail.TransactionID,
		MenuID:        transaction_detail.MenuID,
		Quantity:      transaction_detail.Quantity,
	}
}

func (tr *TransactionDetailCreate) CreateToDomain() transaction_details.Domain {
	return transaction_details.Domain{
		TransactionID: tr.TransactionID,
		MenuID:        tr.MenuID,
		Quantity:      tr.Quantity,
	}
}

func CreateToListDomain(data []TransactionDetailCreate) (result []transaction_details.Domain) {
	result = []transaction_details.Domain{}
	for _, val := range data {
		result = append(result, val.CreateToDomain())
	}
	return result
}
