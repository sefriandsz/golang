package Helper

import (
	"github.com/labstack/echo"
	"net/http"
)

func ResponseError(c echo.Context,statusCode uint,status string, message string) (error) {
	return c.JSON(http.StatusBadRequest,map[string]interface{}{
		"status" : status,
		"status_code" : statusCode,
		"message" : message,
	})
}

func ResponseSuccess(c echo.Context,statusCode uint,status string, data map[string]interface{}) (error) {
	return c.JSON(http.StatusBadRequest,map[string]interface{}{
		"status" : status,
		"status_code" : statusCode,
		"data" : data,
	})
}