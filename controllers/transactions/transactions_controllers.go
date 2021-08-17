package transactions

import (
	"errors"
	"injar/app/middleware"
	"injar/controllers/transactions/request"
	"injar/controllers/transactions/response"
	"injar/usecase/transactions"
	"net/http"
	"strconv"
	"strings"

	controller "injar/controllers"

	echo "github.com/labstack/echo/v4"
)

type TransactionsController struct {
	TransactionsUC transactions.Usecase
	jwtAuth        *middleware.ConfigJWT
}

func NewTransactionsController(uc transactions.Usecase) *TransactionsController {
	return &TransactionsController{
		TransactionsUC: uc,
	}
}

func (ctrl *TransactionsController) GetByUserID(c echo.Context) error {
	ctx := c.Request().Context()
	user, err := ctrl.jwtAuth.GetUser(c)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > 50 {
		limit = 10
	}

	resp, total, err := ctrl.TransactionsUC.GetByUserID(ctx, page, limit, user.ID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.Transaction{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	pagination := controller.PaginationRes(page, total, limit)
	return controller.NewSuccessResponsePagination(c, responseController, pagination)
}

func (ctrl *TransactionsController) GetById(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	favourite, err := ctrl.TransactionsUC.GetByID(ctx, id)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(favourite))

}

func (ctrl *TransactionsController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	user, err := ctrl.jwtAuth.GetUser(c)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	req := request.Transactions{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	domainReq.UserID = user.ID
	resp, err := ctrl.TransactionsUC.Store(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}

func (ctrl *TransactionsController) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	if strings.TrimSpace(id) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	req := request.Transactions{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	idInt, _ := strconv.Atoi(id)
	domainReq.ID = idInt
	resp, err := ctrl.TransactionsUC.Delete(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(*resp))
}
