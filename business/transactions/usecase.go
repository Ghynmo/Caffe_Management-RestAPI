package transactions

import (
	"context"
	"time"
)

type TransactionUseCase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewTransactionUseCase(repo Repository, timeout time.Duration) UseCase {
	return &TransactionUseCase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *TransactionUseCase) GetTransactionController(ctx context.Context) ([]Domain, error) {
	user, err := uc.Repo.GetTransaction(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return user, nil
}

func (uc *TransactionUseCase) GetTransactionByIDController(ctx context.Context, id int) (Domain, error) {
	user, err := uc.Repo.GetTransactionByID(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *TransactionUseCase) CreateTransactionController(ctx context.Context, data Domain) (Domain, error) {
	user, err := uc.Repo.CreateTransaction(ctx, data)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *TransactionUseCase) UpdateTransactionController(ctx context.Context, data Domain, id int) (Domain, error) {
	user, err := uc.Repo.UpdateTransaction(ctx, data, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *TransactionUseCase) DeleteTransactionController(ctx context.Context, id int) (Domain, error) {
	user, err := uc.Repo.DeleteTransaction(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}
