package users

// input
type RegisterUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Passion  string `json:"passion" binding:"required"`
	Password string `json:"password" binding:"required"`
}
