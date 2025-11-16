package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// fmt.Println("lo logramos")

	// Crear router
	r := gin.Default()

	// Uso de CORS
	r.Use(cors.Default()) // All origins allowed by default

	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"}, // Permite todos los orígenes (puedes especificar dominios en lugar de "*")
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// }))

	// Ruta simple
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Probando nueva estructura",
		})
	})

	// Iniciar el server
	r.Run(":8080")
}
