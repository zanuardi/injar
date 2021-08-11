package routes

import (
	"injar/controllers/categories"
	"injar/controllers/users"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware        middleware.JWTConfig
	UserController       users.UserController
	CategoriesController categories.CategoryController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	auth := e.Group("v1/api/auth")
	auth.POST("/register", cl.UserController.Store)
	auth.POST("/login", cl.UserController.Login)

	category := e.Group("v1/api/categories")
	category.GET("", cl.CategoriesController.GetAll)
	category.POST("", cl.CategoriesController.Store)
	category.POST("/:id", cl.CategoriesController.Update)

}
