package inventories

import (
	"miniProject/business/inventories"
	"miniProject/controllers"
	"miniProject/controllers/inventories/request"
	"miniProject/controllers/inventories/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type InventoryController struct {
	InventoryUseCase inventories.UseCase
}

func NewInventoryController(inventoryUseCase inventories.UseCase) *InventoryController {
	return &InventoryController{
		InventoryUseCase: inventoryUseCase,
	}
}

func (handler InventoryController) GetInventoriesController(c echo.Context) error {

	ctx := c.Request().Context()
	inventory, err := handler.InventoryUseCase.GetInventoriesController(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.FromListDomain(inventory))
}

func (handler InventoryController) GetInventoryByIDController(c echo.Context) error {

	id, _ := strconv.Atoi(c.QueryParam("id"))

	ctx := c.Request().Context()

	inventory, err := handler.InventoryUseCase.GetInventoryByIDController(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.FromDomain(inventory))
}

func (handler InventoryController) CreateInventoryController(c echo.Context) error {
	inventoryInsert := request.InventoryInsert{}
	c.Bind(&inventoryInsert)

	ctx := c.Request().Context()

	inventory, err := handler.InventoryUseCase.CreateInventoryController(ctx, inventoryInsert.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.FromDomain(inventory))
}

func (handler InventoryController) UpdateInventoryController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	inventoryInsert := request.InventoryInsert{}
	c.Bind(&inventoryInsert)

	ctx := c.Request().Context()
	inventory, err := handler.InventoryUseCase.UpdateInventoryController(ctx, inventoryInsert.ToDomain(), id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.UpdateSuccessResponse(c, responses.FromDomain(inventory))
}

func (handler InventoryController) DeleteInventoryController(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))

	ctx := c.Request().Context()
	inventory, err := handler.InventoryUseCase.DeleteInventoryController(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.DeleteSuccessResponse(c, responses.FromDomain(inventory))
}
