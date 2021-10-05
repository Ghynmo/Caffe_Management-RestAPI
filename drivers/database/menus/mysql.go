package menus

import (
	"context"
	"miniProject/business/menus"

	"gorm.io/gorm"
)

type MysqlMenuRepository struct {
	Conn *gorm.DB
}

func NewMysqlMenuRepository(conn *gorm.DB) menus.Repository {
	return &MysqlMenuRepository{
		Conn: conn,
	}
}

func (DB *MysqlMenuRepository) GetMenu(ctx context.Context) ([]menus.Domain, error) {
	return []menus.Domain{}, nil

	// var CurrentMenu []Menu
	// result := DB.Conn.Find(&CurrentMenu)
	// if result.Error != nil {
	// 	return []menus.Domain{}, result.Error
	// }
	// return ToDomains(CurrentMenu), nil
}

func (DB *MysqlMenuRepository) GetMenuByID(ctx context.Context, id int) (menus.Domain, error) {
	return menus.Domain{}, nil

	// var Currentmenu Menu
	// result := DB.Conn.Where("id = ?", id).First(&Currentmenu)
	// if result.Error != nil {
	// 	return menus.Domain{}, result.Error
	// }
	// return Currentmenu.ToDomain(), nil
}

func (DB *MysqlMenuRepository) CreateMenu(ctx context.Context, data menus.Domain) (menus.Domain, error) {
	return menus.Domain{}, nil

	// InsertMenu := FromDomain(data)
	// result := DB.Conn.Create(&InsertMenu)
	// if result.Error != nil {
	// 	return menus.Domain{}, result.Error
	// }

	// return InsertMenu.ToDomain(), nil
}

func (DB *MysqlMenuRepository) UpdateMenu(ctx context.Context, data menus.Domain, id int) (menus.Domain, error) {
	return menus.Domain{}, nil

	// Currentmenu := FromDomain(data)
	// result := DB.Conn.Where("id = ?", id).Updates(Currentmenu).First(&Currentmenu)
	// if result.Error != nil {
	// 	return menus.Domain{}, result.Error
	// }
	// return Currentmenu.ToDomain(), nil
}

func (DB *MysqlMenuRepository) DeleteMenu(ctx context.Context, id int) (menus.Domain, error) {
	return menus.Domain{}, nil

	// var Currentmenu Menu
	// result := DB.Conn.Where("id = ?", id).Delete(&Currentmenu)
	// if result.Error != nil {
	// 	return menus.Domain{}, result.Error
	// }
	// return Currentmenu.ToDomain(), nil
}
