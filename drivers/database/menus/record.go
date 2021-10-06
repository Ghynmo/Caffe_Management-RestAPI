package menus

import (
	"miniProject/business/menus"

	"gorm.io/gorm"
)

type Menus struct {
	gorm.Model
	Name     string
	Category string
	Price    int
	Stock    int
}

func (menu *Menus) ToDomain() menus.Domain {
	return menus.Domain{

		ID:        uint(menu.ID),
		Name:      menu.Name,
		Category:  menu.Category,
		Price:     menu.Price,
		Stock:     menu.Stock,
		CreatedAt: menu.CreatedAt,
		UpdatedAt: menu.CreatedAt,
	}
}

func FromDomain(domain menus.Domain) Menus {
	return Menus{
		Name:     domain.Name,
		Category: domain.Category,
		Price:    domain.Price,
		Stock:    domain.Stock,
	}
}

func ToDomains(u []Menus) []menus.Domain {
	var Domains []menus.Domain

	for _, val := range u {
		Domains = append(Domains, val.ToDomain())
	}
	return Domains
}
