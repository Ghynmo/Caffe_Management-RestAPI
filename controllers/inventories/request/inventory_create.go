package request

type InventoryInsert struct {
	Name        string `json:"name"`
	Stock       int    `json:"stock"`
	MeasureType string `json:"measure_type"`
	UnitPrice   int    `json:"unit_price"`
}
