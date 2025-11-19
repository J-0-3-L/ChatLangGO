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

// @Summary Obtener usuario por ID
// @Description Obtiene los detalles de un usuario especificando su ID
// @Accept json
// @Produce json
// @Param id path int true "ID del usuario"
// @Success 200 {object} User
// @Failure 404 {object} ErrorResponse
// @Router /user/{id} [get]
func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Usando nueva ubicacion de rutas",
	})
}
