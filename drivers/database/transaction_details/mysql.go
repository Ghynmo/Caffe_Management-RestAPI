package transaction_details

import (
	"context"
	"fmt"
	"miniProject/business/transaction_details"

	"gorm.io/gorm"
)

type MysqlTransactionDetailRepository struct {
	Conn *gorm.DB
}

func NewMysqlTransactionDetailRepository(conn *gorm.DB) transaction_details.Repository {
	return &MysqlTransactionDetailRepository{
		Conn: conn,
	}
}

func (DB *MysqlTransactionDetailRepository) GetTransactionDetails(ctx context.Context) ([]transaction_details.Domain, error) {
	var CurrentTransactionDetail []TransactionDetails
	result := DB.Conn.Preload("Menu").Preload("Transaction.User").Find(&CurrentTransactionDetail)
	if result.Error != nil {
		return []transaction_details.Domain{}, result.Error
	}
	return ToDomains(CurrentTransactionDetail), nil
}

func (DB *MysqlTransactionDetailRepository) GetTransactionDetailByID(ctx context.Context, id int) (transaction_details.Domain, error) {
	var Currenttransaction TransactionDetails
	result := DB.Conn.Preload("Menu").Preload("Transaction.User").Where("id = ?", id).First(&Currenttransaction)
	if result.Error != nil {
		return transaction_details.Domain{}, result.Error
	}
	return Currenttransaction.ToDomain(), nil
}

func (DB *MysqlTransactionDetailRepository) CreateTransactionDetail(ctx context.Context, data transaction_details.Domain) (transaction_details.Domain, error) {
	InsertTransactionDetail := FromDomain(data)

	fmt.Println("After pass from Usecase, Conv to driver struct", InsertTransactionDetail)

	result := DB.Conn.Create(&InsertTransactionDetail)
	if result.Error != nil {
		return transaction_details.Domain{}, result.Error
	}

	resp_result := DB.Conn.Preload("Menu").Preload("Transaction.User").Where("id = ?", InsertTransactionDetail.ID).First(&InsertTransactionDetail)
	if result.Error != nil {
		return transaction_details.Domain{}, resp_result.Error
	}

	return InsertTransactionDetail.ToDomain(), nil
}

func (DB *MysqlTransactionDetailRepository) UpdateTransactionDetail(ctx context.Context, data transaction_details.Domain, id int) (transaction_details.Domain, error) {
	Currenttransaction := FromDomain(data)
	result := DB.Conn.Where("id = ?", id).Updates(Currenttransaction).First(&Currenttransaction)
	if result.Error != nil {
		return transaction_details.Domain{}, result.Error
	}
	return Currenttransaction.ToDomain(), nil
}
