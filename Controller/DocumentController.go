package Controller

import (
	"fmt"
	"github.com/labstack/echo"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"restapi/Core"
	"restapi/Helper"
	"restapi/Model"
	"strings"
)

type DocumentController struct{}

func (h *DocumentController)Upload(c echo.Context) error {
	var tokenString = ""
	tokenString = c.Request().Header.Get("Authorization")
	if tokenString != ""{
		tokenString = strings.Split(tokenString, "Bearer ")[1]
	}

	userData := Authorize(tokenString)
	if userData == nil {
		return Helper.ResponseError(c,http.StatusUnauthorized,"unauthorized","Unauthorized")
	}

	file, err := c.FormFile("file")
	if err != nil {
		return Helper.ResponseError(c,http.StatusBadRequest,"bad request","Please insert file")
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	extension := filepath.Ext(file.Filename)
	filename := Helper.String(20)+extension


	var existingUser Model.User
	query := Core.DB.
		Where("code = ?", userData.Code).
		First(&existingUser)

	if query.RecordNotFound(){
		return Helper.ResponseError(c,http.StatusUnauthorized,"unauthorized","Unauthorized")
	}

	document := Model.Document{
		Code:Helper.String(12),
		FileName:filename,
		UserId:existingUser.ID,
	}

	exc := Core.DB.Create(&document)
	if exc.Error != nil {
		return Helper.ResponseError(c,http.StatusBadRequest,"bad request",fmt.Sprintf(exc.Error.Error()))
	}

	dst, err := os.Create("./Document/"+filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	data := map[string]interface{}{
		"file_name": filename,
	}

	return Helper.ResponseSuccess(c,http.StatusCreated,"success",data)
}

