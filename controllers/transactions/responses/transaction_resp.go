package responses

import (
	"miniProject/business/transactions"
	user "miniProject/controllers/users/responses"
)

type TransactionResponse struct {
	UserID int  `json:"user_id"`
	Price  int  `json:"price"`
	Status bool `json:"status"`
	User   user.UserResponse
}

func FromDomain(domain transactions.Domain) TransactionResponse {
	return TransactionResponse{
		UserID: domain.UserID,
		Price:  domain.Price,
		Status: domain.Status,
		User:   user.FromDomain(domain.User),
	}
}

func FromListDomain(data []transactions.Domain) (result []TransactionResponse) {
	result = []TransactionResponse{}
	for _, val := range data {
		result = append(result, FromDomain(val))
	}
	return result
}
