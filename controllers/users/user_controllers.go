package users

import (
	"injar/app/middleware"
	controller "injar/controllers"
	"injar/controllers/users/request"
	"injar/usecase/users"
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

type UserController struct {
	userUseCase users.Usecase
	jwtAuth     *middleware.ConfigJWT
}

func NewUserController(uc users.Usecase) *UserController {
	return &UserController{
		userUseCase: uc,
	}
}

func (ctrl *UserController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Users{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.userUseCase.Store(ctx, req.ToDomain())
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

	token, err := ctrl.userUseCase.CreateToken(ctx, req.Username, req.Password)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	response := struct {
		Token string `json:"token"`
	}{Token: token}

	return controller.NewSuccessResponse(c, response)
}

func (ctrl *UserController) FindById(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	user, err := ctrl.userUseCase.GetByID(ctx, id)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controller.NewSuccessResponse(c, user)
}

func (ctrl *UserController) FindByToken(c echo.Context) error {
	ctx := c.Request().Context()

	user, err := ctrl.jwtAuth.GetUser(c)

	id, err := ctrl.userUseCase.GetByID(ctx, user.ID)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controller.NewSuccessResponse(c, id)
}
