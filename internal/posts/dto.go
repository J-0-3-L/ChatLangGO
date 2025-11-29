package posts

import "time"

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

type UserPostsByID struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
}

type UserPostResponse struct {
	ID        uint          `json:"id"`
	Content   string        `json:"content"`
	CreatedAt time.Time     `json:"created_at"`
	User      UserPostsByID `json:"user"`
}

type UpdatePostInput struct {
	Content string `json:"content" binding:"required"`
}
