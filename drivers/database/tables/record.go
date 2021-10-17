package tables

import (
	"miniProject/business/tables"
	"time"
)

type Tables struct {
	ID        int `gorm:"primaryKey"`
	Capacity  int
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
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

func ToListDomain(data []Tables) (result []tables.Domain) {
	result = []tables.Domain{}
	for _, val := range data {
		result = append(result, val.ToDomain())
	}
	return result
}
