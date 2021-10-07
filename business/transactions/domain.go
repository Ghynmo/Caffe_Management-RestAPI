package transactions

import (
	"context"
	// "miniProject/business/tables"
	"time"
)

type Domain struct {
	ID        uint
	UserID    uint
	Menu      string
	Quantity  int
	Payment   int
	Price     int
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UseCase interface {
	BuyTransactionController(ctx context.Context, data Domain) (Domain, error)
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
