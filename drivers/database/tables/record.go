package tables

import (
	"miniProject/business/tables"

	"gorm.io/gorm"
)

type Tables struct {
	gorm.Model
	Capacity int
	Status   bool
}

func (table *Tables) ToDomain() tables.Domain {
	return tables.Domain{

		ID:        table.ID,
		Capacity:  table.Capacity,
		Status:    table.Status,
		CreatedAt: table.CreatedAt,
		UpdatedAt: table.CreatedAt,
	}
}

func FromDomain(domain tables.Domain) Tables {
	return Tables{

		Capacity: domain.Capacity,
		Status:   domain.Status,
	}
}

func ToDomains(u []Tables) []tables.Domain {
	var Domains []tables.Domain

	for _, val := range u {
		Domains = append(Domains, val.ToDomain())
	}
	return Domains
}
