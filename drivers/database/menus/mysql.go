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
	var CurrentMenu []Menus
	result := DB.Conn.Find(&CurrentMenu)
	if result.Error != nil {
		return []menus.Domain{}, result.Error
	}
	return ToDomains(CurrentMenu), nil
}

func (DB *MysqlMenuRepository) GetMenuByID(ctx context.Context, id uint) (menus.Domain, error) {
	var Currentmenu Menus
	result := DB.Conn.Where("id = ?", id).First(&Currentmenu)
	if result.Error != nil {
		return menus.Domain{}, result.Error
	}
	return Currentmenu.ToDomain(), nil
}

func (DB *MysqlMenuRepository) CreateMenu(ctx context.Context, data menus.Domain) (menus.Domain, error) {

	InsertMenu := FromDomain(data)
	req_result := DB.Conn.Create(&InsertMenu)
	if req_result.Error != nil {
		return menus.Domain{}, req_result.Error
	}
	resp_result, err := DB.GetMenuByID(ctx, InsertMenu.ID)
	if err != nil {
		return menus.Domain{}, err
	}
	return resp_result, nil
}

func (DB *MysqlMenuRepository) UpdateMenu(ctx context.Context, data menus.Domain, id uint) (menus.Domain, error) {
	Currentmenu := FromDomain(data)
	result := DB.Conn.Where("id = ?", id).Updates(Currentmenu).First(&Currentmenu)
	if result.Error != nil {
		return menus.Domain{}, result.Error
	}
	return Currentmenu.ToDomain(), nil
}

func (DB *MysqlMenuRepository) DeleteMenu(ctx context.Context, id uint) (menus.Domain, error) {
	var Currentmenu Menus
	result := DB.Conn.Where("id = ?", id).Delete(&Currentmenu)
	if result.Error != nil {
		return menus.Domain{}, result.Error
	}
	return Currentmenu.ToDomain(), nil
}
