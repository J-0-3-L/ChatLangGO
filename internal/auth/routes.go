package auth

import "github.com/gin-gonic/gin"

func AuthRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", Register)
		auth.POST("/login", Login)
	}
}
