package transaction_details

import (
	"context"
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

func (DB *MysqlTransactionDetailRepository) GetTransactionDetail(ctx context.Context) ([]transaction_details.Domain, error) {
	return []transaction_details.Domain{}, nil

	// var CurrentTransactionDetail []TransactionDetail
	// result := DB.Conn.Find(&CurrentTransactionDetail)
	// if result.Error != nil {
	// 	return []transaction_details.Domain{}, result.Error
	// }
	// return ToDomains(CurrentTransactionDetail), nil
}

func (DB *MysqlTransactionDetailRepository) GetTransactionDetailByID(ctx context.Context, id int) (transaction_details.Domain, error) {
	return transaction_details.Domain{}, nil

	// var Currenttransaction TransactionDetail
	// result := DB.Conn.Where("id = ?", id).First(&Currenttransaction)
	// if result.Error != nil {
	// 	return transaction_details.Domain{}, result.Error
	// }
	// return Currenttransaction.ToDomain(), nil
}

func (DB *MysqlTransactionDetailRepository) CreateTransactionDetail(ctx context.Context, data transaction_details.Domain) (transaction_details.Domain, error) {
	return transaction_details.Domain{}, nil

	// InsertTransactionDetail := FromDomain(data)
	// result := DB.Conn.Create(&InsertTransactionDetail)
	// if result.Error != nil {
	// 	return transaction_details.Domain{}, result.Error
	// }

	// return InsertTransactionDetail.ToDomain(), nil
}

func (DB *MysqlTransactionDetailRepository) UpdateTransactionDetail(ctx context.Context, data transaction_details.Domain, id int) (transaction_details.Domain, error) {
	return transaction_details.Domain{}, nil

	// Currenttransaction := FromDomain(data)
	// result := DB.Conn.Where("id = ?", id).Updates(Currenttransaction).First(&Currenttransaction)
	// if result.Error != nil {
	// 	return transaction_details.Domain{}, result.Error
	// }
	// return Currenttransaction.ToDomain(), nil
}

func (DB *MysqlTransactionDetailRepository) DeleteTransactionDetail(ctx context.Context, id int) (transaction_details.Domain, error) {
	return transaction_details.Domain{}, nil

	// var Currenttransaction TransactionDetail
	// result := DB.Conn.Where("id = ?", id).Delete(&Currenttransaction)
	// if result.Error != nil {
	// 	return transaction_details.Domain{}, result.Error
	// }
	// return Currenttransaction.ToDomain(), nil
}
