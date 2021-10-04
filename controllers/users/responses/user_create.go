package responses

import "miniProject/business/users"

type UserInsert struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

func CreateFromDomain(domain users.Domain) UserInsert {
	return UserInsert{
		Name:    domain.Name,
		Email:   domain.Email,
		Address: domain.Address,
		Phone:   domain.Phone,
	}
}
