package menus

import (
	"gorm.io/gorm"
)

type Menus struct {
	gorm.Model
	RecipeID    int
	Name        string
	Category    string
	Price       int
	Stock       int
	Description string
}

// func (menu *Menus) ToDomain() menus.Domain {
// 	return menus.Domain{

// 		ID:            menu.ID,
// 		UserID:        menu.UserID,
// 		TableDetailID: menu.TableDetailID,
// 		Payment:       menu.Payment,
// 		Price:         menu.Price,
// 		Status:        menu.Status,
// 		CreatedAt:     menu.CreatedAt,
// 		UpdatedAt:     menu.CreatedAt,
// 	}
// }

// func FromDomain(domain menus.Domain) Menus {
// 	return Menus{
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

// func ToDomains(u []Menus) []menus.Domain {
// 	var Domains []menus.Domain

// 	for _, val := range u {
// 		Domains = append(Domains, val.ToDomain())
// 	}
// 	return Domains
// }
