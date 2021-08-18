package routes

import (
	"injar/controllers/categories"
	"injar/controllers/favourites"
	"injar/controllers/files"
	"injar/controllers/transactions"
	"injar/controllers/users"
	"injar/controllers/webinars"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware          middleware.JWTConfig
	UserController         users.UserController
	CategoriesController   categories.CategoryController
	WebinarController      webinars.WebinarController
	FavouritesController   favourites.FavouritesController
	TransactionsController transactions.TransactionsController
	FileController         files.FileController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	r := e.Group("/v1/api")

	//Auth ...
	auth := r.Group("/auth")

	auth.POST("/register", cl.UserController.Store)
	auth.POST("/login", cl.UserController.Login)

	//Using Bearer Token ...
	r.Use(middleware.JWTWithConfig(cl.JWTMiddleware))

	//Users ...
	user := r.Group("/profile")

	user.GET("", cl.UserController.FindByToken)
	user.PUT("", cl.UserController.Update)

	//Categories ...
	category := r.Group("/categories")

	category.GET("/select", cl.CategoriesController.SelectAll)
	category.GET("", cl.CategoriesController.FindAll)
	category.GET("/id/:id", cl.CategoriesController.FindById)
	category.POST("", cl.CategoriesController.Store)
	category.PUT("/id/:id", cl.CategoriesController.Update)
	category.DELETE("/id/:id", cl.CategoriesController.Delete)

	//Webinars ...
	webinar := r.Group("/webinars")

	webinar.GET("/select", cl.WebinarController.SelectAll)
	webinar.GET("", cl.WebinarController.FindAll)
	webinar.GET("/id/:id", cl.WebinarController.FindById)
	webinar.POST("", cl.WebinarController.Store)
	webinar.PUT("/id/:id", cl.WebinarController.Update)
	webinar.DELETE("/id/:id", cl.WebinarController.Delete)

	//Favourites ...
	favourites := r.Group("/favourites")

	favourites.GET("", cl.FavouritesController.GetByUserID)
	favourites.GET("/id/:id", cl.FavouritesController.GetById)
	favourites.POST("", cl.FavouritesController.Store)
	favourites.DELETE("/id/:id", cl.FavouritesController.Delete)

	//transactions ...
	transactions := r.Group("/transactions")

	transactions.GET("", cl.TransactionsController.GetByUserID)
	transactions.GET("/id/:id", cl.TransactionsController.GetById)
	transactions.POST("", cl.TransactionsController.Store)
	transactions.DELETE("/id/:id", cl.TransactionsController.Delete)

	fileRoute := r.Group("/files")
	fileRoute.POST("", cl.FileController.Store)
}
