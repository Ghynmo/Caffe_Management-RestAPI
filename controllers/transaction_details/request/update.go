package request

import "miniProject/business/transaction_details"

type TransactionDetailUpdate struct {
	Quantity int `json:"quantity"`
}

func (tr *TransactionDetailUpdate) UpdateToDomain() transaction_details.Domain {
	return transaction_details.Domain{
		Quantity: tr.Quantity,
	}
}

func ToListDomain(data []TransactionDetailUpdate) (result []transaction_details.Domain) {
	result = []transaction_details.Domain{}
	for _, val := range data {
		result = append(result, val.UpdateToDomain())
	}
	return result
}
