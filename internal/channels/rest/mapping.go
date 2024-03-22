package rest

import "authentication-service/internal/canonical"

func (c *LoginRequest) toCanonical() canonical.Login {
	return canonical.Login{
		UserName:     c.UserName,
		Registration: c.Registration,
		Password:     c.Password,
	}
}

func (c *CreateUserRequest) toCanonical() canonical.User {
	return canonical.User{
		UserName:     c.UserName,
		Registration: c.Registration,
		Email:        c.Email,
		Password:     c.Password,
	}
}
