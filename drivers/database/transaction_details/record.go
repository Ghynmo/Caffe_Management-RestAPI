package transaction_details

import (
	trDetails "miniProject/business/transaction_details"

	"gorm.io/gorm"
)

type TransactionDetails struct {
	gorm.Model
	TransactionID uint
	MenuID        uint
	Quantity      int
}

func (trDetail *TransactionDetails) ToDomain() trDetails.Domain {
	return trDetails.Domain{
		ID:            trDetail.ID,
		TransactionID: trDetail.TransactionID,
		MenuID:        trDetail.MenuID,
		Quantity:      trDetail.Quantity,
		CreatedAt:     trDetail.CreatedAt,
		UpdatedAt:     trDetail.UpdatedAt,
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
