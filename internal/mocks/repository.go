package mocks

import (
	"authentication-service/internal/canonical"
	"context"

	"github.com/stretchr/testify/mock"
)

type LoginRepositoryMock struct {
	mock.Mock
}

func (l *LoginRepositoryMock) GetUser(ctx context.Context, login canonical.Login) (*canonical.User, error) {
	args := l.Called(login)

	return args.Get(0).(*canonical.User), args.Error(1)
}

func (l *LoginRepositoryMock) MockGetUser(userName string, user canonical.User, errorReturned error, times int) {
	l.On("GetUser", mock.MatchedBy(func(login canonical.Login) bool {
		return login.UserName == userName
	})).Return(&user, errorReturned).Times(times)
}
