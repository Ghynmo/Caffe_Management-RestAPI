package users

import (
	"miniProject/business/users"
	// "miniProject/drivers/database/transactions"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	MemberID int
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Address  string
	Phone    string
	Point    int
	// Transactions []transactions.Transactions `gorm:"foreignKey:UserID"`
}

func (user *Users) ToDomain() users.Domain {
	return users.Domain{
		MemberID:  user.MemberID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Address:   user.Address,
		Phone:     user.Phone,
		Point:     user.Point,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.CreatedAt,
	}
}

func FromDomain(domain users.Domain) Users {
	return Users{
		MemberID: domain.MemberID,
		Name:     domain.Name,
		Email:    domain.Email,
		Password: domain.Password,
		Address:  domain.Address,
		Phone:    domain.Phone,
		Point:    domain.Point,
	}
}

func ToDomains(u []Users) []users.Domain {
	var Domains []users.Domain

	for _, val := range u {
		Domains = append(Domains, val.ToDomain())
	}
	return Domains
}
