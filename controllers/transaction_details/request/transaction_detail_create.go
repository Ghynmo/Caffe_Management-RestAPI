package request

type TransactionDetailInsert []struct {
	TransactionID int    `json:"transaction_id"`
	Menu          string `json:"menu"`
	Quantity      int    `json:"quantity"`
}

// func (tr *TransactionDetailInsert) ToDomain() transaction_details.Domain {
// 	return transaction_details.Domain{
// 		TransactionDetailID: tr.TransactionDetailID,
// 		Menu:                tr.Menu,
// 		Quantity:            tr.Quantity,
// 	}
// }
