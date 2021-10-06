package transaction_details

import (
	"context"
	"time"
)

type Domain struct {
	ID            uint
	TransactionID uint
	MenuID        uint
	Quantity      int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type UseCase interface {
	GetTransactionDetailsController(ctx context.Context) ([]Domain, error)
	GetTransactionDetailByIDController(ctx context.Context, id uint) (Domain, error)
	CreateTransactionDetailController(ctx context.Context, data Domain) (Domain, error)
	UpdateTransactionDetailController(ctx context.Context, data Domain, id uint) (Domain, error)
}

type Repository interface {
	GetTransactionDetails(ctx context.Context) ([]Domain, error)
	GetTransactionDetailByID(ctx context.Context, id uint) (Domain, error)
	CreateTransactionDetail(ctx context.Context, data Domain) (Domain, error)
	UpdateTransactionDetail(ctx context.Context, data Domain, id uint) (Domain, error)
}
