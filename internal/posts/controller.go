package posts

import (
	"net/http"

	"go000/internal/config"
	"go000/internal/models"

	"github.com/gin-gonic/gin"
)

// Listar todos los posts

// @Summary Listar todos los posts
// @Description Retorna todos los posts publicados por todos los usuarios
// @Tags Posts
// @Produce json
// @Success 200 {object} map[string]interface{} "posts"
// @Router /posts/ [get]
func AllPosts(c *gin.Context) {
	var posts []models.Post

	// Salida de posts con user completo
	config.DB.Preload("User").Order("created_at desc").Find(&posts)

	resp := make([]PostResponse, len(posts))

	for i, p := range posts {
		resp[i] = PostResponse{
			ID:      p.ID,
			Content: p.Content,
			UserID:  p.UserID,
		}
	}

	c.JSON(200, gin.H{"posts": resp})

	//c.JSON(http.StatusOK, gin.H{"posts": posts})
}

// Crear post por cada user
// @Summary Crear Post
// @Description Crear nuevo post de usuario logeado
// @Tags posts
// @Accept json
// @Produce json
// @Param post body posts.CreatePostInput true "Contenido del post"
// @Security BearerAuth
// @Sucess 200 {object} posts.PostResponse
// @Failure 400 {object} map[string]string
// @Router /posts/ [post]
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

	response := PostResponse{
		ID:      post.ID,
		Content: post.Content,
		UserID:  post.UserID,
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Post creado",
		"post":    response,
	})
}

func GetPostsUser(c *gin.Context) {

	userID := c.Param("id")

	var posts []models.Post
	config.DB.Where("user_id=?", userID).Preload("User").Order("created_at desc").Find(&posts)

	c.JSON(http.StatusOK, gin.H{"posts": posts})

}
