package responses

import "miniProject/business/menus"

type UserMenu struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"`
}

func UserMenuFromDomain(domain menus.Domain) UserMenu {
	return UserMenu{
		Name:     domain.Name,
		Category: domain.Category,
		Price:    domain.Price,
	}
}
