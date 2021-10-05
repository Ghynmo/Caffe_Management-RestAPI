package ingredients

import (
	"gorm.io/gorm"
)

type Ingredients struct {
	gorm.Model
	RecipeID   int
	MaterialID int
	Quantity   int
}

// func (ingredient *Ingredients) ToDomain() ingredients.Domain {
// 	return ingredients.Domain{

// 		ID:            ingredient.ID,
// 		UserID:        ingredient.UserID,
// 		TableDetailID: ingredient.TableDetailID,
// 		Payment:       ingredient.Payment,
// 		Price:         ingredient.Price,
// 		Status:        ingredient.Status,
// 		CreatedAt:     ingredient.CreatedAt,
// 		UpdatedAt:     ingredient.CreatedAt,
// 	}
// }

// func FromDomain(domain ingredients.Domain) Ingredients {
// 	return Ingredients{
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

// func ToDomains(u []Ingredients) []ingredients.Domain {
// 	var Domains []ingredients.Domain

// 	for _, val := range u {
// 		Domains = append(Domains, val.ToDomain())
// 	}
// 	return Domains
// }
