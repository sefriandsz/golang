package Output

import "restapi/Model"

type AuthorizeOutput struct {
	Code                         string `json:"code" form:"code" query:"code"`
	Email                        string `json:"email" form:"email" query:"email"`
}

func NewAuthorizeOutput(user *Model.User) AuthorizeOutput{
	return AuthorizeOutput{
		Code:user.Code,
		Email:user.Email,
	}
}