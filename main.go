package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// fmt.Println("lo logramos")

	// Crear router
	r := gin.Default()

	// Ruta simple
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Probando Gin en Backend",
		})
	})

	// Iniciar el server
	r.Run(":8080")
}
