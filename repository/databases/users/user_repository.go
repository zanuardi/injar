package users

import (
	"context"
	"injar/usecase/users"

	"gorm.io/gorm"
)

type mysqlUsersRepository struct {
	DB *gorm.DB
}

func NewMySQLUserRepository(db *gorm.DB) users.Repository {
	return &mysqlUsersRepository{
		DB: db,
	}
}

func (repo *mysqlUsersRepository) Fetch(ctx context.Context, page, perpage int) ([]users.Domain, int, error) {
	rec := []Users{}

	offset := (page - 1) * perpage
	err := repo.DB.Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		return []users.Domain{}, 0, err
	}

	var totalData int64
	err = repo.DB.Count(&totalData).Error
	if err != nil {
		return []users.Domain{}, 0, err
	}

	var domainNews []users.Domain
	for _, value := range rec {
		domainNews = append(domainNews, value.toDomain())
	}
	return domainNews, int(totalData), nil
}

func (repo *mysqlUsersRepository) GetByID(ctx context.Context, userId int) (users.Domain, error) {
	rec := Users{}
	err := repo.DB.Where("id = ?", userId).First(&rec).Error
	if err != nil {
		return users.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (repo *mysqlUsersRepository) GetByUsername(ctx context.Context, username string) (users.Domain, error) {
	rec := Users{}
	err := repo.DB.Where("username = ?", username).First(&rec).Error
	if err != nil {
		return users.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (repo *mysqlUsersRepository) Store(ctx context.Context, userDomain *users.Domain) error {
	rec := fromDomain(*userDomain)

	result := repo.DB.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
