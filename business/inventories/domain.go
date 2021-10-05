package inventories

import (
	"context"
	"time"
)

type Domain struct {
	ID          uint
	Name        int
	Measure     int
	MeasureType string
	UnitPrice   int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UseCase interface {
	GetInventoryController(ctx context.Context) ([]Domain, error)
	GetInventoryByIDController(ctx context.Context, id int) (Domain, error)
	CreateInventoryController(ctx context.Context, data Domain) (Domain, error)
	UpdateInventoryController(ctx context.Context, data Domain, id int) (Domain, error)
	DeleteInventoryController(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	GetInventory(ctx context.Context) ([]Domain, error)
	GetInventoryByID(ctx context.Context, id int) (Domain, error)
	CreateInventory(ctx context.Context, data Domain) (Domain, error)
	UpdateInventory(ctx context.Context, data Domain, id int) (Domain, error)
	DeleteInventory(ctx context.Context, id int) (Domain, error)
}
