package transactions

import (
	"context"
	"injar/usecase/transactions"

	"gorm.io/gorm"
)

type mysqlTransactionsRepository struct {
	DB *gorm.DB
}

func NewMySQLTransactionsRepository(conn *gorm.DB) transactions.Repository {
	return &mysqlTransactionsRepository{
		DB: conn,
	}
}

func (repo *mysqlTransactionsRepository) GetByUserID(ctx context.Context, userID int) ([]transactions.Domain, error) {
	rec := []Transactions{}

	err := repo.DB.Joins("Users").Joins("Webinars").Where("transactions.user_id = ?", userID).Find(&rec).Error
	if err != nil {
		return []transactions.Domain{}, err
	}

	tranasctionDomain := []transactions.Domain{}
	for _, value := range rec {
		tranasctionDomain = append(tranasctionDomain, value.toDomain())
	}

	return tranasctionDomain, nil
}

func (repo *mysqlTransactionsRepository) GetByID(ctx context.Context, ID int) (transactions.Domain, error) {
	rec := Transactions{}
	err := repo.DB.Joins("Users").Joins("Webinars").Where("transactions.id = ?", ID).First(&rec).Error
	if err != nil {
		return transactions.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (repo *mysqlTransactionsRepository) Store(ctx context.Context, transactionsDomain *transactions.Domain) (transactions.Domain, error) {
	rec := fromDomain(*transactionsDomain)

	result := repo.DB.Create(&rec)
	if result.Error != nil {
		return transactions.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (repo *mysqlTransactionsRepository) Delete(ctx context.Context, transactionsDomain *transactions.Domain) (transactions.Domain, error) {
	rec := fromDomain(*transactionsDomain)

	result := repo.DB.Delete(rec)

	if result.Error != nil {
		return transactions.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}
