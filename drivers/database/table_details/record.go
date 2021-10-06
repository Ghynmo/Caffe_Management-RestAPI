package table_details

import (
	"miniProject/business/table_details"

	"gorm.io/gorm"
)

type TableDetails struct {
	gorm.Model
	TotalCapacity int
	Status        bool
}

func (table_detail *TableDetails) ToDomain() table_details.Domain {
	return table_details.Domain{
		TotalCapacity: table_detail.TotalCapacity,
		Status:        table_detail.Status,
	}
}

func FromDomain(domain table_details.Domain) TableDetails {
	return TableDetails{
		TotalCapacity: domain.TotalCapacity,
		Status:        domain.Status,
	}
}

func ToDomains(u []TableDetails) []table_details.Domain {
	var Domains []table_details.Domain

	for _, val := range u {
		Domains = append(Domains, val.ToDomain())
	}
	return Domains
}
