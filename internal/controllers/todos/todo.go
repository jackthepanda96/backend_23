package todos

import (
	"apibe23/internal/helper"
	"apibe23/internal/models"
	"apibe23/internal/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type TodoController struct {
	model *models.TodoModel
}

func NewTodoController(m *models.TodoModel) *TodoController {
	return &TodoController{
		model: m,
	}
}

func (tc *TodoController) CreateTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		var userID = utils.DecodeToken(c.Get("user").(*jwt.Token))

		var input TodoRequest

		err := c.Bind(&input)

		if err != nil {
			return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
		}

		err = tc.model.InsertTodo(ToModelTodo(input, userID))
		if err != nil {
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}

		return c.JSON(201, helper.ResponseFormat(201, "success insert todo", nil))
	}
}

func (tc *TodoController) ShowMyTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		var userID = utils.DecodeToken(c.Get("user").(*jwt.Token))

		result, err := tc.model.GetTodoByUser(userID)
		if err != nil {
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}

		return c.JSON(200, helper.ResponseFormat(200, "success get my todo", result))
	}
}
