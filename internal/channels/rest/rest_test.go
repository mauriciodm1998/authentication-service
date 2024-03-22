package rest

import (
	"authentication-service/internal/canonical"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (c *MockService) Login(_ context.Context, login canonical.Login) (string, error) {
	args := c.Called(login)
	return args.Get(0).(string), args.Error(1)
}

func (c *MockService) CreateUser(ctx context.Context, user canonical.User) error {
	args := c.Called()

	return args.Error(0)
}

func (c *MockService) MockLogin(login canonical.Login, token string, errorToReturn error, times int) {
	c.On("Login", mock.MatchedBy(func(l canonical.Login) bool {
		return l.UserName == login.UserName
	})).Return(token, errorToReturn).Times(times)
}

var (
	svc *MockService

	rest Login
)

func init() {
	svc = new(MockService)

	rest = &login{
		service: svc,
	}
}

func TestStart(t *testing.T) {
	go func() {
		err := rest.Start()
		assert.NoError(t, err)
	}()

	<-time.After(100 * time.Millisecond)
}

func TestLogin(t *testing.T) {
	login := canonical.Login{
		UserName: "loginfulano",
		Password: "12345",
	}

	svc.MockLogin(login, "fake-token", nil, 1)

	request := LoginRequest{
		UserName: "loginfulano",
		Password: "12345",
	}

	req := createJsonRequest(http.MethodPost, "/login", request)

	rec := httptest.NewRecorder()

	err := rest.Login(echo.New().NewContext(req, rec))

	assert.Equal(t, http.StatusOK, rec.Result().StatusCode)
	assert.Nil(t, err)
	svc.AssertExpectations(t)
}

func TestLoginError(t *testing.T) {
	login := canonical.Login{
		UserName: "loginfulano",
		Password: "12345",
	}

	svc.MockLogin(login, "fake-token", errors.New("generic error"), 1)

	request := LoginRequest{
		UserName: "loginfulano",
		Password: "12345",
	}

	req := createJsonRequest(http.MethodPost, "/login", request)

	rec := httptest.NewRecorder()

	err := rest.Login(echo.New().NewContext(req, rec))

	assert.Error(t, err)
	svc.AssertExpectations(t)
}

func TestLoginErrorBind(t *testing.T) {
	req := createJsonRequest(http.MethodPost, "/login", "")

	rec := httptest.NewRecorder()

	err := rest.Login(echo.New().NewContext(req, rec))

	assert.Equal(t, "code=400, message={invalid data}", err.Error())
	svc.AssertExpectations(t)
}

func createJsonRequest(method, endpoint string, request interface{}) *http.Request {
	json, _ := json.Marshal(request)
	req := httptest.NewRequest(method, endpoint, bytes.NewReader(json))
	req.Header.Set("Content-Type", "application/json")
	return req
}
