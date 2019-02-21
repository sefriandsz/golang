package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"restapi/Controller"
)
func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// CORS middleware setup
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE, echo.HEAD},
	}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	authController := &Controller.AuthController{}
	documentController := &Controller.DocumentController{}
	e.POST("/login", authController.Login)
	e.POST("/register", authController.Register)
	loginRequest := e.Group("/my")
	loginRequest.Use(middleware.JWT([]byte("secret")))
	loginRequest.POST("/upload", documentController.Upload)
	e.Logger.Fatal(e.Start(":8080"))
}