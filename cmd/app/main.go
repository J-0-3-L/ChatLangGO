package main

import (
	"go000/internal/auth"
	"go000/internal/config"
	"go000/internal/handlers"
	"go000/internal/posts"

	_ "go000/internal/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// @title ChatLanGO API
// @version 1.0
// @description Una app de red social simple de chateo
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Ingresa el token JWT: Bearer {token}
// @host localhost:8080
// @BasePath /

func main() {

	config.ConnectionDB()

	config.MigrateModels()

	// Crear router
	r := gin.Default()

	// Uso de CORS
	//r.Use(cors.Default()) // All origins allowed by default

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Permite todos los orígenes (puedes especificar dominios en lugar de "*")
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Config de swagger
	config.ConfigSwagger(r)

	// Rutas
	handlers.RegisterRoutes(r)
	auth.AuthRoutes(r)
	posts.PostRoutes(r)

	// Iniciar el server --->  go run ./cmd/app/main.go
	r.Run(":8080")
}
