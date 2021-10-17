package responses

import "miniProject/business/users"

type UserResponse struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

func FromDomain(domain users.Domain) UserResponse {
	return UserResponse{
		Name:    domain.Name,
		Email:   domain.Email,
		Address: domain.Address,
		Phone:   domain.Phone,
	}
}

func FromListDomain(data []users.Domain) (result []UserResponse) {
	result = []UserResponse{}
	for _, val := range data {
		result = append(result, FromDomain(val))
	}
	return result
}
