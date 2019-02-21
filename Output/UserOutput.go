package Output

import (
	"restapi/Model"
)

type UserOutput struct {
	Code                         string `json:"code" form:"code" query:"code"`
	Name                         string `json:"name" form:"username" query:"name"`
	Username                     string `json:"username" form:"username" query:"username"`
	Email                        string `json:"email" form:"email" query:"email"`
	Status                     string `json:"status" form:"status" query:"status"`
}

func NewUserOutput(user *Model.User) UserOutput{
	return UserOutput{
		Code:user.Code,
		Name:user.Name,
		Username:user.Username,
		Email:user.Email,
		Status:user.Status,
	}
}