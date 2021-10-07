package main

import (
	"log"
	_middleware "miniProject/app/middlewares"
	"miniProject/app/routes"
	_mysqlDriver "miniProject/drivers/mysql"
	"time"

	_userUseCase "miniProject/business/users"
	_userController "miniProject/controllers/users"
	_userRepository "miniProject/drivers/database/users"

	_inventoryUseCase "miniProject/business/inventories"
	_inventoryController "miniProject/controllers/inventories"
	_inventoryRepository "miniProject/drivers/database/inventories"

	_menuUseCase "miniProject/business/menus"
	_menuController "miniProject/controllers/menus"
	_menuRepository "miniProject/drivers/database/menus"

	_tableUseCase "miniProject/business/tables"
	_tableController "miniProject/controllers/tables"
	_tableRepository "miniProject/drivers/database/tables"

	_transactionUseCase "miniProject/business/transactions"
	_transactionController "miniProject/controllers/transactions"
	_transactionRepository "miniProject/drivers/database/transactions"

	_transaction_detailUseCase "miniProject/business/transaction_details"
	_transaction_detailController "miniProject/controllers/transaction_details"
	_transaction_detailRepository "miniProject/drivers/database/transaction_details"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`app/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func DbMigrate(db *gorm.DB) {
	db.AutoMigrate(&_userRepository.Users{})
	db.AutoMigrate(&_transactionRepository.Transactions{})
	db.AutoMigrate(&_transaction_detailRepository.TransactionDetails{})
	db.AutoMigrate(&_inventoryRepository.Inventories{})
	db.AutoMigrate(&_menuRepository.Menus{})
	db.AutoMigrate(&_tableRepository.Tables{})

}

func main() {

	//Koneksi ke Database
	configDB := _mysqlDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

	configJWT := _middleware.ConfigJwt{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	Conn := configDB.InitialDB()
	DbMigrate(Conn)

	e := echo.New()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	userRepository := _userRepository.NewMysqlUserRepository(Conn)
	userUseCase := _userUseCase.NewUserUseCase(userRepository, timeoutContext, configJWT)
	userController := _userController.NewUserController(userUseCase)

	menuRepository := _menuRepository.NewMysqlMenuRepository(Conn)
	menuUseCase := _menuUseCase.NewMenuUseCase(menuRepository, timeoutContext)
	menuController := _menuController.NewMenuController(menuUseCase)

	transactionRepository := _transactionRepository.NewMysqlTransactionRepository(Conn)
	transactionUseCase := _transactionUseCase.NewTransactionUseCase(transactionRepository, timeoutContext)
	transactionController := _transactionController.NewTransactionController(transactionUseCase)

	transaction_detailRepository := _transaction_detailRepository.NewMysqlTransactionDetailRepository(Conn)
	transaction_detailUseCase := _transaction_detailUseCase.NewTransactionDetailUseCase(transaction_detailRepository, timeoutContext)
	transaction_detailController := _transaction_detailController.NewTransactionDetailController(transaction_detailUseCase)

	inventoryRepository := _inventoryRepository.NewMysqlInventoryRepository(Conn)
	inventoryUseCase := _inventoryUseCase.NewInventoryUseCase(inventoryRepository, timeoutContext)
	inventoryController := _inventoryController.NewInventoryController(inventoryUseCase)

	tableRepository := _tableRepository.NewMysqlTableRepository(Conn)
	tableUseCase := _tableUseCase.NewTableUseCase(tableRepository, timeoutContext)
	tableController := _tableController.NewTableController(tableUseCase)

	routesInit := routes.ControllerList{
		JWTConfig:                   configJWT.Init(),
		UserController:              *userController,
		InventoryController:         *inventoryController,
		MenuController:              *menuController,
		TableController:             *tableController,
		TransactionController:       *transactionController,
		TransactionDetailController: *transaction_detailController,
	}

	routesInit.RouteRegister(*e)

	log.Fatal(e.Start((viper.GetString("server.address"))))
}
