package responses

import "miniProject/business/menus"

type MenuResponse struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"`
	Stock    int    `json:"stock"`
}

func FromDomain(domain menus.Domain) MenuResponse {
	return MenuResponse{
		Name:     domain.Name,
		Category: domain.Category,
		Price:    domain.Price,
		Stock:    domain.Stock,
	}
}

func FromListDomain(data []menus.Domain) (result []MenuResponse) {
	result = []MenuResponse{}
	for _, val := range data {
		result = append(result, FromDomain(val))
	}
	return result
}
