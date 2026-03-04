package app

import (
	"backend/api"

	"github.com/gin-gonic/gin"
)

type App struct {
	router         *gin.Engine
	userController api.UserControllerInterface
}

func NewApp() *App {
	return &App{
		router: gin.Default(),
	}
}

func (a *App) Run() {
	a.router.POST("/api/login", a.userController.Login)
	a.router.Run(":8001")
}
