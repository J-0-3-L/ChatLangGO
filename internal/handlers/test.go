package handlers

import (
	"github.com/gin-gonic/gin"
)

// Test godoc
// @Summary Endpoint de prueba validada
// @Description Devuelve un mensaje simple para verificar el funcionamiento del servidor.
// @Tags Test
// @Produce json
// @Success 200 {object} map[string]string
// @Router /test/ [get]
func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Usando nueva ubicacion de rutas",
	})
}

func RegisterRoutes(r *gin.Engine) {
	allRoutes := r.Group("/test")
	{
		allRoutes.GET("/", Test)
	}
}
