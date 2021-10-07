package inventories

import (
	"miniProject/business/inventories"

	"gorm.io/gorm"
)

type Inventories struct {
	gorm.Model
	Name        string
	Stock       int
	MeasureType string
	UnitPrice   int
}

func (inventory *Inventories) ToDomain() inventories.Domain {
	return inventories.Domain{
		ID:          inventory.ID,
		Name:        inventory.Name,
		Stock:       inventory.Stock,
		MeasureType: inventory.MeasureType,
		UnitPrice:   inventory.UnitPrice,
		CreatedAt:   inventory.CreatedAt,
		UpdatedAt:   inventory.CreatedAt,
	}
}

func FromDomain(domain inventories.Domain) Inventories {
	return Inventories{
		Name:        domain.Name,
		Stock:       domain.Stock,
		MeasureType: domain.MeasureType,
		UnitPrice:   domain.UnitPrice,
	}
}

func ToDomains(u []Inventories) []inventories.Domain {
	var Domains []inventories.Domain

	for _, val := range u {
		Domains = append(Domains, val.ToDomain())
	}
	return Domains
}
