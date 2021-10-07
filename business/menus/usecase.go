package menus

import (
	"context"
	"time"
)

type MenuUseCase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewMenuUseCase(repo Repository, timeout time.Duration) UseCase {
	return &MenuUseCase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *MenuUseCase) GetMenusController(ctx context.Context) ([]Domain, error) {
	user, err := uc.Repo.GetMenu(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return user, nil
}

func (uc *MenuUseCase) GetMenuByIDController(ctx context.Context, id uint) (Domain, error) {
	user, err := uc.Repo.GetMenuByID(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *MenuUseCase) CreateMenuController(ctx context.Context, data Domain) (Domain, error) {
	user, err := uc.Repo.CreateMenu(ctx, data)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *MenuUseCase) SaveApiController(ctx context.Context, data Domain) (Domain, error) {
	user, err := uc.Repo.CreateMenu(ctx, data)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *MenuUseCase) UpdateMenuController(ctx context.Context, data Domain, id uint) (Domain, error) {
	user, err := uc.Repo.UpdateMenu(ctx, data, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *MenuUseCase) DeleteMenuController(ctx context.Context, id uint) (Domain, error) {
	user, err := uc.Repo.DeleteMenu(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *MenuUseCase) GetMenuAPI(ctx context.Context, key string) (Domain, error) {
	user, err := uc.Repo.GetMenuAPI(ctx, key)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}
