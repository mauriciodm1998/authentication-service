package service

import (
	"authentication-service/internal/canonical"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/undefinedlabs/go-mpatch"
	"golang.org/x/crypto/bcrypt"
)

type LoginRepositoryMock struct {
	mock.Mock
}

func (l *LoginRepositoryMock) GetUser(ctx context.Context, login canonical.Login) (*canonical.User, error) {
	args := l.Called(login)

	return args.Get(0).(*canonical.User), args.Error(1)
}

func (l *LoginRepositoryMock) CreateUser(ctx context.Context, user canonical.User) error {
	args := l.Called()

	return args.Error(0)
}

func (l *LoginRepositoryMock) MockGetUser(userName string, user canonical.User, errorReturned error, times int) {
	l.On("GetUser", mock.MatchedBy(func(login canonical.Login) bool {
		return login.UserName == userName
	})).Return(&user, errorReturned).Times(times)
}

var (
	service LoginService

	repositoryMock *LoginRepositoryMock
)

func init() {
	repositoryMock = new(LoginRepositoryMock)
	service = &loginService{
		Repository: repositoryMock,
	}
}

func TestLogin(t *testing.T) {
	login := canonical.Login{
		UserName: "fulano",
		Password: "passwordteste123",
	}

	user := canonical.User{
		Id:       12321321,
		UserName: "fulano",
		Password: "$2a$10$GjN8aPVbp5u/jFZFwmgda.XpLnj7oqb6hPsN1v57JhbCIqN/M.04O",
	}

	repositoryMock.MockGetUser(login.UserName, user, nil, 1)

	token, err := service.Login(context.Background(), login)

	assert.Nil(t, err)
	assert.NotNil(t, token)
	repositoryMock.AssertExpectations(t)
}

func TestLoginGetUserError(t *testing.T) {
	login := canonical.Login{
		UserName: "fulano",
		Password: "passwordteste123",
	}

	user := canonical.User{
		Id:       12312321,
		UserName: "fulano",
		Password: "$2a$10$GjN8aPVbp5u/jFZFwmgda.XpLnj7oqb6hPsN1v57JhbCIqN/M.04O",
	}

	repositoryMock.MockGetUser(login.UserName, user, errors.New("generic error"), 1)

	_, err := service.Login(context.Background(), login)

	assert.Equal(t, errors.New("generic error"), err)
}

func TestLoginCheckPasswordError(t *testing.T) {
	login := canonical.Login{
		UserName: "fulano",
		Password: "passwordteste123",
	}

	user := canonical.User{
		Id:       12312321,
		UserName: "fulano",
		Password: "$2a$10$GjN8aPVbp5u/jFZFwmg/M.04O",
	}

	patch, _ := mpatch.PatchMethod(bcrypt.CompareHashAndPassword, func(hashedPassword []byte, password []byte) error {
		return errors.New("generic error")
	})

	defer patch.Unpatch()

	repositoryMock.MockGetUser(login.UserName, user, nil, 1)

	_, err := service.Login(context.Background(), login)

	assert.Equal(t, errors.New("generic error"), err)
}
