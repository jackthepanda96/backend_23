package users

import "github.com/labstack/echo/v4"

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
}

type Services interface {
	Register(newUser User) error
	Login(email string, password string) (User, string, error)
}

type Query interface {
	Register(newUser User) error
	Login(email string) (User, error)
}
