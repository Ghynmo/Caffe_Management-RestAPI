package menus

import (
	"miniProject/business/menus"
	"time"
)

type Menus struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"unique;not null"`
	Category  string
	Price     int `gorm:"default:0"`
	Stock     int `gorm:"default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (menu *Menus) ToDomain() menus.Domain {
	return menus.Domain{

		ID:        int(menu.ID),
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
