package transaction_details

import (
	trDetails "miniProject/business/transaction_details"
	"miniProject/drivers/database/menus"
	"miniProject/drivers/database/transactions"
)

type TransactionDetails struct {
	ID            int                       `gorm:"primaryKey"`
	TransactionID int                       `gorm:"not null"`
	Transaction   transactions.Transactions `gorm:"foreignKey:TransactionID"`
	MenuID        int
	Menu          menus.Menus `gorm:"foreignKey:MenuID"`
	Quantity      int         `gorm:"default:0"`
}

func (trDetail *TransactionDetails) ToDomain() trDetails.Domain {
	return trDetails.Domain{
		ID:            trDetail.ID,
		TransactionID: trDetail.TransactionID,
		Transaction:   trDetail.Transaction.ToDomain(),
		MenuID:        trDetail.MenuID,
		Menu:          trDetail.Menu.Name,
		Menus:         trDetail.Menu.ToDomain(),
		Quantity:      trDetail.Quantity,
		Name:          trDetail.Transaction.User.Name,
	}
}

func FromDomain(domain trDetails.Domain) TransactionDetails {
	return TransactionDetails{
		TransactionID: domain.TransactionID,
		MenuID:        domain.MenuID,
		Quantity:      domain.Quantity,
	}
}

func ToDomains(u []TransactionDetails) []trDetails.Domain {
	var Domains []trDetails.Domain

	for _, val := range u {
		Domains = append(Domains, val.ToDomain())
	}
	return Domains
}

func ToListDomain(data []TransactionDetails) (result []trDetails.Domain) {
	result = []trDetails.Domain{}
	for _, val := range data {
		result = append(result, val.ToDomain())
	}
	return result
}
