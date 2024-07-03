package factory

import (
	"apibe23/configs"
	repoTodo "apibe23/internal/features/todos/repository"
	"apibe23/internal/features/users/handler"
	"apibe23/internal/features/users/repository"
	"apibe23/internal/features/users/services"
	"apibe23/internal/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitFactory(e *echo.Echo) {
	cfg := configs.ImportSetting()
	db, _ := configs.ConnectDB(cfg)
	db.AutoMigrate(&repository.User{}, &repoTodo.Todo{})
	um := repository.NewUserModel(db)
	us := services.NewUserService(um)
	uc := handler.NewUserController(us)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())

	// Register

	// t.GET("", tc.ShowMyTodo())
	// t.POST("", tc.CreateTodo())

	routes.InitRoute(e, uc, nil)
}
