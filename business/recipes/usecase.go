package recipes

import (
	"context"
	"time"
)

type RecipeUseCase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewRecipeUseCase(repo Repository, timeout time.Duration) UseCase {
	return &RecipeUseCase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *RecipeUseCase) GetRecipeController(ctx context.Context) ([]Domain, error) {
	user, err := uc.Repo.GetRecipe(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return user, nil
}

func (uc *RecipeUseCase) GetRecipeByIDController(ctx context.Context, id int) (Domain, error) {
	user, err := uc.Repo.GetRecipeByID(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *RecipeUseCase) CreateRecipeController(ctx context.Context, data Domain) (Domain, error) {
	user, err := uc.Repo.CreateRecipe(ctx, data)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *RecipeUseCase) UpdateRecipeController(ctx context.Context, data Domain, id int) (Domain, error) {
	user, err := uc.Repo.UpdateRecipe(ctx, data, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *RecipeUseCase) DeleteRecipeController(ctx context.Context, id int) (Domain, error) {
	user, err := uc.Repo.DeleteRecipe(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}
