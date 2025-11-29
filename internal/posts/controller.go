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
// @Tags Posts
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

// Listar posts de un user
// @Summary      Listar todos los post de un user
// @Description  Retorna todos los posts creados por un user específico
// @Tags         Posts
// @Param        id   path      int  true  "ID del usuario"
// @Produce      json
// @Success      200  {object}  map[string][]UserPostResponse
// @Failure      400  {object}  map[string]string "error"
// @Router       /posts/{id} [get]
func GetPostsUser(c *gin.Context) {

	userID := c.Param("id")

	var posts []models.Post
	config.DB.Where("user_id=?", userID).Preload("User").Order("created_at desc").Find(&posts)

	resp := make([]UserPostResponse, len(posts))

	for i, p := range posts {
		resp[i] = UserPostResponse{
			ID:        p.ID,
			Content:   p.Content,
			CreatedAt: p.CreatedAt,
			User: UserPostsByID{
				ID:        p.User.ID,
				Username:  p.User.Username,
				AvatarURL: p.User.AvatarURL,
			},
		}
	}
	c.JSON(http.StatusOK, gin.H{"posts": resp})

}

//Actualizar post de un user

// @Summary        Actualizar un post
// @Description    El usuario dueño del post puede actualizar su contenido
// @Tags           Posts
// @Accept         json
// @Produce        json
// @Param          id path int true "ID del post"
// @Param          body body UpdatePostInput true "Contenido nuevo"
// @Security       BearerAuth
// @Success        200 {object} map[string]interface{}
// @Failure        400 {object} map[string]interface{}
// @Failure        401 {object} map[string]interface{}
// @Failure        403 {object} map[string]interface{}
// @Failure        404 {object} map[string]interface{}
// @Router         /posts/{id} [put]
func UpdatePost(c *gin.Context) {

	id := c.Param("id")
	//userID:=c.GetUint("user_id")

	var post models.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post no encontrado"})
		return
	}

	var body UpdatePostInput
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Contenido Invalido"})
		return
	}

	post.Content = body.Content

	config.DB.Save(&post)

	c.JSON(http.StatusOK, gin.H{"message": "Post Actualizado", "post": post})

}
