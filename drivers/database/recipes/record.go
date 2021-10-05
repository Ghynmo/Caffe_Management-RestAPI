package recipes

import (
	"gorm.io/gorm"
)

type Recipes struct {
	gorm.Model
	Name              string
	Step              string
	EstimateInMinutes int
}

// func (recipe *Recipes) ToDomain() recipes.Domain {
// 	return recipes.Domain{

// 		ID:            recipe.ID,
// 		UserID:        recipe.UserID,
// 		TableDetailID: recipe.TableDetailID,
// 		Payment:       recipe.Payment,
// 		Price:         recipe.Price,
// 		Status:        recipe.Status,
// 		CreatedAt:     recipe.CreatedAt,
// 		UpdatedAt:     recipe.CreatedAt,
// 	}
// }

// func FromDomain(domain recipes.Domain) Recipes {
// 	return Recipes{
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

// func ToDomains(u []Recipes) []recipes.Domain {
// 	var Domains []recipes.Domain

// 	for _, val := range u {
// 		Domains = append(Domains, val.ToDomain())
// 	}
// 	return Domains
// }
