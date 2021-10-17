package transactions

import (
	"miniProject/business/transactions"
	"miniProject/drivers/database/tables"
	"miniProject/drivers/database/users"
	"time"
)

type Transactions struct {
	ID        int             `gorm:"primaryKey"`
	UserID    int             `gorm:"not null"`
	Price     int             `gorm:"default:0"`
	Status    bool            `gorm:"default:false"`
	User      users.Users     `gorm:"foreignKey:UserID"`
	Table     []tables.Tables `gorm:"many2many:transaction_tablenum"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (transaction *Transactions) ToDomain() transactions.Domain {
	return transactions.Domain{
		ID:        transaction.ID,
		UserID:    transaction.UserID,
		Price:     transaction.Price,
		Status:    transaction.Status,
		User:      transaction.User.ToDomain(),
		CreatedAt: transaction.CreatedAt,
		UpdatedAt: transaction.CreatedAt,
		// TransactionDetails: transaction_details.ToListDomain(transaction.TransactionsDetails),
	}
}

func FromDomain(domain transactions.Domain) Transactions {
	return Transactions{
		UserID: domain.UserID,
		Price:  domain.Price,
		Status: domain.Status,
	}
}

func ToDomains(u []Transactions) []transactions.Domain {
	var Domains []transactions.Domain

	for _, val := range u {
		Domains = append(Domains, val.ToDomain())
	}
	return Domains
}

// func (trd *Transactions) ToTrDetails() transaction_details.Domain {
// 	return transaction_details.Domain{
// 		TransactionID:                 uint(transaction.ID),
// 		TransactionDetails: transaction_details.ToListDomain(transaction.TransactionsDetails),
// 	}
// }
