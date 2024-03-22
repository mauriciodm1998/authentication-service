package rest

type LoginRequest struct {
	UserName     string `json:"user_name"`
	Registration string `json:"registration"`
	Password     string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type Response struct {
	Message string `json:"message"`
}

type CreateUserRequest struct {
	UserName     string `json:"user_name"`
	Registration string `json:"registration"`
	Password     string `json:"password"`
	Email        string `json:"email"`
}
