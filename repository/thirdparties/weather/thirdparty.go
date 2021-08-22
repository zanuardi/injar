package weather

import (
	"context"
	"encoding/json"
	"injar/usecase/weather"
	"io/ioutil"
	"net/http"
)

type Weather struct {
	httpClient http.Client
}

func NewWeatherRepository() weather.Repository {
	return &Weather{
		httpClient: http.Client{},
	}
}

func (repo *Weather) GetAll(ctx context.Context, cityName string) (res weather.Domain, err error) {
	req, _ := http.NewRequest("GET", "https://api.openweathermap.org/data/2.5/weather?q="+cityName+"&appid=f331625da3d58a3a89378b293c7ad913", nil)
	resp, err := repo.httpClient.Do(req)
	if err != nil {
		return res, err
	}

	defer resp.Body.Close()

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal([]byte(body), &res)
	return res, err

}
