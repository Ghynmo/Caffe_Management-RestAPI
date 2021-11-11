package transaction_details

import (
	"miniProject/business/transaction_details"
	"miniProject/controllers"
	"miniProject/controllers/transaction_details/request"
	"miniProject/controllers/transaction_details/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TransactionDetailController struct {
	TransactionDetailUseCase transaction_details.UseCase
}

func NewTransactionDetailController(transaction_detailUseCase transaction_details.UseCase) *TransactionDetailController {
	return &TransactionDetailController{
		TransactionDetailUseCase: transaction_detailUseCase,
	}
}

func (handler TransactionDetailController) GetTransactionDetailsController(c echo.Context) error {

	ctx := c.Request().Context()
	transaction_detail, err := handler.TransactionDetailUseCase.GetTransactionDetailsController(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, transaction_detail)
}

func (handler TransactionDetailController) GetTransactionDetailByIDController(c echo.Context) error {

	id, _ := strconv.Atoi(c.QueryParam("id"))

	ctx := c.Request().Context()

	transaction_detail, err := handler.TransactionDetailUseCase.GetTransactionDetailByIDController(ctx, int(id))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.FromDomain(transaction_detail))
}

func (handler TransactionDetailController) CreateTransactionDetailController(c echo.Context) error {
	transaction_detailInsert := request.TransactionDetailCreate{}

	c.Bind(&transaction_detailInsert)

	ctx := c.Request().Context()

	transaction_detail, err := handler.TransactionDetailUseCase.CreateTransactionDetailController(ctx, transaction_detailInsert.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.FromDomain(transaction_detail))
}

func (handler TransactionDetailController) UpdateTransactionDetailController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	transaction_detailInsert := transaction_details.Domain{}
	c.Bind(&transaction_detailInsert)

	ctx := c.Request().Context()
	transaction_detail, err := handler.TransactionDetailUseCase.UpdateTransactionDetailController(ctx, transaction_detailInsert, int(id))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.UpdateSuccessResponse(c, responses.FromDomain(transaction_detail))
}
