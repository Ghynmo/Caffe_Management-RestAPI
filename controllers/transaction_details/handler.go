package transaction_details

import (
	"miniProject/business/transaction_details"
	"miniProject/controllers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TransactionDetailController struct {
	TransactionDetailUseCase transaction_details.UseCase
}

func NewTransactionDetailController(menuUseCase transaction_details.UseCase) *TransactionDetailController {
	return &TransactionDetailController{
		TransactionDetailUseCase: menuUseCase,
	}
}

func (handler TransactionDetailController) GetTransactionDetailsController(c echo.Context) error {

	ctx := c.Request().Context()
	menu, err := handler.TransactionDetailUseCase.GetTransactionDetailsController(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, menu)
}

func (handler TransactionDetailController) GetTransactionDetailByIDController(c echo.Context) error {

	id, _ := strconv.Atoi(c.QueryParam("id"))

	ctx := c.Request().Context()

	menu, err := handler.TransactionDetailUseCase.GetTransactionDetailByIDController(ctx, uint(id))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, menu)
}

func (handler TransactionDetailController) CreateTransactionDetailController(c echo.Context) error {
	menuInsert := transaction_details.Domain{}
	c.Bind(&menuInsert)

	ctx := c.Request().Context()

	menu, err := handler.TransactionDetailUseCase.CreateTransactionDetailController(ctx, menuInsert)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, menu)
}

func (handler TransactionDetailController) UpdateTransactionDetailController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	menuInsert := transaction_details.Domain{}
	c.Bind(&menuInsert)

	ctx := c.Request().Context()
	menu, err := handler.TransactionDetailUseCase.UpdateTransactionDetailController(ctx, menuInsert, uint(id))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, menu)
}
