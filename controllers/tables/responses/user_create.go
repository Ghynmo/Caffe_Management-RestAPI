package responses

import "miniProject/business/tables"

type TableInsert struct {
	Capacity int  `json:"capacity"`
	Status   bool `json:"status"`
}

func CreateFromDomain(domain tables.Domain) TableInsert {
	return TableInsert{
		Capacity: domain.Capacity,
		Status:   domain.Status,
	}
}
