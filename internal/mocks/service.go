package mocks

import (
	"authentication-service/internal/canonical"
	"context"

	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (c *MockService) Login(_ context.Context, login canonical.Login) (string, error) {
	args := c.Called(login)
	return args.Get(0).(string), args.Error(1)
}

func (c *MockService) MockLogin(login canonical.Login, token string, errorToReturn error, times int) {
	c.On("Login", mock.MatchedBy(func(l canonical.Login) bool {
		return l.UserName == login.UserName
	})).Return(token, errorToReturn).Times(times)
}
