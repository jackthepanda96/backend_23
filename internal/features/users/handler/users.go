package handler

import (
	"apibe23/internal/features/users"
	"apibe23/internal/helper"

	"fmt"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	srv users.Services
}

func NewUserController(s users.Services) users.Handler {
	return &UserController{
		srv: s,
	}
}

func (uc *UserController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RegisterRequest
		err := c.Bind(&input)
		fmt.Println(input.Address)
		if err != nil {
			return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
		}

		err = uc.srv.Register(ToModelUsers(input))

		if err != nil {
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}

		return c.JSON(201, helper.ResponseFormat(201, "success insert data", nil))
	}
}

func (uc *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginRequest
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
		}

		result, token, err := uc.srv.Login(input.Email, input.Password)

		if err != nil {
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}

		return c.JSON(200, helper.ResponseFormat(200, "success login", ToLoginReponse(result, token)))
	}
}

// func GetAllUsers(c echo.Context) error {
// 	if len(daftarUser) == 0 {
// 		return c.JSON(404, "empty")
// 	}

// 	return c.JSON(200, daftarUser)
// }
