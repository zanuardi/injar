package webinars

import (
	"injar/repository/databases/categories"
	"injar/repository/databases/users"
	"injar/usecase/webinars"

	"time"

	"gorm.io/gorm"
)

type Webinars struct {
	ID          int
	UserID      int
	User        users.Users `gorm:"foreignKey:UserID;references:ID"`
	CategoryID  int
	Categories  categories.Categories `gorm:"foreignKey:CategoryID;references:ID"`
	ImageUrl    string
	Name        string
	Description string
	Price       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

func (rec *Webinars) toDomain() webinars.Domain {
	return webinars.Domain{
		ID:           rec.ID,
		UserID:       rec.UserID,
		CategoryID:   rec.CategoryID,
		CategoryName: rec.Categories.Name,
		ImageUrl:     rec.ImageUrl,
		Name:         rec.Name,
		Description:  rec.Description,
		Price:        rec.Price,
		CreatedAt:    rec.CreatedAt,
		UpdatedAt:    rec.UpdatedAt,
		DeletedAt:    rec.DeletedAt,
	}
}

func fromDomain(webinarsDomain webinars.Domain) *Webinars {
	return &Webinars{
		ID:          webinarsDomain.ID,
		UserID:      webinarsDomain.UserID,
		CategoryID:  webinarsDomain.CategoryID,
		ImageUrl:    webinarsDomain.ImageUrl,
		Name:        webinarsDomain.Name,
		Description: webinarsDomain.Description,
		Price:       webinarsDomain.Price,
		CreatedAt:   webinarsDomain.CreatedAt,
		UpdatedAt:   webinarsDomain.UpdatedAt,
		DeletedAt:   webinarsDomain.DeletedAt,
	}
}
