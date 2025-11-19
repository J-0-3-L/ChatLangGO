package main

import (
	"go000/internal/config"
	"go000/internal/handlers"

	_ "go000/cmd/app/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// @title Example API
// @version 1.0
// @description A simple API to demonstrate Swagger with Gin
// @host localhost:8080
// @BasePath /api/v1

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

	// Config de swagger
	config.ConfigSwagger(r)

	// Rutas
	handlers.RegisterRoutes(r)

	// Iniciar el server
	r.Run(":8080")
}
