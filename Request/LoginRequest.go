package Request

type LoginRequest struct {
	Email  string `json:"email" form:"email" valid:"required,email"`
	Password string `json:"password" form:"password" valid:"required"`
}
