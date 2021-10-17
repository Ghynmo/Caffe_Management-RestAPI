package request

import "miniProject/business/menus"

type MenuInsert struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"`
	Stock    int    `json:"stock"`
}

func (inv *MenuInsert) ToDomain() menus.Domain {
	return menus.Domain{
		Name:     inv.Name,
		Category: inv.Category,
		Price:    inv.Price,
		Stock:    inv.Stock,
	}
}
