package responses

import (
	"miniProject/business/users"
	"time"
)

type UserResponse struct {
	ID       int    `json:"id"`
	MemberID int    `json:"memberid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	// Password  string    `json:"password"`
	Token     string    `json:"token"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	Point     int       `json:"point"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}

func FromDomain(domain users.Domain) UserResponse {
	return UserResponse{
		MemberID: domain.MemberID,
		ID:       domain.ID,
		Name:     domain.Name,
		Email:    domain.Email,
		// Password:  domain.Password,
		Token:     domain.Token,
		Address:   domain.Address,
		Phone:     domain.Phone,
		Point:     domain.Point,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.CreatedAt,
	}
}
