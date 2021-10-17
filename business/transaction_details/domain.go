package transaction_details

import (
	"context"
	"miniProject/business/menus"
	"miniProject/business/transactions"
)

type Domain struct {
	ID            int
	TransactionID int
	MenuID        int
	Menu          string
	Quantity      int
	Transaction   transactions.Domain
	Menus         menus.Domain
}

type UseCase interface {
	GetTransactionDetailsController(ctx context.Context) ([]Domain, error)
	GetTransactionDetailByIDController(ctx context.Context, id int) (Domain, error)
	CreateTransactionDetailController(ctx context.Context, data Domain) (Domain, error)
	UpdateTransactionDetailController(ctx context.Context, data Domain, id int) (Domain, error)
}

type Repository interface {
	GetTransactionDetails(ctx context.Context) ([]Domain, error)
	GetTransactionDetailByID(ctx context.Context, id int) (Domain, error)
	CreateTransactionDetail(ctx context.Context, data Domain) (Domain, error)
	UpdateTransactionDetail(ctx context.Context, data Domain, id int) (Domain, error)
}
