package inventories

import (
	"context"
	"time"
)

type InventoryUseCase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewInventoryUseCase(repo Repository, timeout time.Duration) UseCase {
	return &InventoryUseCase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *InventoryUseCase) GetInventoriesController(ctx context.Context) ([]Domain, error) {
	user, err := uc.Repo.GetInventories(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return user, nil
}

func (uc *InventoryUseCase) GetInventoryByIDController(ctx context.Context, id int) (Domain, error) {
	user, err := uc.Repo.GetInventoryByID(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *InventoryUseCase) CreateInventoryController(ctx context.Context, data Domain) (Domain, error) {
	user, err := uc.Repo.CreateInventory(ctx, data)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *InventoryUseCase) UpdateInventoryController(ctx context.Context, data Domain, id int) (Domain, error) {
	user, err := uc.Repo.UpdateInventory(ctx, data, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *InventoryUseCase) DeleteInventoryController(ctx context.Context, id int) (Domain, error) {
	user, err := uc.Repo.DeleteInventory(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}
