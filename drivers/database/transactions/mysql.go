package transactions

import (
	"context"
	"fmt"
	"miniProject/business/transactions"

	"gorm.io/gorm"
)

type MysqlTransactionRepository struct {
	Conn *gorm.DB
}

func NewMysqlTransactionRepository(conn *gorm.DB) transactions.Repository {
	return &MysqlTransactionRepository{
		Conn: conn,
	}
}

func (DB *MysqlTransactionRepository) BuyTransaction(ctx context.Context, data transactions.Domain) (transactions.Domain, error) {
	CurrentTransaction := FromDomain(data)
	fmt.Println(CurrentTransaction)
	result := DB.Conn.Create(&CurrentTransaction)
	if result.Error != nil {
		return transactions.Domain{}, result.Error
	}
	return CurrentTransaction.ToDomain(), nil
}

func (DB *MysqlTransactionRepository) GetTransactionByID(ctx context.Context, id uint) (transactions.Domain, error) {
	var Currenttransaction Transactions
	result := DB.Conn.Where("id = ?", id).First(&Currenttransaction)
	if result.Error != nil {
		return transactions.Domain{}, result.Error
	}
	return Currenttransaction.ToDomain(), nil
}

func (DB *MysqlTransactionRepository) CreateTransaction(ctx context.Context, data transactions.Domain) (transactions.Domain, error) {
	InsertTransaction := FromDomain(data)
	result := DB.Conn.Create(&InsertTransaction)
	if result.Error != nil {
		return transactions.Domain{}, result.Error
	}
	return InsertTransaction.ToDomain(), nil
}

func (DB *MysqlTransactionRepository) UpdateTransaction(ctx context.Context, data transactions.Domain, id uint) (transactions.Domain, error) {
	Currenttransaction := FromDomain(data)
	result := DB.Conn.Where("id = ?", id).Updates(Currenttransaction).First(&Currenttransaction)
	if result.Error != nil {
		return transactions.Domain{}, result.Error
	}
	return Currenttransaction.ToDomain(), nil
}
