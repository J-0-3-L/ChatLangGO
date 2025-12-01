package posts

import (
	"go000/internal/auth"

	"github.com/gin-gonic/gin"
)

func PostRoutes(r *gin.Engine) {

	posts := r.Group("/posts")

	posts.GET("/", AllPosts)
	posts.GET("/:id", GetPostsUser)

	posts.Use(auth.AuthRequired())
	{
		posts.POST("/", CreatePost)
		posts.PUT("/:id", UpdatePost)
		posts.DELETE("/:id", DeletePost)
	}

}
