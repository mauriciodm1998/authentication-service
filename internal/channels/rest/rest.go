package rest

import (
	"authentication-service/internal/config"
	"authentication-service/internal/middlewares"
	"authentication-service/internal/service"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type Login interface {
	RegisterGroup(*echo.Group)
	Login(echo.Context) error
	Start() error
}

type login struct {
	service service.LoginService
}

func NewRestChannel() Login {
	return &login{
		service: service.NewLoginService(),
	}
}

func (u *login) RegisterGroup(g *echo.Group) {
	g.POST("/login", u.Login)
	g.POST("/", u.Create)
}

func (u *login) Start() error {
	router := echo.New()

	router.Use(middlewares.Logger)

	mainGroup := router.Group("/api")

	authGroup := mainGroup.Group("/user")
	u.RegisterGroup(authGroup)

	return router.Start(":" + config.Get().Server.Port)
}

func (u *login) Login(c echo.Context) error {
	var userLogin LoginRequest

	if err := c.Bind(&userLogin); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, Response{
			Message: fmt.Errorf("invalid data").Error(),
		})
	}

	token, err := u.service.Login(c.Request().Context(), userLogin.toCanonical())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, Response{err.Error()})
	}

	return c.JSON(http.StatusOK, TokenResponse{
		Token: token,
	})
}

func (u *login) Create(c echo.Context) error {
	var userRequest CreateUserRequest

	if err := c.Bind(&userRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, Response{
			Message: fmt.Errorf("invalid data").Error(),
		})
	}

	err := u.service.CreateUser(c.Request().Context(), userRequest.toCanonical())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, Response{
			Message: err.Error(),
		})
	}

	return c.NoContent(http.StatusNoContent)
}
