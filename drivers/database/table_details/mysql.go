package table_details

import (
	"context"
	"miniProject/business/table_details"

	"gorm.io/gorm"
)

type MysqlTableDetailRepository struct {
	Conn *gorm.DB
}

func NewMysqlTableDetailRepository(conn *gorm.DB) table_details.Repository {
	return &MysqlTableDetailRepository{
		Conn: conn,
	}
}

func (DB *MysqlTableDetailRepository) GetTableDetail(ctx context.Context) ([]table_details.Domain, error) {
	return []table_details.Domain{}, nil

	// var CurrentTableDetail []TableDetail
	// result := DB.Conn.Find(&CurrentTableDetail)
	// if result.Error != nil {
	// 	return []table_details.Domain{}, result.Error
	// }
	// return ToDomains(CurrentTableDetail), nil
}

func (DB *MysqlTableDetailRepository) GetTableDetailByID(ctx context.Context, id int) (table_details.Domain, error) {
	return table_details.Domain{}, nil

	// var Currenttransaction TableDetail
	// result := DB.Conn.Where("id = ?", id).First(&Currenttransaction)
	// if result.Error != nil {
	// 	return table_details.Domain{}, result.Error
	// }
	// return Currenttransaction.ToDomain(), nil
}

func (DB *MysqlTableDetailRepository) CreateTableDetail(ctx context.Context, data table_details.Domain) (table_details.Domain, error) {
	return table_details.Domain{}, nil

	// InsertTableDetail := FromDomain(data)
	// result := DB.Conn.Create(&InsertTableDetail)
	// if result.Error != nil {
	// 	return table_details.Domain{}, result.Error
	// }

	// return InsertTableDetail.ToDomain(), nil
}

func (DB *MysqlTableDetailRepository) UpdateTableDetail(ctx context.Context, data table_details.Domain, id int) (table_details.Domain, error) {
	return table_details.Domain{}, nil

	// Currenttransaction := FromDomain(data)
	// result := DB.Conn.Where("id = ?", id).Updates(Currenttransaction).First(&Currenttransaction)
	// if result.Error != nil {
	// 	return table_details.Domain{}, result.Error
	// }
	// return Currenttransaction.ToDomain(), nil
}

func (DB *MysqlTableDetailRepository) DeleteTableDetail(ctx context.Context, id int) (table_details.Domain, error) {
	return table_details.Domain{}, nil

	// var Currenttransaction TableDetail
	// result := DB.Conn.Where("id = ?", id).Delete(&Currenttransaction)
	// if result.Error != nil {
	// 	return table_details.Domain{}, result.Error
	// }
	// return Currenttransaction.ToDomain(), nil
}
