package controllers

import (
	"injar/configs"
	"injar/models/base"
	"injar/models/users"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterUserController(c echo.Context) error {
	var userRegister users.UserRegister
	c.Bind(&userRegister)

	var userDB users.User
	userDB.Name = userRegister.Name
	userDB.Username = userRegister.Username
	userDB.Email = userRegister.Email
	userDB.Password = userRegister.Password

	result := configs.DB.Create(&userDB)
	if result.Error != nil {
		var response = base.BaseResponse{
			Code:    http.StatusInternalServerError,
			Status:  false,
			Message: "Failed register User",
			Data:    nil,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	var response = base.BaseResponse{
		Code:    http.StatusOK,
		Status:  true,
		Message: "Success register user",
		Data:    userDB,
	}
	return c.JSON(http.StatusOK, response)
}

func LoginUserController(c echo.Context) error {
	var userLogin users.UserLogin
	c.Bind(&userLogin)

	var response = base.BaseResponse{
		Code:    http.StatusOK,
		Status:  true,
		Message: "Success",
		Data:    userLogin,
	}
	return c.JSON(http.StatusOK, response)
}

// func DetailUserController(c echo.Context) error {
// 	userId, _ := strconv.Atoi(c.Param("userId"))

// 	var user = User{userId, "Alterra", "alta@gmail.com", "Malang"}

// 	var response = BaseResponse{
// 		Code:    http.StatusOK,
// 		Status:  true,
// 		Message: "Success",
// 		Data:    user,
// 	}
// 	return c.JSON(http.StatusOK, response)
// }

func GetUserController(c echo.Context) error {
	var users []users.User

	result := configs.DB.Find(&users)
	if result.Error != nil {
		var response = base.BaseResponse{
			Code:    http.StatusInternalServerError,
			Status:  false,
			Message: "Failed register User",
			Data:    nil,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	var response = base.BaseResponse{
		Code:    http.StatusOK,
		Status:  true,
		Message: "Success get Data",
		Data:    users,
	}
	return c.JSON(http.StatusOK, response)
}
