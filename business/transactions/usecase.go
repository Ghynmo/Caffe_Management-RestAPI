package transactions

import (
	"context"
	"fmt"
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

func (uc *TransactionUseCase) BuyTransactionController(ctx context.Context, data Domain) (Domain, error) {
	user, err := uc.Repo.BuyTransaction(ctx, data)
	if err != nil {
		return Domain{}, err
	}
	fmt.Println(user)
	return user, nil
}

func (uc *TransactionUseCase) GetTransactionByIDController(ctx context.Context, id uint) (Domain, error) {
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

func (uc *TransactionUseCase) UpdateTransactionController(ctx context.Context, data Domain, id uint) (Domain, error) {
	user, err := uc.Repo.UpdateTransaction(ctx, data, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}
