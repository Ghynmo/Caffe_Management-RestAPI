package recipes

import "time"

type Domain struct {
	ID                int
	Name              string
	Step              string
	EstimateInMinutes int
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
