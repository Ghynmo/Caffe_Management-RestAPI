package transactions

import (
	"miniProject/business/transactions"

	"gorm.io/gorm"
)

type Transactions struct {
	gorm.Model
	UserID   uint
	Menu     string
	Quantity int
	Payment  int
	Price    int
	Status   bool
}

func (transaction *Transactions) ToDomain() transactions.Domain {
	return transactions.Domain{
		ID:        uint(transaction.ID),
		UserID:    uint(transaction.UserID),
		Payment:   transaction.Payment,
		Price:     transaction.Price,
		Status:    transaction.Status,
		CreatedAt: transaction.CreatedAt,
		UpdatedAt: transaction.CreatedAt,
	}
}

func FromDomain(domain transactions.Domain) Transactions {
	return Transactions{
		UserID:  domain.UserID,
		Payment: domain.Payment,
		Price:   domain.Price,
		Status:  domain.Status,
	}
}

func ToDomains(u []Transactions) []transactions.Domain {
	var Domains []transactions.Domain

	for _, val := range u {
		Domains = append(Domains, val.ToDomain())
	}
	return Domains
}

func ToListDomain(data []Transactions) (result []transactions.Domain) {
	result = []transactions.Domain{}
	for _, val := range data {
		result = append(result, val.ToDomain())
	}
	return result
}
