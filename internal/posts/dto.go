package posts

// import "go000/internal/models"

// type AllPostsResponse struct {
// 	Posts []models.Post `json:"posts"`
// }

type CreatePostInput struct {
	Content string `json:"content" binding:"required"`
}
