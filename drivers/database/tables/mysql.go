package tables

import (
	"context"
	"miniProject/business/tables"

	"gorm.io/gorm"
)

type MysqlTableRepository struct {
	Conn *gorm.DB
}

func NewMysqlTableRepository(conn *gorm.DB) tables.Repository {
	return &MysqlTableRepository{
		Conn: conn,
	}
}

func (DB *MysqlTableRepository) GetTables(ctx context.Context) ([]tables.Domain, error) {
	var CurrentTable []Tables
	result := DB.Conn.Find(&CurrentTable)
	if result.Error != nil {
		return []tables.Domain{}, result.Error
	}
	return ToDomains(CurrentTable), nil
}

func (DB *MysqlTableRepository) GetTableByID(ctx context.Context, id int) (tables.Domain, error) {
	var Currenttable Tables
	result := DB.Conn.Where("id = ?", id).First(&Currenttable)
	if result.Error != nil {
		return tables.Domain{}, result.Error
	}
	return Currenttable.ToDomain(), nil
}

func (DB *MysqlTableRepository) CreateTable(ctx context.Context, data tables.Domain) (tables.Domain, error) {
	InsertTable := FromDomain(data)
	result := DB.Conn.Create(&InsertTable)
	if result.Error != nil {
		return tables.Domain{}, result.Error
	}

	return InsertTable.ToDomain(), nil
}

func (DB *MysqlTableRepository) UpdateTable(ctx context.Context, data tables.Domain, id int) (tables.Domain, error) {
	Currenttable := FromDomain(data)
	result := DB.Conn.Where("id = ?", id).Updates(Currenttable).First(&Currenttable)
	if result.Error != nil {
		return tables.Domain{}, result.Error
	}
	return Currenttable.ToDomain(), nil
}

func (DB *MysqlTableRepository) DeleteTable(ctx context.Context, id int) (tables.Domain, error) {
	var Currenttable Tables
	result := DB.Conn.Where("id = ?", id).Delete(&Currenttable)
	if result.Error != nil {
		return tables.Domain{}, result.Error
	}
	return Currenttable.ToDomain(), nil
}
