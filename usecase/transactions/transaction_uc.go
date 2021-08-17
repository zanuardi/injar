package transactions

import (
	"context"
	usecase "injar/usecase"
	"time"
)

type transactionsUsecase struct {
	transactionsRepository Repository
	contextTimeout         time.Duration
}

func NewTransactionsUsecase(timeout time.Duration, cr Repository) Usecase {
	return &transactionsUsecase{
		contextTimeout:         timeout,
		transactionsRepository: cr,
	}
}

func (uc *transactionsUsecase) GetByUserID(ctx context.Context, UserID int) ([]Domain, error) {
	resp, err := uc.transactionsRepository.GetByUserID(ctx, UserID)
	if err != nil {
		return []Domain{}, err
	}
	return resp, nil
}

func (uc *transactionsUsecase) GetByID(ctx context.Context, ID int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if ID <= 0 {
		return Domain{}, usecase.ErrNotFound
	}
	res, err := uc.transactionsRepository.GetByID(ctx, ID)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (uc *transactionsUsecase) Store(ctx context.Context, favouriteDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedTransaction, _ := uc.transactionsRepository.GetByID(ctx, favouriteDomain.ID)

	if existedTransaction != (Domain{}) {
		return Domain{}, usecase.ErrDuplicateData
	}

	result, err := uc.transactionsRepository.Store(ctx, favouriteDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (uc *transactionsUsecase) Delete(ctx context.Context, transactionsDomain *Domain) (*Domain, error) {
	existedTransaction, err := uc.transactionsRepository.GetByID(ctx, transactionsDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	transactionsDomain.ID = existedTransaction.ID

	result, err := uc.transactionsRepository.Delete(ctx, transactionsDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}
