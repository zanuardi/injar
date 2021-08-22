package weather

import "context"

type Domain struct {
	Name       string      `json:"name"`
	Base       string      `json:"base"`
	Coord      interface{} `json:"coord"`
	Weather    interface{} `json:"weather"`
	Clouds     interface{} `json:"clouds"`
	Main       interface{} `json:"main"`
	Visibility int         `json:"visibility"`
	Wind       interface{} `json:"wind"`
}

type Usecase interface {
	GetAll(ctx context.Context, cityName string) (Domain, error)
}

type Repository interface {
	GetAll(ctx context.Context, cityName string) (Domain, error)
}
