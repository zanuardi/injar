package categories

import (
	"errors"
	"injar/businesses/categories"
	"injar/controllers/categories/request"
	"injar/controllers/categories/response"
	"net/http"
	"strconv"
	"strings"

	controller "injar/controllers"

	echo "github.com/labstack/echo/v4"
)

type CategoryController struct {
	categoryUsecase categories.Usecase
}

func NewCategoryController(cu categories.Usecase) *CategoryController {
	return &CategoryController{
		categoryUsecase: cu,
	}
}

func (ctrl *CategoryController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := ctrl.categoryUsecase.GetAll(ctx)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.Categories{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController)
}

func (ctrl *CategoryController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Categories{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	resp, err := ctrl.categoryUsecase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}

func (ctrl *CategoryController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.QueryParam("id")
	if strings.TrimSpace(id) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	req := request.Categories{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	idInt, _ := strconv.Atoi(id)
	domainReq.ID = idInt
	resp, err := ctrl.categoryUsecase.Update(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(*resp))
}
