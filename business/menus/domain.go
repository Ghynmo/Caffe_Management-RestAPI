package menus

import (
	"context"
	"time"
)

type Domain struct {
	ID        int
	Name      string
	Category  string
	Price     int
	Stock     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UseCase interface {
	GetMenusController(ctx context.Context) ([]Domain, error)
	GetMenuByIDController(ctx context.Context, id int) (Domain, error)
	CreateMenuController(ctx context.Context, data Domain) (Domain, error)
	UpdateMenuController(ctx context.Context, data Domain, id int) (Domain, error)
	DeleteMenuController(ctx context.Context, id int) (Domain, error)
	GetMenuAPI(ctx context.Context, key string) (Domain, error)
}

type Repository interface {
	GetMenu(ctx context.Context) ([]Domain, error)
	GetMenuByID(ctx context.Context, id int) (Domain, error)
	CreateMenu(ctx context.Context, data Domain) (Domain, error)
	UpdateMenu(ctx context.Context, data Domain, id int) (Domain, error)
	DeleteMenu(ctx context.Context, id int) (Domain, error)
	GetMenuAPI(ctx context.Context, key string) (Domain, error)
}
