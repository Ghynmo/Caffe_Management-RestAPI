package transactions

import (
	"miniProject/business/transactions"

	"gorm.io/gorm"
)

type Transactions struct {
	gorm.Model
	UserID        uint
	TableDetailID uint
	Payment       int
	Price         int
	Status        bool
}

func (transaction *Transactions) ToDomain() transactions.Domain {
	return transactions.Domain{

		ID:            transaction.ID,
		UserID:        transaction.UserID,
		TableDetailID: transaction.TableDetailID,
		Payment:       transaction.Payment,
		Price:         transaction.Price,
		Status:        transaction.Status,
		CreatedAt:     transaction.CreatedAt,
		UpdatedAt:     transaction.CreatedAt,
	}
}

func FromDomain(domain transactions.Domain) Transactions {
	return Transactions{
		UserID:        domain.UserID,
		TableDetailID: domain.TableDetailID,
		Payment:       domain.Payment,
		Price:         domain.Price,
		Status:        domain.Status,
	}
}

func ToDomains(u []Transactions) []transactions.Domain {
	var Domains []transactions.Domain

	for _, val := range u {
		Domains = append(Domains, val.ToDomain())
	}
	return Domains
}
