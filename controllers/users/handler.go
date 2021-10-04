package users

import (
	"miniProject/business/users"
	"miniProject/controllers"
	"miniProject/controllers/users/request"
	"miniProject/controllers/users/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserUseCase users.UseCase
}

func NewUserController(userUseCase users.UseCase) *UserController {
	return &UserController{
		UserUseCase: userUseCase,
	}
}

func (handler UserController) Login(c echo.Context) error {
	userLogin := request.UserLogin{}
	c.Bind(&userLogin)

	ctx := c.Request().Context()
	user, err := handler.UserUseCase.Login(ctx, userLogin.Email, userLogin.Password)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.FromDomain(user))
}

func (handler UserController) GetUserController(c echo.Context) error {

	ctx := c.Request().Context()
	user, err := handler.UserUseCase.GetUserController(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, user)
}

func (handler UserController) GetUserByIDController(c echo.Context) error {

	id, _ := strconv.Atoi(c.QueryParam("id"))

	ctx := c.Request().Context()

	user, err := handler.UserUseCase.GetUserByIDController(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, user)
}

func (handler UserController) CreateUserController(c echo.Context) error {
	userInsert := users.Domain{}
	c.Bind(&userInsert)

	ctx := c.Request().Context()

	user, err := handler.UserUseCase.CreateUserController(ctx, userInsert)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.CreateFromDomain(user))
}
