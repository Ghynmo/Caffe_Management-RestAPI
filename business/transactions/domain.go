package transactions

import "time"

type Domain struct {
	ID            int
	UserID        int
	TableDetailID int
	Payment       int
	Price         int
	Status        bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
