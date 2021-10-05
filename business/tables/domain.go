package tables

import (
	"context"

	"gorm.io/gorm"
)

type Domain struct {
	gorm.Model
	Capacity int
	Status   bool
}

type UseCase interface {
	GetTableController(ctx context.Context) ([]Domain, error)
	GetTableByIDController(ctx context.Context, id int) (Domain, error)
	CreateTableController(ctx context.Context, data Domain) (Domain, error)
	UpdateTableController(ctx context.Context, data Domain, id int) (Domain, error)
	DeleteTableController(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	GetTable(ctx context.Context) ([]Domain, error)
	GetTableByID(ctx context.Context, id int) (Domain, error)
	CreateTable(ctx context.Context, data Domain) (Domain, error)
	UpdateTable(ctx context.Context, data Domain, id int) (Domain, error)
	DeleteTable(ctx context.Context, id int) (Domain, error)
}
