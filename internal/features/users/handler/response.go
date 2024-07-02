package handler

import (
	"apibe23/internal/features/users"
)

type LoginResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"hp"`
	Token    string `json:"token"`
}

func ToLoginReponse(input users.User, tkn string) LoginResponse {
	return LoginResponse{
		ID:       input.ID,
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Phone:    input.Phone,
		Token:    tkn,
	}
}
