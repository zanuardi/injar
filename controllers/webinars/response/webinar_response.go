package response

import (
	"injar/repository/databases/categories"
	"injar/usecase/webinars"
	"time"

	"gorm.io/gorm"
)

type Webinar struct {
	ID           int                   `json:"id"`
	UserID       int                   `json:"user_id"`
	CategoryID   int                   `json:"category_id"`
	Categories   categories.Categories `json:"categories"`
	CategoryName string                `json:"category_name"`
	ImageUrl     string                `json:"image_url"`
	Name         string                `json:"name"`
	Description  string                `json:"description"`
	Price        string                `json:"price"`
	CreatedAt    time.Time             `json:"created_at"`
	UpdatedAt    time.Time             `json:"updated_at"`
	DeletedAt    gorm.DeletedAt        `json:"deleted_at"`
}

func FromDomain(domain webinars.Domain) Webinar {
	return Webinar{
		ID:         domain.ID,
		UserID:     domain.UserID,
		CategoryID: domain.CategoryID,
		// Categories:  domain.Categories,
		ImageUrl:    domain.ImageUrl,
		Name:        domain.Name,
		Description: domain.Description,
		Price:       domain.Price,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
		DeletedAt:   domain.DeletedAt,
	}
}
