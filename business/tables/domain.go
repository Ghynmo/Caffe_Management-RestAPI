package tables

import "time"

type Domain struct {
	ID        int
	Capacity  int
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
