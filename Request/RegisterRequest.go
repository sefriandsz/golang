package Request

type RegisterRequest struct {
	Name                         string `json:"name" form:"username" valid:"required"`
	Username                     string `json:"username" form:"username" valid:"required"`
	Email                        string `json:"email" form:"email" valid:"required,email"`
	Password                     string `json:"password" form:"password" valid:"required"`
	ConfirmPassword              string `json:"confirm_password" form:"confirm_password" valid:"required"`
}
