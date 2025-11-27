package posts

// import "go000/internal/models"

// type AllPostsResponse struct {
// 	Posts []models.Post `json:"posts"`
// }

type CreatePostInput struct {
	Content string `json:"content" binding:"required"`
}

type PostResponse struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
}
