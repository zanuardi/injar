package categories_test

import (
	"context"
	"errors"
	businesses "injar/usecase"
	categoryMock "injar/usecase/categories/mocks"

	"os"
	"testing"

	category "injar/usecase/categories"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	categoryRepository categoryMock.Repository
	categoryUsecase    category.Usecase
)

func setup() {
	categoryUsecase = category.NewCategoryUsecase(2, &categoryRepository)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())

}

func TestGetById(t *testing.T) {
	t.Run("test case 1, valid test ", func(t *testing.T) {
		domain := category.Domain{
			ID:   1,
			Name: "Teknologi",
		}
		categoryRepository.On("GetByID", context.Background(), mock.AnythingOfType("int")).Return(domain, nil).Once()

		result, err := categoryRepository.GetByID(context.Background(), 1)

		assert.Nil(t, err)

		assert.Equal(t, domain.ID, result.ID)
		assert.Equal(t, domain.Name, result.Name)
	})

	t.Run("test case 2, invalid id", func(t *testing.T) {
		result, err := categoryUsecase.GetByID(context.Background(), -1)

		assert.Equal(t, result, category.Domain{})
		assert.Equal(t, err, businesses.ErrCategoryNotFound)
	})

}

func TestGetByName(t *testing.T) {
	t.Run("test case 1, valid test ", func(t *testing.T) {
		domain := category.Domain{
			ID:   1,
			Name: "Teknologi",
		}
		categoryRepository.On("GetByName", context.Background(), mock.AnythingOfType("string")).Return(domain, nil).Once()

		result, err := categoryRepository.GetByName(context.Background(), "Teknologi")

		assert.Nil(t, err)

		assert.Equal(t, domain.ID, result.ID)
		assert.Equal(t, domain.Name, result.Name)
	})

	t.Run("test case 2, invalid name", func(t *testing.T) {
		result, err := categoryUsecase.GetByName(context.Background(), "")

		assert.Equal(t, result, category.Domain{})
		assert.Equal(t, err, businesses.ErrCategoryNotFound)
	})

}

func TestGetAll(t *testing.T) {
	t.Run("test case 1, get all", func(t *testing.T) {
		domain := []category.Domain{
			{
				ID:   1,
				Name: "Teknologi",
			},
			{
				ID:   2,
				Name: "Bisnis",
			},
		}
		categoryRepository.On("Find", context.Background()).Return(domain, nil).Once()

		result, err := categoryUsecase.GetAll(context.Background())

		assert.Equal(t, 2, len(result))
		assert.Nil(t, err)
	})

	t.Run("test case 2, repository error", func(t *testing.T) {
		errRepository := errors.New("mysql not running")
		categoryRepository.On("Find", context.Background()).Return([]category.Domain{}, errRepository).Once()

		result, err := categoryUsecase.GetAll(context.Background())

		assert.Equal(t, 0, len(result))
		assert.Equal(t, errRepository, err)
	})
}

func TestStore(t *testing.T) {

	t.Run("test case 2, duplicate data", func(t *testing.T) {
		domain := category.Domain{
			ID:   1,
			Name: "hiking",
		}
		errRepository := errors.New("duplicate data")
		categoryRepository.On("GetByName", mock.Anything, mock.AnythingOfType("string")).Return(domain, errRepository).Once()

		_, err := categoryUsecase.Store(context.Background(), &domain)

		assert.Equal(t, err, businesses.ErrDuplicateData)
	})

}
