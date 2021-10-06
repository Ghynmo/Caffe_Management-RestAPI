package transactions

import (
	"context"
	"time"
)

type Domain struct {
	ID            uint
	UserID        uint
	TableDetailID uint
	Payment       int
	Price         int
	Status        bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type UseCase interface {
	GetTransactionsController(ctx context.Context) ([]Domain, error)
	GetTransactionByIDController(ctx context.Context, id uint) (Domain, error)
	CreateTransactionController(ctx context.Context, data Domain) (Domain, error)
	UpdateTransactionController(ctx context.Context, data Domain, id uint) (Domain, error)
}

type Repository interface {
	GetTransactions(ctx context.Context) ([]Domain, error)
	GetTransactionByID(ctx context.Context, id uint) (Domain, error)
	CreateTransaction(ctx context.Context, data Domain) (Domain, error)
	UpdateTransaction(ctx context.Context, data Domain, id uint) (Domain, error)
}
