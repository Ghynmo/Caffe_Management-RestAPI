package transactions

import (
	"fmt"
	"miniProject/business/transactions"
	"miniProject/controllers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TransactionController struct {
	TransactionUseCase transactions.UseCase
}

func NewTransactionController(transactionUseCase transactions.UseCase) *TransactionController {
	return &TransactionController{
		TransactionUseCase: transactionUseCase,
	}
}

func (handler TransactionController) BuyTransactionController(c echo.Context) error {
	DataRequest := transactions.Domain{}
	c.Bind(&DataRequest)
	fmt.Println(DataRequest)
	ctx := c.Request().Context()
	transaction, err := handler.TransactionUseCase.BuyTransactionController(ctx, DataRequest)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, transaction)
}

func (handler TransactionController) GetTransactionByIDController(c echo.Context) error {

	id, _ := strconv.Atoi(c.QueryParam("id"))

	ctx := c.Request().Context()

	transaction, err := handler.TransactionUseCase.GetTransactionByIDController(ctx, uint(id))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, transaction)
}

func (handler TransactionController) CreateTransactionController(c echo.Context) error {
	transactionInsert := transactions.Domain{}
	c.Bind(&transactionInsert)

	ctx := c.Request().Context()

	transaction, err := handler.TransactionUseCase.CreateTransactionController(ctx, transactionInsert)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, transaction)
}

func (handler TransactionController) UpdateTransactionController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	transactionInsert := transactions.Domain{}
	c.Bind(&transactionInsert)

	ctx := c.Request().Context()
	transaction, err := handler.TransactionUseCase.UpdateTransactionController(ctx, transactionInsert, uint(id))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, transaction)
}
