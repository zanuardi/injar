package users_test

import (
	"context"
	"errors"
	"injar/app/middleware"
	"injar/helpers/encrypt"
	businesses "injar/usecase"
	"injar/usecase/users"
	userMock "injar/usecase/users/mocks"

	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userRepository userMock.Repository
	userUsecase    users.Usecase
	jwtAuth        *middleware.ConfigJWT
)

func setup() {
	jwtAuth = &middleware.ConfigJWT{SecretJWT: "secret", ExpiresDuration: 2}
	userUsecase = users.NewUserUsecase(&userRepository, jwtAuth, 2)

}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())

}

func TestGetById(t *testing.T) {
	t.Run("test case 1, valid test ", func(t *testing.T) {
		domain := users.Domain{
			ID:       1,
			Name:     "Zanuardi Novanda",
			Username: "zanuardi",
			Email:    "zanuardi@gmail.com",
		}
		userRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(domain, nil).Once()

		result, err := userUsecase.GetByID(context.Background(), 1)

		assert.Equal(t, err, nil)
		assert.Equal(t, result, domain)
	})

	// t.Run("test case 2, invalid id", func(t *testing.T) {
	// 	userRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(users.Domain{}, businesses.ErrNotFound).Once()
	// 	result, err := userUsecase.GetByID(context.Background(), -1)

	// 	assert.Equal(t, result, users.Domain{})
	// 	assert.Equal(t, err, businesses.ErrNotFound)
	// })

}

func TestRegister(t *testing.T) {
	t.Run("test case 1, valid test ", func(t *testing.T) {
		domain := users.Domain{
			ID:       1,
			Name:     "Zanuardi Novanda",
			Username: "zanuardi",
			Email:    "zanuardi@gmail.com",
		}
		userRepository.On("GetByUsername", mock.Anything, mock.AnythingOfType("string")).Return(users.Domain{}, nil).Once()
		userRepository.On("Store", mock.Anything, mock.AnythingOfType("*users.Domain")).Return(nil).Once()

		err := userUsecase.Store(context.Background(), &domain)

		assert.NoError(t, err)
	})

	t.Run("test case 2, duplicate data", func(t *testing.T) {
		domain := users.Domain{
			ID:       1,
			Name:     "Zanuardi Novanda",
			Username: "zanuardi",
			Email:    "zanuardi@gmail.com",
		}
		errRepository := errors.New("duplicate data")
		userRepository.On("GetByUsername", mock.Anything, mock.AnythingOfType("string")).Return(domain, errRepository).Once()

		err := userUsecase.Store(context.Background(), &domain)

		assert.Equal(t, err, errRepository)
	})

	t.Run("test case 3, duplicate data", func(t *testing.T) {
		domain := users.Domain{
			ID:       1,
			Name:     "Zanuardi Novanda",
			Username: "zanuardi",
			Email:    "zanuardi@gmail.com",
		}
		userRepository.On("GetByUsername", mock.Anything, mock.AnythingOfType("string")).Return(domain, businesses.ErrDuplicateData).Once()

		err := userUsecase.Store(context.Background(), &domain)

		assert.Equal(t, err, businesses.ErrDuplicateData)
	})

	t.Run("test case 4, duplicate data", func(t *testing.T) {

		userRepository.On("GetByUsername", mock.Anything, mock.AnythingOfType("string")).Return(users.Domain{}, businesses.ErrInternalServer).Once()

		err := userUsecase.Store(context.Background(), &users.Domain{})

		assert.Equal(t, err, businesses.ErrInternalServer)
	})

	t.Run("test case 5, register failed", func(t *testing.T) {
		domain := users.Domain{
			ID:       1,
			Name:     "Zanuardi Novanda",
			Username: "zanuardi",
			Email:    "zanuardi@gmail.com",
		}
		errRepository := errors.New("register failed")
		userRepository.On("GetByUsername", mock.Anything, mock.AnythingOfType("string")).Return(users.Domain{}, nil).Once()
		userRepository.On("Store", mock.Anything, mock.AnythingOfType("*users.Domain")).Return(errRepository).Once()

		err := userUsecase.Store(context.Background(), &domain)

		assert.Equal(t, err, errRepository)
	})

}

func TestLogin(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		pass, _ := encrypt.Hash("12345")
		usersDomain := users.Domain{
			ID:       1,
			Name:     "Zanuardi Novanda",
			Username: "zanuardi",
			Email:    "zanuardi@gmail.com",
			Password: pass,
		}

		userRepository.On("GetByUsername", mock.Anything, mock.AnythingOfType("string")).Return(usersDomain, nil).Once()

		_, err := userUsecase.CreateToken(context.Background(), "zanuardi", "12345")
		assert.Nil(t, err)
	})
	t.Run("test case 2, password error", func(t *testing.T) {
		pass, _ := encrypt.Hash("12345")
		usersDomain := users.Domain{
			ID:       1,
			Name:     "Zanuardi Novanda",
			Username: "zanuardi",
			Email:    "zanuardi@gmail.com",
			Password: pass,
		}

		userRepository.On("GetByUsername", mock.Anything, mock.AnythingOfType("string")).Return(usersDomain, nil).Once()

		_, err := userUsecase.CreateToken(context.Background(), "zanuardi", "1234567")
		assert.Equal(t, err, businesses.ErrInternalServer)

	})

	t.Run("test case 3, error record", func(t *testing.T) {

		errRepository := errors.New("error record")
		userRepository.On("GetByUsername", mock.Anything, mock.AnythingOfType("string")).Return(users.Domain{}, errRepository).Once()

		result, err := userUsecase.CreateToken(context.Background(), "zanuard", "12345")

		assert.Equal(t, err, errRepository)
		assert.Equal(t, result, result)
	})

}

func TestUpdate(t *testing.T) {
	t.Run("test case 1, valid test ", func(t *testing.T) {
		domain := users.Domain{
			ID:       1,
			Name:     "Zanuardi Novanda",
			Username: "zanuardi",
			Email:    "zanuardi@gmail.com",
		}
		userRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(domain, nil).Once()
		userRepository.On("Update", mock.Anything, mock.AnythingOfType("*users.Domain")).Return(users.Domain{}, nil).Once()

		_, err := userUsecase.Update(context.Background(), &domain)

		assert.NoError(t, err)
	})

	t.Run("test case 2, duplicate data", func(t *testing.T) {
		domain := users.Domain{
			ID:       1,
			Name:     "Zanuardi Novanda",
			Username: "zanuardi",
			Email:    "zanuardi@gmail.com",
		}
		userRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(domain, nil).Once()
		userRepository.On("GetByUsername", mock.Anything, mock.AnythingOfType("string")).Return(users.Domain{}, businesses.ErrInternalServer).Once()
		userRepository.On("Update", mock.Anything, mock.AnythingOfType("*users.Domain")).Return(users.Domain{}, nil).Once()

		res, err := userUsecase.Update(context.Background(), &users.Domain{})

		assert.Equal(t, res, res)

		assert.Equal(t, err, nil)
	})
}
