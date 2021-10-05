package recipes

import (
	"context"
	"miniProject/business/recipes"

	"gorm.io/gorm"
)

type MysqlRecipeRepository struct {
	Conn *gorm.DB
}

func NewMysqlRecipeRepository(conn *gorm.DB) recipes.Repository {
	return &MysqlRecipeRepository{
		Conn: conn,
	}
}

func (DB *MysqlRecipeRepository) GetRecipe(ctx context.Context) ([]recipes.Domain, error) {
	return []recipes.Domain{}, nil

	// var CurrentRecipe []Recipe
	// result := DB.Conn.Find(&CurrentRecipe)
	// if result.Error != nil {
	// 	return []recipes.Domain{}, result.Error
	// }
	// return ToDomains(CurrentRecipe), nil
}

func (DB *MysqlRecipeRepository) GetRecipeByID(ctx context.Context, id int) (recipes.Domain, error) {
	return recipes.Domain{}, nil

	// var Currentrecipe Recipe
	// result := DB.Conn.Where("id = ?", id).First(&Currentrecipe)
	// if result.Error != nil {
	// 	return recipes.Domain{}, result.Error
	// }
	// return Currentrecipe.ToDomain(), nil
}

func (DB *MysqlRecipeRepository) CreateRecipe(ctx context.Context, data recipes.Domain) (recipes.Domain, error) {
	return recipes.Domain{}, nil

	// InsertRecipe := FromDomain(data)
	// result := DB.Conn.Create(&InsertRecipe)
	// if result.Error != nil {
	// 	return recipes.Domain{}, result.Error
	// }

	// return InsertRecipe.ToDomain(), nil
}

func (DB *MysqlRecipeRepository) UpdateRecipe(ctx context.Context, data recipes.Domain, id int) (recipes.Domain, error) {
	return recipes.Domain{}, nil

	// Currentrecipe := FromDomain(data)
	// result := DB.Conn.Where("id = ?", id).Updates(Currentrecipe).First(&Currentrecipe)
	// if result.Error != nil {
	// 	return recipes.Domain{}, result.Error
	// }
	// return Currentrecipe.ToDomain(), nil
}

func (DB *MysqlRecipeRepository) DeleteRecipe(ctx context.Context, id int) (recipes.Domain, error) {
	return recipes.Domain{}, nil

	// var Currentrecipe Recipe
	// result := DB.Conn.Where("id = ?", id).Delete(&Currentrecipe)
	// if result.Error != nil {
	// 	return recipes.Domain{}, result.Error
	// }
	// return Currentrecipe.ToDomain(), nil
}
