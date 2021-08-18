package main

import (
	_middleware "injar/app/middleware"
	_routes "injar/app/routes"
	_userController "injar/controllers/users"
	_userRepo "injar/repository/databases/users"
	_userUsecase "injar/usecase/users"

	_categoriesController "injar/controllers/categories"
	_categoriesRepo "injar/repository/databases/categories"
	_categoriesUsecase "injar/usecase/categories"

	_webinarsController "injar/controllers/webinars"
	_webinarsRepo "injar/repository/databases/webinars"
	_webinarsUsecase "injar/usecase/webinars"

	_favouritesController "injar/controllers/favourites"
	_favouritesRepo "injar/repository/databases/favourites"
	_favouritesUsecase "injar/usecase/favourites"

	_transactionsController "injar/controllers/transactions"
	_transactionsRepo "injar/repository/databases/transactions"
	_transactionsUsecase "injar/usecase/transactions"

	_filesController "injar/controllers/files"
	_filesRepo "injar/repository/databases/files"
	_filesUsecase "injar/usecase/files"

	_dbDriver "injar/repository/mysql"

	"log"
	"time"

	echo "github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}
	db := configDB.InitDB()

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	e := echo.New()

	// Users ...
	userRepo := _userRepo.NewMySQLUserRepository(db)
	userUsecase := _userUsecase.NewUserUsecase(userRepo, &configJWT, timeoutContext)
	userCtrl := _userController.NewUserController(userUsecase)

	// Categories ...
	categoriesRepo := _categoriesRepo.NewMySQLCategoryRepository(db)
	categoriesUsecase := _categoriesUsecase.NewCategoryUsecase(timeoutContext, categoriesRepo)
	categoriesCtrl := _categoriesController.NewCategoryController(categoriesUsecase)

	// Webinars ...
	webinarsRepo := _webinarsRepo.NewMySQLWebinarRepository(db)
	webinarsUsecase := _webinarsUsecase.NewWebinarUsecase(timeoutContext, webinarsRepo)
	webinarsCtrl := _webinarsController.NewWebinarController(webinarsUsecase)

	// favourites ...
	favouritesRepo := _favouritesRepo.NewMySQLFavouritesRepository(db)
	favouritesUsecase := _favouritesUsecase.NewFavouritesUsecase(timeoutContext, favouritesRepo)
	favouritesCtrl := _favouritesController.NewFavouritesController(favouritesUsecase)

	// transactions ...
	transactionsRepo := _transactionsRepo.NewMySQLTransactionsRepository(db)
	transactionsUsecase := _transactionsUsecase.NewTransactionsUsecase(timeoutContext, transactionsRepo)
	transactionsCtrl := _transactionsController.NewTransactionsController(transactionsUsecase)

	// files ...
	filesRepo := _filesRepo.NewFileRepository(db)
	filesUsecase := _filesUsecase.NewFileUC(timeoutContext, filesRepo)
	filesCtrl := _filesController.NewFileController(filesUsecase, db)

	routesInit := _routes.ControllerList{
		JWTMiddleware:          configJWT.Init(),
		UserController:         *userCtrl,
		CategoriesController:   *categoriesCtrl,
		WebinarController:      *webinarsCtrl,
		FavouritesController:   *favouritesCtrl,
		TransactionsController: *transactionsCtrl,
		FileController:         *filesCtrl,
	}
	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
