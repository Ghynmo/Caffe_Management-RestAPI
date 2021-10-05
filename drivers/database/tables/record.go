package tables

import (
	"gorm.io/gorm"
)

type Tables struct {
	gorm.Model
	Capacity int
	Status   bool
}

// func (table *Tables) ToDomain() tables.Domain {
// 	return tables.Domain{

// 		ID:            table.ID,
// 		UserID:        table.UserID,
// 		TableDetailID: table.TableDetailID,
// 		Payment:       table.Payment,
// 		Price:         table.Price,
// 		Status:        table.Status,
// 		CreatedAt:     table.CreatedAt,
// 		UpdatedAt:     table.CreatedAt,
// 	}
// }

// func FromDomain(domain tables.Domain) Tables {
// 	return Tables{
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

// func ToDomains(u []Tables) []tables.Domain {
// 	var Domains []tables.Domain

// 	for _, val := range u {
// 		Domains = append(Domains, val.ToDomain())
// 	}
// 	return Domains
// }
