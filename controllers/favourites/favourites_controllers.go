package favourites

import (
	"errors"
	"injar/controllers/favourites/request"
	"injar/controllers/favourites/response"
	"injar/usecase/favourites"
	"net/http"
	"strconv"
	"strings"

	controller "injar/controllers"

	echo "github.com/labstack/echo/v4"
)

type FavouritesController struct {
	FavouritesUC favourites.Usecase
}

func NewFavouritesController(uc favourites.Usecase) *FavouritesController {
	return &FavouritesController{
		FavouritesUC: uc,
	}
}

func (ctrl *FavouritesController) GetByUserID(c echo.Context) error {
	ctx := c.Request().Context()
	userID, err := strconv.Atoi(c.Param("user_id"))

	resp, err := ctrl.FavouritesUC.GetByUserID(ctx, userID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.Favourite{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController)
}

func (ctrl *FavouritesController) GetById(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	favourite, err := ctrl.FavouritesUC.GetByID(ctx, id)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controller.NewSuccessResponse(c, favourite)

}

func (ctrl *FavouritesController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Favourites{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	resp, err := ctrl.FavouritesUC.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}

func (ctrl *FavouritesController) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	if strings.TrimSpace(id) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	req := request.Favourites{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	idInt, _ := strconv.Atoi(id)
	domainReq.ID = idInt
	resp, err := ctrl.FavouritesUC.Delete(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(*resp))
}
