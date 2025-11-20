package config

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	//"github.com/swaggo/gin-swagger"
	//"github.com/swaggo/gin-swagger/swaggerFiles"
	//"github.com/swaggo/gin-swagger"
)

// @title Mi API de Ejemplo
// @version 1.0
// @description Esta es la documentación de la API usando Gin y Swagger.
func ConfigSwagger(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
