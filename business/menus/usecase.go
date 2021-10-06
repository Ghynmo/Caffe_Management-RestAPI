package menus

import (
	"context"
	"miniProject/business/masak_openapi"
	"time"
)

type MenuUseCase struct {
	Repo           Repository
	RepoMasak      masak_openapi.Repository
	contextTimeout time.Duration
}

func NewMenuUseCase(repo Repository, rm masak_openapi.Repository, timeout time.Duration) UseCase {
	return &MenuUseCase{
		Repo:           repo,
		RepoMasak:      rm,
		contextTimeout: timeout,
	}
}

func (uc *MenuUseCase) APIDetail(menu Domain, key string) (Domain, error) {
	newMenu, err := uc.RepoMasak.GetMenuApi(menu.Name)
	if err != nil {
		return Domain{}, err
	}
	menu.Name = newMenu.Name

	return menu, nil
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

func (uc *MenuUseCase) APIz(key string) (Domain, error) {
	var menu Domain
	newMenu, err := uc.RepoMasak.GetMenuApi(key)
	if err != nil {
		return Domain{}, err
	}
	menu.Name = newMenu.Name

	return menu, nil
}
