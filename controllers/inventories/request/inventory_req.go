package request

import "miniProject/business/inventories"

type InventoryInsert struct {
	Name        string `json:"name"`
	Stock       int    `json:"stock"`
	MeasureType string `json:"measure_type"`
	UnitPrice   int    `json:"unit_price"`
}

func (inv *InventoryInsert) ToDomain() inventories.Domain {
	return inventories.Domain{
		Name:        inv.Name,
		Stock:       inv.Stock,
		MeasureType: inv.MeasureType,
		UnitPrice:   inv.UnitPrice,
	}
}
