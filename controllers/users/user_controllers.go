package users

import (
	"injar/app/middleware"
	controller "injar/controllers"
	"injar/controllers/users/request"
	"injar/controllers/users/response"
	"injar/usecase/users"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type UserController struct {
	userUsecase users.Usecase
	jwtAuth     *middleware.ConfigJWT
}

func NewUserController(uc users.Usecase) *UserController {
	return &UserController{
		userUsecase: uc,
	}
}

func (ctrl *UserController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Users{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.userUsecase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, "Successfully inserted")
}

func (ctrl *UserController) Login(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Login{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	token, err := ctrl.userUsecase.CreateToken(ctx, req.Username, req.Password)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	response := struct {
		Token string `json:"token"`
	}{Token: token}

	return controller.NewSuccessResponse(c, response)
}

func (ctrl *UserController) FindByToken(c echo.Context) error {
	ctx := c.Request().Context()

	user, err := ctrl.jwtAuth.GetUser(c)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	id, err := ctrl.userUsecase.GetByID(ctx, user.ID)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controller.NewSuccessResponse(c, id)
}

func (ctrl *UserController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	user, err := ctrl.jwtAuth.GetUser(c)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	req := request.Users{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	domainReq.ID = user.ID
	resp, err := ctrl.userUsecase.Update(ctx, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(*resp))
}
