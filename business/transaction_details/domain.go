package transaction_details

import (
	"context"
	"time"
)

type Domain struct {
	ID            int
	TransactionID int
	MenuID        int
	Quantity      int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type UseCase interface {
	GetTransactionDetailController(ctx context.Context) ([]Domain, error)
	GetTransactionDetailByIDController(ctx context.Context, id int) (Domain, error)
	CreateTransactionDetailController(ctx context.Context, data Domain) (Domain, error)
	UpdateTransactionDetailController(ctx context.Context, data Domain, id int) (Domain, error)
	DeleteTransactionDetailController(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	GetTransactionDetail(ctx context.Context) ([]Domain, error)
	GetTransactionDetailByID(ctx context.Context, id int) (Domain, error)
	CreateTransactionDetail(ctx context.Context, data Domain) (Domain, error)
	UpdateTransactionDetail(ctx context.Context, data Domain, id int) (Domain, error)
	DeleteTransactionDetail(ctx context.Context, id int) (Domain, error)
}
