package weather

import (
	"context"
	"time"
)

type weatherUsecase struct {
	repo           Repository
	contextTimeout time.Duration
}

func NewWeatherUsecase(timeout time.Duration, repo Repository) Usecase {
	return &weatherUsecase{
		contextTimeout: timeout,
		repo:           repo,
	}
}

func (uc *weatherUsecase) GetAll(ctx context.Context, cityName string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.repo.GetAll(ctx, cityName)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}
