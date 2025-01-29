package request

type UserRequest struct {
	Name     string `json:"name" binding:"required, min = 3, max=255" `
	Email    string `json:"email" binding:"required, email" `
	Password string `json:"password" binding:"required, min = 6, max=21, containsAny=!@#$%&*"`
	Area     string `json:"area" binding:"required"`
}
