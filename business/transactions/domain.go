package transactions

import (
	"context"
	"miniProject/business/users"
	"time"
)

type Domain struct {
	ID        int
	UserID    int
	User      users.Domain
	Price     int
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UseCase interface {
	GetTransactionController(ctx context.Context) ([]Domain, error)
	GetTransactionByIDController(ctx context.Context, id int) (Domain, error)
	CreateTransactionController(ctx context.Context, data Domain) (Domain, error)
	DeleteTransactionController(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	GetTransaction(ctx context.Context) ([]Domain, error)
	GetTransactionByID(ctx context.Context, id int) (Domain, error)
	CreateTransaction(ctx context.Context, data Domain) (Domain, error)
	DeleteTransaction(ctx context.Context, id int) (Domain, error)
}
