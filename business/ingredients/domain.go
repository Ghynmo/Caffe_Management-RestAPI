package ingredients

import "time"

type Domain struct {
	RecipeID   int
	MaterialID int
	Quantity   int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
