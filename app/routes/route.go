package routes

import (
	"miniProject/controllers/inventories"
	"miniProject/controllers/menus"
	"miniProject/controllers/tables"
	"miniProject/controllers/transaction_details"
	"miniProject/controllers/transactions"
	"miniProject/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	UserController              users.UserController
	MenuController              menus.MenuController
	TableController             tables.TableController
	InventoryController         inventories.InventoryController
	TransactionController       transactions.TransactionController
	TransactionDetailController transaction_details.TransactionDetailController
	JWTConfig                   middleware.JWTConfig
}

func (cl *ControllerList) RouteRegister(e echo.Echo) {
	e.Use(middleware.Logger())

	//Users
	e.POST("users/login", cl.UserController.Login)
	e.GET("users", cl.UserController.GetUserController, middleware.JWTWithConfig(cl.JWTConfig))
	e.GET("user", cl.UserController.GetUserByIDController)
	e.POST("user", cl.UserController.CreateUserController)
	e.PUT("user/:id", cl.UserController.UpdateUserController)
	e.DELETE("user", cl.UserController.DeleteUserController)

	//Inventories
	e.GET("inventories", cl.InventoryController.GetInventoriesController)
	e.GET("inventory", cl.InventoryController.GetInventoryByIDController)
	e.POST("inventory", cl.InventoryController.CreateInventoryController)
	e.PUT("inventory/:id", cl.InventoryController.UpdateInventoryController)
	e.DELETE("inventory", cl.InventoryController.DeleteInventoryController)

	//Menus
	e.GET("menus", cl.MenuController.GetMenusController)
	e.GET("menu", cl.MenuController.GetMenuByIDController)
	e.POST("menu", cl.MenuController.CreateMenuController)
	e.PUT("menu/:id", cl.MenuController.UpdateMenuController)
	e.DELETE("menu", cl.MenuController.DeleteMenuController)

	//Tables
	e.GET("tables", cl.TableController.GetTablesController)
	e.GET("table", cl.TableController.GetTableByIDController)
	e.POST("table", cl.TableController.CreateTableController)
	e.PUT("table/:id", cl.TableController.UpdateTableController)
	e.DELETE("table", cl.TableController.DeleteTableController)

	//Transactions
	e.GET("transaction", cl.TransactionController.GetTransactionByIDController)
	e.POST("transaction", cl.TransactionController.CreateTransactionController)
	e.PUT("transaction/:id", cl.TransactionController.UpdateTransactionController)

	//TransactionDetails
	e.GET("transaction-details", cl.TransactionDetailController.GetTransactionDetailsController)
	e.GET("transaction-detail", cl.TransactionDetailController.GetTransactionDetailByIDController)
	e.PUT("transaction-detail/:id", cl.TransactionDetailController.UpdateTransactionDetailController)

	//API
	e.GET("getapimenu/:name", cl.MenuController.GetAPI)
	e.POST("buy", cl.TransactionDetailController.CreateTransactionDetailController)
}
