package categories

import (
	"context"
	usecase "injar/usecase"
	"strings"
	"time"
)

type categoryUsecase struct {
	categoryRepository Repository
	contextTimeout     time.Duration
}

func NewCategoryUsecase(timeout time.Duration, cr Repository) Usecase {
	return &categoryUsecase{
		contextTimeout:     timeout,
		categoryRepository: cr,
	}
}

func (uc *categoryUsecase) FindAll(ctx context.Context, page, perpage int) ([]Domain, int, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 25
	}

	res, total, err := uc.categoryRepository.FindAll(ctx, page, perpage)
	if err != nil {
		return []Domain{}, 0, err
	}

	return res, total, nil
}

func (uc *categoryUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	resp, err := uc.categoryRepository.Find(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return resp, nil
}

func (uc *categoryUsecase) GetByID(ctx context.Context, categoryID int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if categoryID <= 0 {
		return Domain{}, usecase.ErrCategoryNotFound
	}
	res, err := uc.categoryRepository.GetByID(ctx, categoryID)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (uc *categoryUsecase) GetByName(ctx context.Context, categoryName string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if strings.TrimSpace(categoryName) == "" {
		return Domain{}, usecase.ErrCategoryNotFound
	}
	res, err := uc.categoryRepository.GetByName(ctx, categoryName)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (uc *categoryUsecase) Store(ctx context.Context, categoryDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedCategories, _ := uc.categoryRepository.GetByName(ctx, categoryDomain.Name)

	if existedCategories != (Domain{}) {
		return Domain{}, usecase.ErrDuplicateData
	}

	result, err := uc.categoryRepository.Store(ctx, categoryDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (uc *categoryUsecase) Update(ctx context.Context, categoriesDomain *Domain) (*Domain, error) {
	existedCategories, err := uc.categoryRepository.GetByID(ctx, categoriesDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	categoriesDomain.ID = existedCategories.ID

	result, err := uc.categoryRepository.Update(ctx, categoriesDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}

func (uc *categoryUsecase) Delete(ctx context.Context, categoriesDomain *Domain) (*Domain, error) {
	existedCategories, err := uc.categoryRepository.GetByID(ctx, categoriesDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	categoriesDomain.ID = existedCategories.ID

	result, err := uc.categoryRepository.Delete(ctx, categoriesDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}
