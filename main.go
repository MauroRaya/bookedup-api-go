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
	r.POST("/usuarios", controllers.PostUsuario)
	r.DELETE("/usuarios/:id", controllers.DeleteUsuario)

	r.Run("localhost:80")
}
