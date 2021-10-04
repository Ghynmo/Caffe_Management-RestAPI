package menus

import "time"

type Domain struct {
	ID          int
	RecipeID    int
	Name        string
	Category    string
	Price       int
	Stock       int
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
