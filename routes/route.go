package routes

import (
	authcontroller "prakerja11/controllers/auth"
	usercontroller "prakerja11/controllers/user"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoutes(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.POST("/login", authcontroller.LoginController)
	e.POST("/register", authcontroller.RegisterController)

	eAuth := e.Group("")
	eAuth.Use(echojwt.JWT([]byte("123")))
	eAuth.GET("/users", usercontroller.GetUsersController)

}