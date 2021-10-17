package transaction_details

import (
	"context"
	"fmt"
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

func (uc *TransactionDetailUseCase) GetTransactionDetailByIDController(ctx context.Context, id int) (Domain, error) {
	user, err := uc.Repo.GetTransactionDetailByID(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *TransactionDetailUseCase) CreateTransactionDetailController(ctx context.Context, data Domain) (Domain, error) {
	fmt.Println("Usecase", data)

	result, err := uc.Repo.CreateTransactionDetail(ctx, data)
	if err != nil {
		return Domain{}, err
	}

	fmt.Println("after DB", result)

	return result, nil
}

func (uc *TransactionDetailUseCase) UpdateTransactionDetailController(ctx context.Context, data Domain, id int) (Domain, error) {
	user, err := uc.Repo.UpdateTransactionDetail(ctx, data, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *TransactionDetailUseCase) CreateListTransactionDetailController(ctx context.Context, data []Domain) ([]Domain, error) {
	var results []Domain
	fmt.Println(data)

	for _, val := range data {
		result, err := uc.Repo.CreateTransactionDetail(ctx, val)
		if err != nil {
			return []Domain{}, err
		}
		results = append(results, result)
	}

	return results, nil
}
