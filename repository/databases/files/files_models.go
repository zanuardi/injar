package files

import (
	"injar/usecase/files"
	"time"

	"gorm.io/gorm"
)

type File struct {
	ID        int `gorm:"primaryKey"`
	Type      string
	Path      string
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (rec *File) toDomain() files.Domain {
	return files.Domain{
		ID:        rec.ID,
		Type:      rec.Type,
		Path:      rec.Path,
		CreatedAt: rec.CreatedAt,
		DeletedAt: rec.DeletedAt,
	}
}

func fromDomain(fileDomain files.Domain) *File {
	return &File{
		ID:        fileDomain.ID,
		Type:      fileDomain.Type,
		Path:      fileDomain.Path,
		CreatedAt: fileDomain.CreatedAt,
		DeletedAt: fileDomain.DeletedAt,
	}
}
