package request

type TransactionDetailInsert struct {
	TransactionID uint `json:"transaction_id"`
	MenuID        uint `json:"menu_id"`
	Quantity      int  `json:"quantity"`
}
