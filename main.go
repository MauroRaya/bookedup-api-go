package main

import (
	"bookedup/config"
	"bookedup/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	r := gin.Default()

	r.GET("/usuarios", controllers.GetUsuarios)
	r.GET("/usuarios/:id", controllers.GetUsuario)
	r.POST("/usuarios", controllers.PostUsuario)
	r.PATCH("/usuarios/:id", controllers.PatchUsuario)
	r.DELETE("/usuarios/:id", controllers.DeleteUsuario)

	r.Run("localhost:80")
}
