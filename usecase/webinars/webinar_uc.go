package webinars

import (
	"context"
	usecase "injar/usecase"
	"strings"
	"time"
)

type webinarUsecase struct {
	webinarRepository Repository
	contextTimeout    time.Duration
}

func NewWebinarUsecase(timeout time.Duration, wr Repository) Usecase {
	return &webinarUsecase{
		contextTimeout:    timeout,
		webinarRepository: wr,
	}
}

func (uc *webinarUsecase) GetAll(ctx context.Context, name string) ([]Domain, error) {
	resp, err := uc.webinarRepository.GetAll(ctx, name)
	if err != nil {
		return []Domain{}, err
	}
	return resp, nil
}

func (uc *webinarUsecase) FindAll(ctx context.Context, name, category, schedule, priceFrom, priceTo string, page, perpage int) ([]Domain, int, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 25
	}

	res, total, err := uc.webinarRepository.FindAll(ctx, name, category, schedule, priceFrom, priceTo, page, perpage)
	if err != nil {
		return []Domain{}, 0, err
	}

	return res, total, nil
}

func (uc *webinarUsecase) GetByID(ctx context.Context, webinarID int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if webinarID <= 0 {
		return Domain{}, usecase.ErrNotFound
	}
	res, err := uc.webinarRepository.GetByID(ctx, webinarID)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (uc *webinarUsecase) GetByName(ctx context.Context, webinarName string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if strings.TrimSpace(webinarName) == "" {
		return Domain{}, usecase.ErrNotFound
	}
	res, err := uc.webinarRepository.GetByName(ctx, webinarName)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (uc *webinarUsecase) Store(ctx context.Context, webinarDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedWebinar, _ := uc.webinarRepository.GetByName(ctx, webinarDomain.Name)

	if existedWebinar != (Domain{}) {
		return Domain{}, usecase.ErrDuplicateData
	}

	result, err := uc.webinarRepository.Store(ctx, webinarDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (uc *webinarUsecase) Update(ctx context.Context, webinarDomain *Domain) (*Domain, error) {
	existedWebinar, err := uc.webinarRepository.GetByID(ctx, webinarDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	webinarDomain.ID = existedWebinar.ID

	result, err := uc.webinarRepository.Update(ctx, webinarDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}

func (uc *webinarUsecase) Delete(ctx context.Context, webinarDomain *Domain) (*Domain, error) {
	existedWebinar, err := uc.webinarRepository.GetByID(ctx, webinarDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	webinarDomain.ID = existedWebinar.ID

	result, err := uc.webinarRepository.Delete(ctx, webinarDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}
