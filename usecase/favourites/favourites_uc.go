package favourites

import (
	"context"
	usecase "injar/usecase"
	"time"
)

type favouritesUsecase struct {
	favouritesRepository Repository
	contextTimeout       time.Duration
}

func NewFavouritesUsecase(timeout time.Duration, cr Repository) Usecase {
	return &favouritesUsecase{
		contextTimeout:       timeout,
		favouritesRepository: cr,
	}
}

func (uc *favouritesUsecase) GetByUserID(ctx context.Context, UserID int) ([]Domain, error) {
	resp, err := uc.favouritesRepository.GetByUserID(ctx, UserID)
	if err != nil {
		return []Domain{}, err
	}
	return resp, nil
}

func (uc *favouritesUsecase) GetByID(ctx context.Context, ID int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if ID <= 0 {
		return Domain{}, usecase.ErrNotFound
	}
	res, err := uc.favouritesRepository.GetByID(ctx, ID)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (uc *favouritesUsecase) Store(ctx context.Context, favouriteDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedFavourites, _ := uc.favouritesRepository.GetByID(ctx, favouriteDomain.ID)

	if existedFavourites != (Domain{}) {
		return Domain{}, usecase.ErrDuplicateData
	}

	result, err := uc.favouritesRepository.Store(ctx, favouriteDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (uc *favouritesUsecase) Delete(ctx context.Context, favouritesDomain *Domain) (*Domain, error) {
	existedFavourites, err := uc.favouritesRepository.GetByID(ctx, favouritesDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	favouritesDomain.ID = existedFavourites.ID

	result, err := uc.favouritesRepository.Delete(ctx, favouritesDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}
