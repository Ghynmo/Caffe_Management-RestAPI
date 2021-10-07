package users

import (
	"context"
	"errors"
	"miniProject/app/middlewares"
	"time"
)

type UserUseCase struct {
	ConfigJWT      middlewares.ConfigJwt
	Repo           Repository
	contextTimeout time.Duration
}

func NewUserUseCase(repo Repository, timeout time.Duration, jwtconfig middlewares.ConfigJwt) UseCase {
	return &UserUseCase{
		ConfigJWT:      jwtconfig,
		Repo:           repo,
		contextTimeout: timeout,
	}
}

//core business login
func (uc *UserUseCase) Login(ctx context.Context, email string, password string) (Domain, error) {
	if email == "" {
		return Domain{}, errors.New("email empty")
	}
	if password == "" {
		return Domain{}, errors.New("password empty")
	}

	// var err error
	// password, err = encrypt.Hash(password)
	// if err != nil {
	// 	return Domain{}, nil
	// }

	user, err := uc.Repo.Login(ctx, email, password)
	if err != nil {
		return Domain{}, err
	}

	// if !(encrypt.ValidateHash(password, user.Password)) {
	// 	return Domain{}, errors.New("wrong password")
	// }

	user.Token, err = uc.ConfigJWT.GenerateToken(uint(user.ID))
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *UserUseCase) GetUserController(ctx context.Context) ([]Domain, error) {
	user, err := uc.Repo.GetUser(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return user, nil
}

func (uc *UserUseCase) GetUserByIDController(ctx context.Context, id int) (Domain, error) {
	user, err := uc.Repo.GetUserByID(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *UserUseCase) CreateUserController(ctx context.Context, data Domain) (Domain, error) {
	user, err := uc.Repo.CreateUser(ctx, data)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *UserUseCase) UpdateUserController(ctx context.Context, data Domain, id int) (Domain, error) {
	user, err := uc.Repo.UpdateUser(ctx, data, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *UserUseCase) DeleteUserController(ctx context.Context, id int) (Domain, error) {
	user, err := uc.Repo.DeleteUser(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}
