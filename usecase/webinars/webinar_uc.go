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

func (webinarUC *webinarUsecase) GetAll(ctx context.Context, name string) ([]Domain, error) {
	resp, err := webinarUC.webinarRepository.GetAll(ctx, name)
	if err != nil {
		return []Domain{}, err
	}
	return resp, nil
}

func (webinarUC *webinarUsecase) GetByID(ctx context.Context, webinarID int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, webinarUC.contextTimeout)
	defer cancel()

	if webinarID <= 0 {
		return Domain{}, usecase.ErrNotFound
	}
	res, err := webinarUC.webinarRepository.GetByID(ctx, webinarID)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (webinarUC *webinarUsecase) GetByName(ctx context.Context, webinarName string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, webinarUC.contextTimeout)
	defer cancel()

	if strings.TrimSpace(webinarName) == "" {
		return Domain{}, usecase.ErrNotFound
	}
	res, err := webinarUC.webinarRepository.GetByName(ctx, webinarName)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (webinarUC *webinarUsecase) Store(ctx context.Context, webinarDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, webinarUC.contextTimeout)
	defer cancel()

	existedWebinar, _ := webinarUC.webinarRepository.GetByName(ctx, webinarDomain.Name)

	if existedWebinar != (Domain{}) {
		return Domain{}, usecase.ErrDuplicateData
	}

	result, err := webinarUC.webinarRepository.Store(ctx, webinarDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (webinarUC *webinarUsecase) Update(ctx context.Context, webinarDomain *Domain) (*Domain, error) {
	existedWebinar, err := webinarUC.webinarRepository.GetByID(ctx, webinarDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	webinarDomain.ID = existedWebinar.ID

	result, err := webinarUC.webinarRepository.Update(ctx, webinarDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}

func (webinarUC *webinarUsecase) Delete(ctx context.Context, webinarDomain *Domain) (*Domain, error) {
	existedWebinar, err := webinarUC.webinarRepository.GetByID(ctx, webinarDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	webinarDomain.ID = existedWebinar.ID

	result, err := webinarUC.webinarRepository.Delete(ctx, webinarDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}
