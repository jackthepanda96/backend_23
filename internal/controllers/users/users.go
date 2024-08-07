package users

import (
	"apibe23/internal/helper"
	"apibe23/internal/models"
	"apibe23/internal/utils"
	"fmt"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	model *models.UserModel
}

func NewUserController(m *models.UserModel) *UserController {
	return &UserController{
		model: m,
	}
}

func (uc *UserController) Register(c echo.Context) error {
	var input RegisterRequest
	err := c.Bind(&input)
	fmt.Println(input.Address)
	if err != nil {
		return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
	}

	processPw, err := utils.GeneratePassword(input.Password)

	if err != nil {
		return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
	}

	input.Password = string(processPw)

	_, err = uc.model.Register(ToModelUsers(input))
	if err != nil {
		return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
	}
	return c.JSON(201, helper.ResponseFormat(201, "success insert data", nil))
}

func (uc *UserController) Login(c echo.Context) error {
	var input LoginRequest
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
	}

	result, err := uc.model.Login(input.Email)
	if err != nil {
		return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
	}

	err = utils.CheckPassword([]byte(input.Password), []byte(result.Password))

	if err != nil {
		return c.JSON(400, helper.ResponseFormat(400, "input data kurang tepat", nil))
	}

	token, err := utils.GenerateToken(result.ID)
	if err != nil {
		return c.JSON(500, helper.ResponseFormat(500, "privacy error", nil))
	}

	return c.JSON(200, helper.ResponseFormat(200, "success login", ToLoginReponse(result, token)))
}

// func GetAllUsers(c echo.Context) error {
// 	if len(daftarUser) == 0 {
// 		return c.JSON(404, "empty")
// 	}

// 	return c.JSON(200, daftarUser)
// }
