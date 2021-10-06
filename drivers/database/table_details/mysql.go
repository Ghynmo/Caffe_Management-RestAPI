package table_details

import (
	"context"
	"miniProject/business/table_details"

	"gorm.io/gorm"
)

type MysqlTable_detailRepository struct {
	Conn *gorm.DB
}

func NewMysqlTableDetailRepository(conn *gorm.DB) table_details.Repository {
	return &MysqlTable_detailRepository{
		Conn: conn,
	}
}

func (DB *MysqlTable_detailRepository) GetTableDetails(ctx context.Context) ([]table_details.Domain, error) {
	var CurrentTable_detail []TableDetails
	result := DB.Conn.Find(&CurrentTable_detail)
	if result.Error != nil {
		return []table_details.Domain{}, result.Error
	}
	return ToDomains(CurrentTable_detail), nil
}

func (DB *MysqlTable_detailRepository) GetTableDetailByID(ctx context.Context, id uint) (table_details.Domain, error) {
	var Currenttable_detail TableDetails
	result := DB.Conn.Where("id = ?", id).First(&Currenttable_detail)
	if result.Error != nil {
		return table_details.Domain{}, result.Error
	}
	return Currenttable_detail.ToDomain(), nil
}

func (DB *MysqlTable_detailRepository) CreateTableDetail(ctx context.Context, data table_details.Domain) (table_details.Domain, error) {
	InsertTable_detail := FromDomain(data)
	result := DB.Conn.Create(&InsertTable_detail)
	if result.Error != nil {
		return table_details.Domain{}, result.Error
	}
	return InsertTable_detail.ToDomain(), nil
}

func (DB *MysqlTable_detailRepository) UpdateTableDetail(ctx context.Context, data table_details.Domain, id uint) (table_details.Domain, error) {
	Currenttable_detail := FromDomain(data)
	result := DB.Conn.Where("id = ?", id).Updates(Currenttable_detail).First(&Currenttable_detail)
	if result.Error != nil {
		return table_details.Domain{}, result.Error
	}
	return Currenttable_detail.ToDomain(), nil
}
