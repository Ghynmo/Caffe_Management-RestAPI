package responses

import "miniProject/business/tables"

type TableResponse struct {
	ID       int  `json:"id"`
	Capacity int  `json:"capacity"`
	Status   bool `json:"status"`
}

func FromDomain(domain tables.Domain) TableResponse {
	return TableResponse{
		ID:       domain.ID,
		Capacity: domain.Capacity,
		Status:   domain.Status,
	}
}

func FromListDomain(data []tables.Domain) (result []TableResponse) {
	result = []TableResponse{}
	for _, val := range data {
		result = append(result, FromDomain(val))
	}
	return result
}
