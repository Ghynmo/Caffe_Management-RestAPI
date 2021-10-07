package responses

import "miniProject/business/menus"

type MenuInsert struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"`
	Stock    int    `json:"stock"`
}

func CreateFromDomain(domain menus.Domain) MenuInsert {
	return MenuInsert{
		Name:     domain.Name,
		Category: domain.Category,
		Price:    domain.Price,
		Stock:    domain.Stock,
	}
}
