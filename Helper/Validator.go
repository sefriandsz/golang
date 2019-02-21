package Helper

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"net/http"
	"github.com/labstack/echo"
)

func Validator(c echo.Context,req interface{}) (error){
	_, err := govalidator.ValidateStruct(req)
	if err != nil {
		return ResponseError(c,http.StatusBadRequest,"bad request",fmt.Sprintf(err.Error()))
	}

	return nil
}
