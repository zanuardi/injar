package routes

import (
	"injar/controllers/users"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware  middleware.JWTConfig
	UserController users.UserController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	auth := e.Group("v1/api/auth")
	auth.POST("/register", cl.UserController.Store)
	auth.POST("/login", cl.UserController.Login)

}
