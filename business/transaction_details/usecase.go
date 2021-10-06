package transaction_details

import (
	"context"
	"time"
)

type TransactionDetailUseCase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewTransactionDetailUseCase(repo Repository, timeout time.Duration) UseCase {
	return &TransactionDetailUseCase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *TransactionDetailUseCase) GetTransactionDetailsController(ctx context.Context) ([]Domain, error) {
	user, err := uc.Repo.GetTransactionDetails(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return user, nil
}

func (uc *TransactionDetailUseCase) GetTransactionDetailByIDController(ctx context.Context, id uint) (Domain, error) {
	user, err := uc.Repo.GetTransactionDetailByID(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *TransactionDetailUseCase) CreateTransactionDetailController(ctx context.Context, data Domain) (Domain, error) {
	user, err := uc.Repo.CreateTransactionDetail(ctx, data)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *TransactionDetailUseCase) UpdateTransactionDetailController(ctx context.Context, data Domain, id uint) (Domain, error) {
	user, err := uc.Repo.UpdateTransactionDetail(ctx, data, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}
