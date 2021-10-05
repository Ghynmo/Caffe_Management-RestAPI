package inventories

import (
	"context"
	"miniProject/business/inventories"

	"gorm.io/gorm"
)

type MysqlInventoryRepository struct {
	Conn *gorm.DB
}

func NewMysqlInventoryRepository(conn *gorm.DB) inventories.Repository {
	return &MysqlInventoryRepository{
		Conn: conn,
	}
}

func (DB *MysqlInventoryRepository) GetInventory(ctx context.Context) ([]inventories.Domain, error) {
	return []inventories.Domain{}, nil

	// var CurrentInventory []Inventory
	// result := DB.Conn.Find(&CurrentInventory)
	// if result.Error != nil {
	// 	return []inventories.Domain{}, result.Error
	// }
	// return ToDomains(CurrentInventory), nil
}

func (DB *MysqlInventoryRepository) GetInventoryByID(ctx context.Context, id int) (inventories.Domain, error) {
	return inventories.Domain{}, nil

	// var Currentinventory Inventory
	// result := DB.Conn.Where("id = ?", id).First(&Currentinventory)
	// if result.Error != nil {
	// 	return inventories.Domain{}, result.Error
	// }
	// return Currentinventory.ToDomain(), nil
}

func (DB *MysqlInventoryRepository) CreateInventory(ctx context.Context, data inventories.Domain) (inventories.Domain, error) {
	return inventories.Domain{}, nil

	// InsertInventory := FromDomain(data)
	// result := DB.Conn.Create(&InsertInventory)
	// if result.Error != nil {
	// 	return inventories.Domain{}, result.Error
	// }

	// return InsertInventory.ToDomain(), nil
}

func (DB *MysqlInventoryRepository) UpdateInventory(ctx context.Context, data inventories.Domain, id int) (inventories.Domain, error) {
	return inventories.Domain{}, nil

	// Currentinventory := FromDomain(data)
	// result := DB.Conn.Where("id = ?", id).Updates(Currentinventory).First(&Currentinventory)
	// if result.Error != nil {
	// 	return inventories.Domain{}, result.Error
	// }
	// return Currentinventory.ToDomain(), nil
}

func (DB *MysqlInventoryRepository) DeleteInventory(ctx context.Context, id int) (inventories.Domain, error) {
	return inventories.Domain{}, nil

	// var Currentinventory Inventory
	// result := DB.Conn.Where("id = ?", id).Delete(&Currentinventory)
	// if result.Error != nil {
	// 	return inventories.Domain{}, result.Error
	// }
	// return Currentinventory.ToDomain(), nil
}
