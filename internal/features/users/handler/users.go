package handler

import (
	"apibe23/internal/features/users"
	"apibe23/internal/helper"
	"net/http"
	"strings"

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
		if err != nil {
			c.Logger().Error("register parse error:", err.Error())
			return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
		}

		err = uc.srv.Register(ToModelUsers(input))

		if err != nil {
			errCode := 500
			if strings.ContainsAny(err.Error(), "tidak valid") {
				errCode = 400
			}
			return c.JSON(errCode, helper.ResponseFormat(errCode, err.Error(), nil))
		}

		return c.JSON(201, helper.ResponseFormat(201, "success insert data", nil))
	}
}

func (uc *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginRequest
		err := c.Bind(&input)
		if err != nil {
			c.Logger().Error("login parse error:", err.Error())
			return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
		}

		result, token, err := uc.srv.Login(input.Email, input.Password)

		if err != nil {
			errCode := 500
			if strings.ContainsAny(err.Error(), "tidak ditemukan") {
				errCode = 400
			}
			return c.JSON(errCode, helper.ResponseFormat(errCode, err.Error(), nil))
		}

		return c.JSON(200, helper.ResponseFormat(200, "success login", ToLoginReponse(result, token)))
	}
}

func (uc *UserController) Register2() echo.HandlerFunc {

	return func(c echo.Context) error {
		// Get name
		name := c.FormValue("name")
		// Get avatar
		avatar, err := c.FormFile("avatar")
		if err != nil {
			return err
		}

		err = uc.srv.Register2(name, avatar)
		if err != nil {
			return err
		}

		return c.HTML(http.StatusOK, "<b>Thank you! "+name+"</b>")
	}
	// Get name and email

}

// func GetAllUsers(c echo.Context) error {
// 	if len(daftarUser) == 0 {
// 		return c.JSON(404, "empty")
// 	}

// 	return c.JSON(200, daftarUser)
// }
