package table_details

import (
	"gorm.io/gorm"
)

type TableDetails struct {
	gorm.Model
	Capacity int
	Status   bool
}

// func (transaction *TableDetails) ToDomain() transactions.Domain {
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

// func FromDomain(domain transactions.Domain) TableDetails {
// 	return TableDetails{
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

// func ToDomains(u []TableDetails) []transactions.Domain {
// 	var Domains []transactions.Domain

// 	for _, val := range u {
// 		Domains = append(Domains, val.ToDomain())
// 	}
// 	return Domains
// }
