package responses

type TransactionDetailInsert struct {
	TransactionID int    `json:"transaction_id"`
	Menu          string `json:"menu"`
	Quantity      int    `json:"quantity"`
}
