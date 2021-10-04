package users

import (
	"context"
	"miniProject/business/users"

	"gorm.io/gorm"
)

type MysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(conn *gorm.DB) users.Repository {
	return &MysqlUserRepository{
		Conn: conn,
	}
}

func (DB *MysqlUserRepository) Login(ctx context.Context, email string, password string) (users.Domain, error) {
	var CurrentUser Users
	result := DB.Conn.First(&CurrentUser, "email = ? AND password = ?", email, password)
	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return CurrentUser.ToDomain(), nil
}

func (DB *MysqlUserRepository) GetUser(ctx context.Context) ([]users.Domain, error) {
	var CurrentUser []Users
	result := DB.Conn.Find(&CurrentUser)
	if result.Error != nil {
		return []users.Domain{}, result.Error
	}
	return ToDomains(CurrentUser), nil
}

func (DB *MysqlUserRepository) GetUserByID(ctx context.Context, id int) (users.Domain, error) {
	var Currentuser Users
	result := DB.Conn.Where("id = ?", id).First(&Currentuser)
	if result.Error != nil {
		return users.Domain{}, result.Error
	}
	return Currentuser.ToDomain(), nil
}

func (DB *MysqlUserRepository) CreateUser(ctx context.Context, data users.Domain) (users.Domain, error) {
	InsertUser := FromDomain(data)
	result := DB.Conn.Create(&InsertUser)
	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return InsertUser.ToDomain(), nil
}

func (DB *MysqlUserRepository) UpdateUser(ctx context.Context, data users.Domain, id int) (users.Domain, error) {
	Currentuser := FromDomain(data)
	result := DB.Conn.Where("id = ?", id).Updates(Currentuser).First(&Currentuser)
	if result.Error != nil {
		return users.Domain{}, result.Error
	}
	return Currentuser.ToDomain(), nil
}

func (DB *MysqlUserRepository) DeleteUser(ctx context.Context, id int) (users.Domain, error) {
	var Currentuser Users
	result := DB.Conn.Where("id = ?", id).Delete(&Currentuser)
	if result.Error != nil {
		return users.Domain{}, result.Error
	}
	return Currentuser.ToDomain(), nil
}
