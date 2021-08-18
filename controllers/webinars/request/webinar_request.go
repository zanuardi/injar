package request

import (
	"injar/usecase/webinars"
	"time"
)

type Webinars struct {
	UserID      int       `json:"user_id"`
	CategoryID  int       `json:"category_id"`
	ImageUrl    string    `json:"image_url"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Schedule    time.Time `json:"schedule"`
}

func (req *Webinars) ToDomain() *webinars.Domain {
	return &webinars.Domain{
		UserID:      req.UserID,
		CategoryID:  req.CategoryID,
		ImageUrl:    req.ImageUrl,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Schedule:    req.Schedule,
	}
}
