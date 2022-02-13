package main

import (
	"golang-experiment/handlers"
	"golang-experiment/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	modules := r.Group("/modules")
	{
		modules.GET("/", handlers.GetModules)
		modules.POST("/", middlewares.AuthMiddleware, middlewares.CreateModuleMiddleware(), handlers.CreateModule)
		modules.PUT("/:id", middlewares.AuthMiddleware, middlewares.EditModuleMiddleware(), handlers.UpdateModule)
		modules.DELETE("/:id", middlewares.AuthMiddleware, middlewares.DeleteModuleMiddleware(), handlers.DeleteModule)
	}

	r.Run()
}
