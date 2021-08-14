package main

import (
	_middleware "injar/app/middleware"
	_routes "injar/app/routes"
	_userUsecase "injar/businesses/users"
	_userController "injar/controllers/users"
	_userRepo "injar/drivers/databases/users"

	_categoriesUsecase "injar/businesses/categories"
	_categoriesController "injar/controllers/categories"
	_categoriesRepo "injar/drivers/databases/categories"

	_dbDriver "injar/drivers/mysql"

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

	routesInit := _routes.ControllerList{
		JWTMiddleware:        configJWT.Init(),
		UserController:       *userCtrl,
		CategoriesController: *categoriesCtrl,
	}
	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
