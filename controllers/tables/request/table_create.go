package request

import "miniProject/business/tables"

type TableInsert struct {
	Capacity int  `json:"capacity"`
	Status   bool `json:"status"`
}

func (tb *TableInsert) ToDomain() tables.Domain {
	return tables.Domain{
		Capacity: tb.Capacity,
		Status:   tb.Status,
	}
}
