package request

import (
	"miniProject/business/transactions"
)

type TransactionBuy struct {
	UserID int `json:"user_id"`
	// TransacDetail []transaction_details.Domain `json:"menu_data"`
}

func (tr *TransactionBuy) ToDomain() transactions.Domain {
	return transactions.Domain{
		UserID: tr.UserID,
		// TransactionDetails: tr.TransacDetail,
	}
}

func ToListDomain(data []TransactionBuy) (result []transactions.Domain) {
	result = []transactions.Domain{}
	for _, val := range data {
		result = append(result, val.ToDomain())
	}
	return result
}
