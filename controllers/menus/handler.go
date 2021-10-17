package menus

import (
	"miniProject/business/menus"
	"miniProject/controllers"
	"miniProject/controllers/menus/request"
	"miniProject/controllers/menus/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MenuController struct {
	MenuUseCase menus.UseCase
}

func NewMenuController(menuUseCase menus.UseCase) *MenuController {
	return &MenuController{
		MenuUseCase: menuUseCase,
	}
}

func (handler MenuController) GetMenusController(c echo.Context) error {

	ctx := c.Request().Context()
	menu, err := handler.MenuUseCase.GetMenusController(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.FromListDomain(menu))
}

func (handler MenuController) GetMenuByIDController(c echo.Context) error {

	id, _ := strconv.Atoi(c.QueryParam("id"))

	ctx := c.Request().Context()

	menu, err := handler.MenuUseCase.GetMenuByIDController(ctx, int(id))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.FromDomain(menu))
}

func (handler MenuController) CreateMenuController(c echo.Context) error {
	menuInsert := request.MenuInsert{}
	c.Bind(&menuInsert)

	ctx := c.Request().Context()

	menu, err := handler.MenuUseCase.CreateMenuController(ctx, menuInsert.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.FromDomain(menu))
}

func (handler MenuController) UpdateMenuController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	menuInsert := request.MenuInsert{}
	c.Bind(&menuInsert)

	ctx := c.Request().Context()
	menu, err := handler.MenuUseCase.UpdateMenuController(ctx, menuInsert.ToDomain(), int(id))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.UpdateSuccessResponse(c, responses.FromDomain(menu))
}

func (handler MenuController) DeleteMenuController(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))

	ctx := c.Request().Context()
	menu, err := handler.MenuUseCase.DeleteMenuController(ctx, int(id))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.DeleteSuccessResponse(c, responses.FromDomain(menu))
}

func (handler MenuController) GetAPI(c echo.Context) error {

	key := c.Param("name")

	ctx := c.Request().Context()
	menu, err := handler.MenuUseCase.GetMenuAPI(ctx, key)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.FromDomain(menu))
}
