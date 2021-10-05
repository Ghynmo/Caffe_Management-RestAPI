package ingredients

import (
	"context"
	"miniProject/business/ingredients"

	"gorm.io/gorm"
)

type MysqlIngredientRepository struct {
	Conn *gorm.DB
}

func NewMysqlIngredientRepository(conn *gorm.DB) ingredients.Repository {
	return &MysqlIngredientRepository{
		Conn: conn,
	}
}

func (DB *MysqlIngredientRepository) GetIngredient(ctx context.Context) ([]ingredients.Domain, error) {
	return []ingredients.Domain{}, nil

	// var CurrentIngredient []Ingredient
	// result := DB.Conn.Find(&CurrentIngredient)
	// if result.Error != nil {
	// 	return []ingredients.Domain{}, result.Error
	// }
	// return ToDomains(CurrentIngredient), nil
}

func (DB *MysqlIngredientRepository) GetIngredientByID(ctx context.Context, id int) (ingredients.Domain, error) {
	return ingredients.Domain{}, nil

	// var Currentingredient Ingredient
	// result := DB.Conn.Where("id = ?", id).First(&Currentingredient)
	// if result.Error != nil {
	// 	return ingredients.Domain{}, result.Error
	// }
	// return Currentingredient.ToDomain(), nil
}

func (DB *MysqlIngredientRepository) CreateIngredient(ctx context.Context, data ingredients.Domain) (ingredients.Domain, error) {
	return ingredients.Domain{}, nil

	// InsertIngredient := FromDomain(data)
	// result := DB.Conn.Create(&InsertIngredient)
	// if result.Error != nil {
	// 	return ingredients.Domain{}, result.Error
	// }

	// return InsertIngredient.ToDomain(), nil
}

func (DB *MysqlIngredientRepository) UpdateIngredient(ctx context.Context, data ingredients.Domain, id int) (ingredients.Domain, error) {
	return ingredients.Domain{}, nil

	// Currentingredient := FromDomain(data)
	// result := DB.Conn.Where("id = ?", id).Updates(Currentingredient).First(&Currentingredient)
	// if result.Error != nil {
	// 	return ingredients.Domain{}, result.Error
	// }
	// return Currentingredient.ToDomain(), nil
}

func (DB *MysqlIngredientRepository) DeleteIngredient(ctx context.Context, id int) (ingredients.Domain, error) {
	return ingredients.Domain{}, nil

	// var Currentingredient Ingredient
	// result := DB.Conn.Where("id = ?", id).Delete(&Currentingredient)
	// if result.Error != nil {
	// 	return ingredients.Domain{}, result.Error
	// }
	// return Currentingredient.ToDomain(), nil
}
