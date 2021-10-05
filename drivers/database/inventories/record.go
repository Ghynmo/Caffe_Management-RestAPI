package inventories

import (
	"gorm.io/gorm"
)

type Inventorys struct {
	gorm.Model
	RecipeID   int
	MaterialID int
	Quantity   int
}

// func (inventory *Inventorys) ToDomain() inventories.Domain {
// 	return inventories.Domain{

// 		ID:            inventory.ID,
// 		UserID:        inventory.UserID,
// 		TableDetailID: inventory.TableDetailID,
// 		Payment:       inventory.Payment,
// 		Price:         inventory.Price,
// 		Status:        inventory.Status,
// 		CreatedAt:     inventory.CreatedAt,
// 		UpdatedAt:     inventory.CreatedAt,
// 	}
// }

// func FromDomain(domain inventories.Domain) Inventorys {
// 	return Inventorys{
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

// func ToDomains(u []Inventorys) []inventories.Domain {
// 	var Domains []inventories.Domain

// 	for _, val := range u {
// 		Domains = append(Domains, val.ToDomain())
// 	}
// 	return Domains
// }
