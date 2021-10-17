package transactions

import (
	"context"
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

func (DB *MysqlTransactionRepository) GetTransaction(ctx context.Context) ([]transactions.Domain, error) {
	var CurrentTransaction []Transactions
	result := DB.Conn.Preload("User").Find(&CurrentTransaction)
	if result.Error != nil {
		return []transactions.Domain{}, result.Error
	}
	return ToDomains(CurrentTransaction), nil
}

func (DB *MysqlTransactionRepository) GetTransactionByID(ctx context.Context, id int) (transactions.Domain, error) {
	var Currenttransaction Transactions
	result := DB.Conn.Preload("User").Where("id = ?", id).First(&Currenttransaction)
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

	resp_result := DB.Conn.Preload("User").Where("id = ?", InsertTransaction.ID).First(&InsertTransaction)
	if result.Error != nil {
		return transactions.Domain{}, resp_result.Error
	}

	return InsertTransaction.ToDomain(), nil
}

func (DB *MysqlTransactionRepository) DeleteTransaction(ctx context.Context, id int) (transactions.Domain, error) {
	var Currenttransaction Transactions
	result := DB.Conn.Where("id = ?", id).Delete(&Currenttransaction)
	if result.Error != nil {
		return transactions.Domain{}, result.Error
	}
	return Currenttransaction.ToDomain(), nil
}
