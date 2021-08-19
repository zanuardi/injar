package favourites_test

import (
	"context"
	businesses "injar/usecase"
	favouriteMock "injar/usecase/favourites/mocks"

	"os"
	"testing"

	favourite "injar/usecase/favourites"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	favouriteRepository favouriteMock.Repository
	favouriteUsecase    favourite.Usecase
)

func setup() {
	favouriteUsecase = favourite.NewFavouritesUsecase(2, &favouriteRepository)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())

}

func TestGetById(t *testing.T) {
	t.Run("test case 1, valid test ", func(t *testing.T) {
		domain := favourite.Domain{
			ID:        1,
			UserID:    1,
			WebinarID: 1,
		}
		favouriteRepository.On("GetByID", context.Background(), mock.AnythingOfType("int")).Return(domain, nil).Once()

		result, err := favouriteRepository.GetByID(context.Background(), 1)

		assert.Nil(t, err)

		assert.Equal(t, domain.ID, result.ID)
		assert.Equal(t, domain.ID, result.ID)
	})

	t.Run("test case 2, invalid id", func(t *testing.T) {
		result, err := favouriteUsecase.GetByID(context.Background(), -1)

		assert.Equal(t, result, favourite.Domain{})
		assert.Equal(t, err, businesses.ErrNotFound)
	})

}
