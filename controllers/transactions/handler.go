package transactions

import (
	"miniProject/business/transactions"
	"miniProject/controllers"
	"miniProject/controllers/transactions/request"
	"miniProject/controllers/transactions/responses"
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

func (handler TransactionController) GetTransactionController(c echo.Context) error {

	ctx := c.Request().Context()
	transaction, err := handler.TransactionUseCase.GetTransactionController(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.FromListDomain(transaction))
}

func (handler TransactionController) GetTransactionByIDController(c echo.Context) error {

	id, _ := strconv.Atoi(c.QueryParam("id"))

	ctx := c.Request().Context()

	transaction, err := handler.TransactionUseCase.GetTransactionByIDController(ctx, int(id))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.FromDomain(transaction))
}

func (handler TransactionController) CreateTransactionController(c echo.Context) error {
	transactionInsert := request.TransactionBuy{}
	c.Bind(&transactionInsert)

	ctx := c.Request().Context()

	transaction, err := handler.TransactionUseCase.CreateTransactionController(ctx, transactionInsert.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.FromDomain(transaction))
}

func (handler TransactionController) DeleteTransactionController(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))

	ctx := c.Request().Context()
	transaction, err := handler.TransactionUseCase.DeleteTransactionController(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.DeleteSuccessResponse(c, transaction)
}
