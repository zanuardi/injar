package categories

import (
	"context"
	"injar/businesses"
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

func (cu *categoryUsecase) Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()

	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 10
	}

	res, total, err := cu.categoryRepository.Fetch(ctx, page, perpage)
	if err != nil {
		return []Domain{}, 0, err
	}

	return res, total, nil
}

func (cu *categoryUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	resp, err := cu.categoryRepository.Find(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return resp, nil
}

func (cu *categoryUsecase) GetByID(ctx context.Context, categoryID int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()

	if categoryID <= 0 {
		return Domain{}, businesses.ErrCategoryNotFound
	}
	res, err := cu.categoryRepository.GetByID(ctx, categoryID)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (cu *categoryUsecase) GetByName(ctx context.Context, categoryName string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()

	if strings.TrimSpace(categoryName) == "" {
		return Domain{}, businesses.ErrCategoryNotFound
	}
	res, err := cu.categoryRepository.GetByName(ctx, categoryName)
	if err != nil {
		return Domain{}, err
	}

	return res, nil
}

func (cu *categoryUsecase) Store(ctx context.Context, categoryDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()

	existedCategories, _ := cu.categoryRepository.GetByName(ctx, categoryDomain.Name)

	if existedCategories != (Domain{}) {
		return Domain{}, businesses.ErrDuplicateData
	}

	// now := time.Now().UTC()
	result := Domain{
		Name:      categoryDomain.Name,
		UpdatedAt: time.Now(),
	}

	result, err := cu.categoryRepository.Store(ctx, &result)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (cu *categoryUsecase) Update(ctx context.Context, categoriesDomain *Domain) (*Domain, error) {
	existedCategories, err := cu.categoryRepository.GetByID(ctx, categoriesDomain.ID)
	if err != nil {
		return &Domain{}, err
	}
	categoriesDomain.ID = existedCategories.ID

	result, err := cu.categoryRepository.Update(ctx, categoriesDomain)
	if err != nil {
		return &Domain{}, err
	}

	return &result, nil
}
