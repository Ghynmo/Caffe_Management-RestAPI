package tables

import (
	"context"
	"time"
)

type TableUseCase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewTableUseCase(repo Repository, timeout time.Duration) UseCase {
	return &TableUseCase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *TableUseCase) GetTableController(ctx context.Context) ([]Domain, error) {
	user, err := uc.Repo.GetTable(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return user, nil
}

func (uc *TableUseCase) GetTableByIDController(ctx context.Context, id int) (Domain, error) {
	user, err := uc.Repo.GetTableByID(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *TableUseCase) CreateTableController(ctx context.Context, data Domain) (Domain, error) {
	user, err := uc.Repo.CreateTable(ctx, data)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *TableUseCase) UpdateTableController(ctx context.Context, data Domain, id int) (Domain, error) {
	user, err := uc.Repo.UpdateTable(ctx, data, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *TableUseCase) DeleteTableController(ctx context.Context, id int) (Domain, error) {
	user, err := uc.Repo.DeleteTable(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}
