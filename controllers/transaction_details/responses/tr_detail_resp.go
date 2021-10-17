package responses

import (
	"miniProject/business/transaction_details"
	// user "miniProject/controllers/users/responses"
)

type TransactionDetailResponse struct {
	TransactionID int    `json:"transaction_id"`
	Menu          string `json:"menu"`
	Quantity      int    `json:"quantity"`
	UserName      string `json:"name"`
}

func FromDomain(domain transaction_details.Domain) TransactionDetailResponse {
	return TransactionDetailResponse{
		TransactionID: domain.TransactionID,
		Menu:          domain.Menu,
		Quantity:      domain.Quantity,
		UserName:      domain.Name,
	}
}

func FromListDomain(data []transaction_details.Domain) (result []TransactionDetailResponse) {
	result = []TransactionDetailResponse{}
	for _, val := range data {
		result = append(result, FromDomain(val))
	}
	return result
}
