package weather

import (
	controller "injar/controllers"
	"injar/usecase/weather"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type WeatherController struct {
	weatherUC weather.Usecase
}

func NewWeatherController(e *echo.Echo, uc weather.Usecase) *WeatherController {
	return &WeatherController{
		weatherUC: uc,
	}
}

func (ctrl *WeatherController) GetByCity(c echo.Context) error {
	ctx := c.Request().Context()
	cityName := c.Param("city_name")
	resp, err := ctrl.weatherUC.GetAll(ctx, cityName)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, resp)
}
