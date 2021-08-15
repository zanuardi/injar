package users

import (
	"context"
	"injar/app/middleware"
	"injar/helpers/encrypt"
	usecase "injar/usecase"
	"strings"
	"time"
)

type userUsecase struct {
	userRepository Repository
	contextTimeout time.Duration
	jwtAuth        *middleware.ConfigJWT
}

func NewUserUsecase(ur Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration) Usecase {
	return &userUsecase{
		userRepository: ur,
		jwtAuth:        jwtauth,
		contextTimeout: timeout,
	}
}

func (uc *userUsecase) CreateToken(ctx context.Context, username, password string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if strings.TrimSpace(username) == "" && strings.TrimSpace(password) == "" {
		return "", usecase.ErrUsernamePasswordNotFound
	}

	userDomain, err := uc.userRepository.GetByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	if !encrypt.ValidateHash(password, userDomain.Password) {
		return "", usecase.ErrInternalServer
	}

	token := uc.jwtAuth.GenerateToken(userDomain.ID)
	return token, nil
}

func (uc *userUsecase) Store(ctx context.Context, userDomain *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedUser, err := uc.userRepository.GetByUsername(ctx, userDomain.Username)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return err
		}
	}
	if existedUser != (Domain{}) {
		return usecase.ErrDuplicateData
	}

	userDomain.Password, err = encrypt.Hash(userDomain.Password)
	if err != nil {
		return usecase.ErrInternalServer
	}
	err = uc.userRepository.Store(ctx, userDomain)
	if err != nil {
		return err
	}

	return nil
}
