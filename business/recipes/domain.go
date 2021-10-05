package recipes

import (
	"context"
	"time"
)

type Domain struct {
	ID                int
	Name              string
	Step              string
	EstimateInMinutes int
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type UseCase interface {
	GetRecipeController(ctx context.Context) ([]Domain, error)
	GetRecipeByIDController(ctx context.Context, id int) (Domain, error)
	CreateRecipeController(ctx context.Context, data Domain) (Domain, error)
	UpdateRecipeController(ctx context.Context, data Domain, id int) (Domain, error)
	DeleteRecipeController(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	GetRecipe(ctx context.Context) ([]Domain, error)
	GetRecipeByID(ctx context.Context, id int) (Domain, error)
	CreateRecipe(ctx context.Context, data Domain) (Domain, error)
	UpdateRecipe(ctx context.Context, data Domain, id int) (Domain, error)
	DeleteRecipe(ctx context.Context, id int) (Domain, error)
}
