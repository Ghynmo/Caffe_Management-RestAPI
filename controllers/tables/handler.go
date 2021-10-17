package tables

import (
	"miniProject/business/tables"
	"miniProject/controllers"
	"miniProject/controllers/tables/request"
	"miniProject/controllers/tables/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TableController struct {
	TableUseCase tables.UseCase
}

func NewTableController(tableUseCase tables.UseCase) *TableController {
	return &TableController{
		TableUseCase: tableUseCase,
	}
}

func (handler TableController) GetTablesController(c echo.Context) error {

	ctx := c.Request().Context()
	table, err := handler.TableUseCase.GetTablesController(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.FromListDomain(table))
}

func (handler TableController) GetTableByIDController(c echo.Context) error {

	id, _ := strconv.Atoi(c.QueryParam("id"))

	ctx := c.Request().Context()

	table, err := handler.TableUseCase.GetTableByIDController(ctx, int(id))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.FromDomain(table))
}

func (handler TableController) CreateTableController(c echo.Context) error {
	tableInsert := request.TableInsert{}
	c.Bind(&tableInsert)

	ctx := c.Request().Context()

	table, err := handler.TableUseCase.CreateTableController(ctx, tableInsert.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.FromDomain(table))
}

func (handler TableController) UpdateTableController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	tableInsert := request.TableInsert{}
	c.Bind(&tableInsert)

	ctx := c.Request().Context()
	table, err := handler.TableUseCase.UpdateTableController(ctx, tableInsert.ToDomain(), int(id))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.UpdateSuccessResponse(c, responses.FromDomain(table))
}

func (handler TableController) DeleteTableController(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))

	ctx := c.Request().Context()
	table, err := handler.TableUseCase.DeleteTableController(ctx, int(id))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.DeleteSuccessResponse(c, responses.FromDomain(table))
}
