package routes

import (
	usercontroller "prakerja11/controllers/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoutes(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.GET("/users", usercontroller.GetUsersController)
}