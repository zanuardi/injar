package files

import (
	"mime/multipart"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        int `gorm:"primaryKey"`
	Type      string
	Path      string
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type Usecase interface {
	// FindByID(id int) (res entity.File, err error)
	Store(fileType string, file *multipart.FileHeader) (Domain, error)
}

type Repository interface {
	FindByID(id int) (Domain, error)
	Store(file *Domain) (Domain, error)
}
