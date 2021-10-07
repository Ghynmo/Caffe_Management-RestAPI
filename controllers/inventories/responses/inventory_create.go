package responses

import "miniProject/business/inventories"

type InventoryInsert struct {
	Name        string `json:"name"`
	Stock       int    `json:"stock"`
	MeasureType string `json:"measure_type"`
	UnitPrice   int    `json:"unit_price"`
}

func CreateFromDomain(domain inventories.Domain) InventoryInsert {
	return InventoryInsert{
		Name:        domain.Name,
		Stock:       domain.Stock,
		MeasureType: domain.MeasureType,
		UnitPrice:   domain.UnitPrice,
	}
}
