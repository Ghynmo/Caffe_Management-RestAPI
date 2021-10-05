package transactions

import (
	"context"
	"time"
)

type Domain struct {
	ID            int
	TransactionID int
	TableDetailID int
	Payment       int
	Price         int
	Status        bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

type UseCase interface {
	GetTransactionController(ctx context.Context) ([]Domain, error)
	GetTransactionByIDController(ctx context.Context, id int) (Domain, error)
	CreateTransactionController(ctx context.Context, data Domain) (Domain, error)
	UpdateTransactionController(ctx context.Context, data Domain, id int) (Domain, error)
	DeleteTransactionController(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	GetTransaction(ctx context.Context) ([]Domain, error)
	GetTransactionByID(ctx context.Context, id int) (Domain, error)
	CreateTransaction(ctx context.Context, data Domain) (Domain, error)
	UpdateTransaction(ctx context.Context, data Domain, id int) (Domain, error)
	DeleteTransaction(ctx context.Context, id int) (Domain, error)
}
