package posts

type CreatePostInput struct {
	Content string `json:"content" binding:"required"`
}
