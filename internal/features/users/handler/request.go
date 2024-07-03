package handler

import (
	"apibe23/internal/features/users"
)

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type RegisterRequest struct {
	Name     string          `json:"name"`
	Password string          `json:"password"`
	Email    string          `json:"email"`
	Phone    string          `json:"hp"`
	Address  []AlamatRequest `json:"address"`
}

type AlamatRequest struct {
	Alamat string `json:"alamat"`
}

func ToModelUsers(r RegisterRequest) users.User {
	return users.User{
		Name:     r.Name,
		Password: r.Password,
		Email:    r.Email,
		Phone:    r.Phone,
	}
}
