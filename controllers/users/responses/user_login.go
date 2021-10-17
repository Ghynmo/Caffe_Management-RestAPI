package responses

import (
	"miniProject/business/users"
	"time"
)

type UserResponseLogin struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	// Password  string    `json:"password"`
	Token     string    `json:"token"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	Point     int       `json:"point"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}

func FromDomain_Login(domain users.Domain) UserResponseLogin {
	return UserResponseLogin{
		ID:    domain.ID,
		Name:  domain.Name,
		Email: domain.Email,
		// Password:  domain.Password,
		Token:     domain.Token,
		Address:   domain.Address,
		Phone:     domain.Phone,
		Point:     domain.Point,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.CreatedAt,
	}
}
