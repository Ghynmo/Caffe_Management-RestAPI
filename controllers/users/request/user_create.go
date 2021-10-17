package request

import "miniProject/business/users"

type UserInsert struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

func (ui *UserInsert) ToDomain() users.Domain {
	return users.Domain{
		Name:     ui.Name,
		Email:    ui.Email,
		Password: ui.Password,
		Address:  ui.Address,
		Phone:    ui.Phone,
		// TransactionDetails: tr.TransacDetail,
	}
}
