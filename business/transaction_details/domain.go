package transaction_details

import "time"

type Domain struct {
	ID            int
	TransactionID int
	MenuID        int
	Quantity      int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
