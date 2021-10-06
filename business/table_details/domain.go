package table_details

import (
	"context"
	"time"
)

type Domain struct {
	ID            uint
	TotalCapacity int
	Status        bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type UseCase interface {
	GetTableDetailsController(ctx context.Context) ([]Domain, error)
	GetTableDetailByIDController(ctx context.Context, id uint) (Domain, error)
	CreateTableDetailController(ctx context.Context, data Domain) (Domain, error)
	UpdateTableDetailController(ctx context.Context, data Domain, id uint) (Domain, error)
}

type Repository interface {
	GetTableDetails(ctx context.Context) ([]Domain, error)
	GetTableDetailByID(ctx context.Context, id uint) (Domain, error)
	CreateTableDetail(ctx context.Context, data Domain) (Domain, error)
	UpdateTableDetail(ctx context.Context, data Domain, id uint) (Domain, error)
}
