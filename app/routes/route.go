package routes

import (
	"miniProject/controllers/users"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	UserController users.UserController
}

func (cl *ControllerList) RouteRegister(e echo.Echo) {
	e.POST("users/login", cl.UserController.Login)
	e.GET("users", cl.UserController.GetUserController)
	e.GET("user", cl.UserController.GetUserByIDController)
	e.POST("user", cl.UserController.CreateUserController)
	e.PUT("user/:id", cl.UserController.UpdateUserController)
	e.DELETE("user", cl.UserController.DeleteUserController)
}
