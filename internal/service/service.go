package service

import (
	"authentication-service/internal/canonical"
	"authentication-service/internal/repositories"
	"authentication-service/internal/security"
	"authentication-service/internal/token"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/sirupsen/logrus"
)

type LoginService interface {
	Login(context.Context, canonical.Login) (string, error)
	CreateUser(ctx context.Context, user canonical.User) error
}

type loginService struct {
	repositories.Repository
}

func NewLoginService() LoginService {
	return &loginService{
		repositories.New(),
	}
}

func (u *loginService) Login(ctx context.Context, user canonical.Login) (string, error) {
	baseUser, err := u.Repository.GetUser(ctx, user)
	if err != nil {
		log.Err(err).Msg("an error occurred when get user")
		return "", err
	}

	if err = security.CheckPassword(baseUser.Password, user.Password); err != nil {
		log.Err(err).Msg("an error occurred when check user password")
		return "", err
	}

	token, err := token.GenerateToken(baseUser.Id, baseUser.Email)
	if err != nil {
		log.Err(err).Msg("an error occurred when generate token")
		return "", err
	}

	return token, nil
}

func (u *loginService) CreateUser(ctx context.Context, user canonical.User) error {
	passEncrypted, err := security.Hash(user.Password)
	if err != nil {
		err = fmt.Errorf("error generating password hash: %w", err)
		logrus.WithError(err).Warn()
		return err
	}

	user.Password = string(passEncrypted)
	user.Id = uuid.New().String()
	err = u.Repository.CreateUser(ctx, user)
	if err != nil {
		return nil
	}

	return nil
}
