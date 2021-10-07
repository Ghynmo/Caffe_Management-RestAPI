package thirdparty

import "miniProject/business/menus"

type MasakApi struct {
	Results Results `json:"results"`
}

type Results struct {
	Name string `json:"title"`
}

func (menu *MasakApi) ToDomain() menus.Domain {
	return menus.Domain{
		Name: menu.Results.Name,
	}
}

// func ToDomains(u []MasakApi) []menus.Domain {
// 	var Domains []menus.Domain

// 	for _, val := range u {
// 		Domains = append(Domains, val.ToDomain())
// 	}
// 	return Domains
// }
