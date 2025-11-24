package auth

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	//Password  string `json:"password"`
	AvatarURL string `json:"avatar_url"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
