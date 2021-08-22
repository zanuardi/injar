package weather

type ResponseWeather []struct {
	Name       string      `json:"name"`
	Base       string      `json:"base"`
	Coord      interface{} `json:"coord"`
	Weather    interface{} `json:"weather"`
	Clouds     interface{} `json:"clouds"`
	Main       interface{} `json:"main"`
	Visibility int         `json:"visibility"`
	Wind       interface{} `json:"wind"`
}
