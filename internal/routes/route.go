package routes

import (
	"apibe23/internal/features/todos"
	"apibe23/internal/features/users"

	"github.com/golang-jwt/jwt"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo, uc users.Handler, tc todos.Handler) {
	e.POST("/users", uc.Register())
	e.POST("/login", uc.Login())
	e.POST("/register", uc.Register2())
	todoRoute(e, tc)
}

func todoRoute(e *echo.Echo, tc todos.Handler) {
	t := e.Group("/todos")
	t.Use(echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte("passkeyJWT"),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	))
}
