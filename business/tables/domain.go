package tables

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
	GetTablesController(ctx context.Context) ([]Domain, error)
	GetTableByIDController(ctx context.Context, id int) (Domain, error)
	CreateTableController(ctx context.Context, data Domain) (Domain, error)
	UpdateTableController(ctx context.Context, data Domain, id int) (Domain, error)
	DeleteTableController(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	GetTables(ctx context.Context) ([]Domain, error)
	GetTableByID(ctx context.Context, id int) (Domain, error)
	CreateTable(ctx context.Context, data Domain) (Domain, error)
	UpdateTable(ctx context.Context, data Domain, id int) (Domain, error)
	DeleteTable(ctx context.Context, id int) (Domain, error)
}
