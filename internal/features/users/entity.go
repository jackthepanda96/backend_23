package users

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID       uint
	Name     string
	Password string
	Email    string
	Phone    string
}

type Handler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	Register2() echo.HandlerFunc
}

type Services interface {
	Register(newUser User) error
	Login(email string, password string) (User, string, error)
	Register2(newUser string, file *multipart.FileHeader) error
}

type Query interface {
	Register(newUser User) error
	Login(email string) (User, error)
}

type LoginValidate struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6,alphanum"`
}
