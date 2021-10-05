package table_details

import (
	"context"
	"time"
)

type Domain struct {
	ID        int
	Capacity  int
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type UseCase interface {
	GetTableDetailController(ctx context.Context) ([]Domain, error)
	GetTableDetailByIDController(ctx context.Context, id int) (Domain, error)
	CreateTableDetailController(ctx context.Context, data Domain) (Domain, error)
	UpdateTableDetailController(ctx context.Context, data Domain, id int) (Domain, error)
	DeleteTableDetailController(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	GetTableDetail(ctx context.Context) ([]Domain, error)
	GetTableDetailByID(ctx context.Context, id int) (Domain, error)
	CreateTableDetail(ctx context.Context, data Domain) (Domain, error)
	UpdateTableDetail(ctx context.Context, data Domain, id int) (Domain, error)
	DeleteTableDetail(ctx context.Context, id int) (Domain, error)
}
