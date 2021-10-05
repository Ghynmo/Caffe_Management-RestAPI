package table_details

import (
	"context"
	"time"
)

type TableDetailUseCase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewTableDetailUseCase(repo Repository, timeout time.Duration) UseCase {
	return &TableDetailUseCase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *TableDetailUseCase) GetTableDetailController(ctx context.Context) ([]Domain, error) {
	user, err := uc.Repo.GetTableDetail(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return user, nil
}

func (uc *TableDetailUseCase) GetTableDetailByIDController(ctx context.Context, id int) (Domain, error) {
	user, err := uc.Repo.GetTableDetailByID(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *TableDetailUseCase) CreateTableDetailController(ctx context.Context, data Domain) (Domain, error) {
	user, err := uc.Repo.CreateTableDetail(ctx, data)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *TableDetailUseCase) UpdateTableDetailController(ctx context.Context, data Domain, id int) (Domain, error) {
	user, err := uc.Repo.UpdateTableDetail(ctx, data, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *TableDetailUseCase) DeleteTableDetailController(ctx context.Context, id int) (Domain, error) {
	user, err := uc.Repo.DeleteTableDetail(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}
