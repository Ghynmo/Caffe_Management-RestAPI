package request

type MenuInsert struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    string `json:"price"`
	Stock    string `json:"stock"`
}
