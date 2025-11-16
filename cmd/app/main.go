package main

import (
	"go000/internal/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

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

	// Rutas
	handlers.RegisterRoutes(r)

	// Iniciar el server
	r.Run(":8080")
}
