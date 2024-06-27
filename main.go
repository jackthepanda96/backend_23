package main

import (
	"apibe23/configs"
	"apibe23/internal/controllers/users"
	"apibe23/internal/models"

	"github.com/labstack/echo/v4"
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
	um := models.NewUserModel(db)
	uc := users.NewUserController(um)
	// Register
	e.POST("/users", uc.Register)
	e.POST("/login", uc.Login)
	// Login
	// Tampilkan semua data
	// e.GET("/users", GetAllUsers)
	e.Start(":5000")
}
