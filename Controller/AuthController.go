package Controller

import (
	"fitco/utils/helper/hash"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"restapi/Core"
	"restapi/Helper"
	"restapi/Model"
	"restapi/Output"
	"restapi/Request"
	"time"
)

type AuthController struct{}
func (h *AuthController) Login(c echo.Context) error {
	req := new(Request.LoginRequest)
	if errs := c.Bind(&req); errs != nil {
		return errs
	}

	_, err := govalidator.ValidateStruct(req)
	if err != nil {
		return Helper.ResponseError(c,http.StatusBadRequest,"bad request",fmt.Sprintf(err.Error()))
	}

	var existingUser Model.User
	query := Core.DB.
		Where("email = ?", req.Email).
		First(&existingUser)

	if query.RecordNotFound(){
		return Helper.ResponseError(c,http.StatusBadRequest,"bad request","User Not Found")
	}

	if(existingUser.Email == req.Email && hash.Check(req.Password,existingUser.Password)){
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["email"] = existingUser.Email
		claims["code"] = existingUser.Code
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		jwtToken, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}

		data := map[string]interface{}{
			"user" : map[string]interface{}{
				"email" : existingUser.Email,
				"code" : existingUser.Code,
			},
			"token": jwtToken,
		}

		return Helper.ResponseSuccess(c,http.StatusOK,"success",data)
	}

	return Helper.ResponseError(c,http.StatusUnauthorized,"unauthorized","Unauthorized")
}

func (h *AuthController) Register(c echo.Context) error {
	req := new(Request.RegisterRequest)
	if errs := c.Bind(&req); errs != nil {
		return errs
	}

	_, err := govalidator.ValidateStruct(req)
	if err != nil {
		return Helper.ResponseError(c,http.StatusBadRequest,"bad request",fmt.Sprintf(err.Error()))
	}

	if req.Password != req.ConfirmPassword {
		return Helper.ResponseError(c,http.StatusBadRequest,"bad request","Password is not match")
	}

	var existingUser Model.User
	query := Core.DB.
		Where("email = ?", req.Email).
		Or("username = ?", req.Username).
		First(&existingUser)

	if !query.RecordNotFound(){
		return Helper.ResponseError(c,http.StatusBadRequest,"bad request","Duplicate username or email")
	}

	user := Model.User{
		Name:req.Name,
		Username:req.Username,
		Email:req.Email,
		Password:hash.Make(req.Password),
		Status:"Active",
		Code:Helper.String(12),
	}

	var userOutput = Output.NewUserOutput(&user)

	exc := Core.DB.Create(&user)
	if exc.Error != nil {
		return Helper.ResponseError(c,http.StatusBadRequest,"bad request",fmt.Sprintf(exc.Error.Error()))
	}
	data := map[string]interface{}{
		"user": userOutput,
	}
	return Helper.ResponseSuccess(c,http.StatusCreated,"success",data)
}

func Authorize(tokenString string)(*Output.AuthorizeOutput){
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})

	if token != nil && err == nil {
		var outputAuth Output.AuthorizeOutput
		var userLogin = mapstructure.Decode(token.Claims, &outputAuth)
		fmt.Println(userLogin)

		return &outputAuth

	}

	return nil
}
