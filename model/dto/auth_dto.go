package dto

type LoginDTO struct {
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

type RegisterDTO struct {
	Name     string `binding:"required" json:"name"`
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}
