package transaction_details

import (
	trDetails "miniProject/business/transaction_details"

	"gorm.io/gorm"
)

type TransactionDetails struct {
	gorm.Model
	TransactionID int //`gorm:"uniqueIndex:idx_table_details;type:BIGINT(255)"`
	Menu          string
	Quantity      int
}

func (trDetail *TransactionDetails) ToDomain() trDetails.Domain {
	return trDetails.Domain{
		TransactionID: trDetail.TransactionID,
		Menu:          trDetail.Menu,
		Quantity:      trDetail.Quantity,
	}
}

func FromDomain(domain trDetails.Domain) TransactionDetails {
	return TransactionDetails{
		TransactionID: domain.TransactionID,
		Menu:          domain.Menu,
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
