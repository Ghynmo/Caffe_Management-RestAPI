package transactions

import (
	"miniProject/drivers/database/table_details"
	"miniProject/drivers/database/users"

	"gorm.io/gorm"
)

type Transactions struct {
	gorm.Model
	UserID        uint
	users         users.Users `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TableDetailD  uint
	table_details []table_details.TableDetails `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Payment       int
	Price         int
	Status        bool
}

// func (transaction *Transactions) ToDomain() transactions.Domain {
// 	return transactions.Domain{

// 		ID:            transaction.ID,
// 		UserID:        transaction.UserID,
// 		TableDetailID: transaction.TableDetailID,
// 		Payment:       transaction.Payment,
// 		Price:         transaction.Price,
// 		Status:        transaction.Status,
// 		CreatedAt:     transaction.CreatedAt,
// 		UpdatedAt:     transaction.CreatedAt,
// 	}
// }

// func FromDomain(domain transactions.Domain) Transactions {
// 	return Transactions{
// 		ID:            domain.ID,
// 		UserID:        domain.UserID,
// 		TableDetailID: domain.TableDetailID,
// 		Payment:       domain.Payment,
// 		Price:         domain.Price,
// 		Status:        domain.Status,
// 		CreatedAt:     domain.CreatedAt,
// 		UpdatedAt:     domain.CreatedAt,
// 	}
// }

// func ToDomains(u []Transactions) []transactions.Domain {
// 	var Domains []transactions.Domain

// 	for _, val := range u {
// 		Domains = append(Domains, val.ToDomain())
// 	}
// 	return Domains
// }
