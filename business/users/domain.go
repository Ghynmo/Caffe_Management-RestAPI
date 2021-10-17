package users

import (
	"context"
	"time"
)

type Domain struct {
	ID        int
	Name      string
	Email     string
	Password  string
	Token     string
	Address   string
	Phone     string
	Point     int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type UseCase interface {
	Login(ctx context.Context, email string, password string) (Domain, error)
	GetUserController(ctx context.Context) ([]Domain, error)
	GetUserByIDController(ctx context.Context, id int) (Domain, error)
	CreateUserController(ctx context.Context, data Domain) (Domain, error)
	UpdateUserController(ctx context.Context, data Domain, id int) (Domain, error)
	DeleteUserController(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	Login(ctx context.Context, email string, password string) (Domain, error)
	GetUser(ctx context.Context) ([]Domain, error)
	GetUserByID(ctx context.Context, id int) (Domain, error)
	CreateUser(ctx context.Context, data Domain) (Domain, error)
	UpdateUser(ctx context.Context, data Domain, id int) (Domain, error)
	DeleteUser(ctx context.Context, id int) (Domain, error)
}
