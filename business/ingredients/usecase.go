package ingredients

import (
	"context"
	"time"
)

type IngredientUseCase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewIngredientUseCase(repo Repository, timeout time.Duration) UseCase {
	return &IngredientUseCase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *IngredientUseCase) GetIngredientController(ctx context.Context) ([]Domain, error) {
	user, err := uc.Repo.GetIngredient(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return user, nil
}

func (uc *IngredientUseCase) GetIngredientByIDController(ctx context.Context, id int) (Domain, error) {
	user, err := uc.Repo.GetIngredientByID(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *IngredientUseCase) CreateIngredientController(ctx context.Context, data Domain) (Domain, error) {
	user, err := uc.Repo.CreateIngredient(ctx, data)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *IngredientUseCase) UpdateIngredientController(ctx context.Context, data Domain, id int) (Domain, error) {
	user, err := uc.Repo.UpdateIngredient(ctx, data, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *IngredientUseCase) DeleteIngredientController(ctx context.Context, id int) (Domain, error) {
	user, err := uc.Repo.DeleteIngredient(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}
