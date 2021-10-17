package responses

import "miniProject/business/inventories"

type InventoryResponse struct {
	Name        string `json:"name"`
	Stock       int    `json:"stock"`
	MeasureType string `json:"measure_type"`
	UnitPrice   int    `json:"unit_price"`
}

func FromDomain(domain inventories.Domain) InventoryResponse {
	return InventoryResponse{
		Name:        domain.Name,
		Stock:       domain.Stock,
		MeasureType: domain.MeasureType,
		UnitPrice:   domain.UnitPrice,
	}
}

func FromListDomain(data []inventories.Domain) (result []InventoryResponse) {
	result = []InventoryResponse{}
	for _, val := range data {
		result = append(result, FromDomain(val))
	}
	return result
}
