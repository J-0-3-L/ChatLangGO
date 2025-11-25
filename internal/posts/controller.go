package posts

import (
	"net/http"

	"go000/internal/config"
	"go000/internal/models"

	"github.com/gin-gonic/gin"
)

// Listar todos los posts
func AllPosts(c *gin.Context) {
	var posts []models.Post

	config.DB.Preload("User").Order("created_at desc").Find(&posts)
	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

// Crear post por cada user
func CreatePost(c *gin.Context) {

	var input CreatePostInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El contenido es requerido"})
		return
	}

	userID := c.GetUint("user_id")

	post := models.Post{
		Content: input.Content,
		UserID:  userID,
	}

	if err := config.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear el post"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Post creado",
		"post":    post,
	})
}
