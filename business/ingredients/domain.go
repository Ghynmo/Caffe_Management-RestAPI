package ingredients

import (
	"context"
	"time"
)

type Domain struct {
	RecipeID   int
	MaterialID int
	Quantity   int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type UseCase interface {
	GetIngredientController(ctx context.Context) ([]Domain, error)
	GetIngredientByIDController(ctx context.Context, id int) (Domain, error)
	CreateIngredientController(ctx context.Context, data Domain) (Domain, error)
	UpdateIngredientController(ctx context.Context, data Domain, id int) (Domain, error)
	DeleteIngredientController(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	GetIngredient(ctx context.Context) ([]Domain, error)
	GetIngredientByID(ctx context.Context, id int) (Domain, error)
	CreateIngredient(ctx context.Context, data Domain) (Domain, error)
	UpdateIngredient(ctx context.Context, data Domain, id int) (Domain, error)
	DeleteIngredient(ctx context.Context, id int) (Domain, error)
}
