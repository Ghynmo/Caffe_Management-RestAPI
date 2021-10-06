package table_details

import (
	"miniProject/business/table_details"
	"miniProject/controllers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TableDetailController struct {
	TableDetailUseCase table_details.UseCase
}

func NewTableDetailController(table_detailUseCase table_details.UseCase) *TableDetailController {
	return &TableDetailController{
		TableDetailUseCase: table_detailUseCase,
	}
}

func (handler TableDetailController) GetTableDetailsController(c echo.Context) error {

	ctx := c.Request().Context()
	table_detail, err := handler.TableDetailUseCase.GetTableDetailsController(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, table_detail)
}

func (handler TableDetailController) GetTableDetailByIDController(c echo.Context) error {

	id, _ := strconv.Atoi(c.QueryParam("id"))

	ctx := c.Request().Context()

	table_detail, err := handler.TableDetailUseCase.GetTableDetailByIDController(ctx, uint(id))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, table_detail)
}

func (handler TableDetailController) CreateTableDetailController(c echo.Context) error {
	table_detailInsert := table_details.Domain{}
	c.Bind(&table_detailInsert)

	ctx := c.Request().Context()

	table_detail, err := handler.TableDetailUseCase.CreateTableDetailController(ctx, table_detailInsert)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, table_detail)
}

func (handler TableDetailController) UpdateTableDetailController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	table_detailInsert := table_details.Domain{}
	c.Bind(&table_detailInsert)

	ctx := c.Request().Context()
	table_detail, err := handler.TableDetailUseCase.UpdateTableDetailController(ctx, table_detailInsert, uint(id))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, table_detail)
}
