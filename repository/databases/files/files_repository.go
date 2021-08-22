package files

import (
	"injar/usecase/files"

	"gorm.io/gorm"
)

type FileRepository struct {
	DB *gorm.DB
}

func NewFileRepository(db *gorm.DB) files.Repository {
	return &FileRepository{
		DB: db,
	}
}

func (repo *FileRepository) FindByID(id int) (files.Domain, error) {
	rec := File{}

	err := repo.DB.Find(&rec, id).Error
	if err != nil {
		return files.Domain{}, err
	}

	return rec.toDomain(), nil
}

func (repo *FileRepository) Store(file *files.Domain) (files.Domain, error) {
	rec := fromDomain(*file)

	err := repo.DB.Create(&rec).Error
	if err != nil {
		return files.Domain{}, err
	}

	return rec.toDomain(), nil
}
