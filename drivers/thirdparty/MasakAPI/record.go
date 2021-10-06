package MasakAPI

import (
	"miniProject/business/masak_openapi"
)

type Name struct {
	Name              string `json:"title"`
	Step              string `json:"step"`
	EstimateInMinutes string `json:"times"`
}

func (rec *Name) ToDomain() masak_openapi.Domain {
	return masak_openapi.Domain{
		Name:              rec.Name,
		Step:              rec.Step,
		EstimateInMinutes: rec.EstimateInMinutes,
	}
}
