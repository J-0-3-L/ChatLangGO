package handlers

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	allRoutes := r.Group("/test")
	{
		allRoutes.GET("/", Test)
	}
}

func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Usando nueva ubicacion de rutas",
	})
}
