package main

import (
	"apibe23/configs"
	"apibe23/internal/features/users/handler"
	"apibe23/internal/features/users/repository"
	"apibe23/internal/features/users/services"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	e := echo.New()

	cfg := configs.ImportSetting()
	db, _ := configs.ConnectDB(cfg)
	db.AutoMigrate(&repository.User{})
	um := repository.NewUserModel(db)
	us := services.NewUserService(um)
	uc := handler.NewUserController(us)

	// tm := models.NewTodoModel(db)
	// tc := todos.NewTodoController(tm)

	// Register
	e.POST("/users", uc.Register())
	e.POST("/login", uc.Login())

	t := e.Group("/todos")
	t.Use(echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte("passkeyJWT"),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	))
	// t.GET("", tc.ShowMyTodo())
	// t.POST("", tc.CreateTodo())

	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(200, "hello world")
	}, echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte("passkeyJWT"),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	))
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())

	e.Logger.Error(e.Start(":8000"))
}
