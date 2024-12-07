package controllers

import (
	"bookedup/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsuarios(c *gin.Context) {
	usuarios, err := models.BuscarUsuarios()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, usuarios)
}

func PostUsuario(c *gin.Context) {
	var novoUsuario models.Usuario

	if err := c.BindJSON(&novoUsuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := models.CriarUsuario(novoUsuario)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	novoUsuario.ID = id

	c.JSON(http.StatusCreated, novoUsuario)
}

func DeleteUsuario(c *gin.Context) {
	id := c.Param("id")

	usuarioRemovido, err := models.RemoverUsuario(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "usuario removido com sucesso",
		"usuario": usuarioRemovido,
	})
}
