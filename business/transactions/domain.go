package transactions

import (
	"context"
	"miniProject/business/transaction_details"
	"time"
)

type Domain struct {
	ID            uint
	UserID        uint
	TransacDetail []transaction_details.Domain
	Payment       int
	Price         int
	Status        bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type UseCase interface {
	GetTransactionByIDController(ctx context.Context, id uint) (Domain, error)
	CreateTransactionController(ctx context.Context, data Domain) (Domain, error)
	UpdateTransactionController(ctx context.Context, data Domain, id uint) (Domain, error)
}

type Repository interface {
	BuyTransaction(ctx context.Context, data Domain) (Domain, error)
	GetTransactionByID(ctx context.Context, id uint) (Domain, error)
	CreateTransaction(ctx context.Context, data Domain) (Domain, error)
	UpdateTransaction(ctx context.Context, data Domain, id uint) (Domain, error)
}
