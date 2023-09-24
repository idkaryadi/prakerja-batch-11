package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id int 		`json:"id"`
	Name string `json:"name"`
}

type BaseResponse struct {
	Status bool 		`json:"status"`
	Message string 		`json:"message"`
	Data interface{} 	`json:"data"`
}

func main(){
	e := echo.New()
	e.GET("/users", GetUsersController)
	e.Start(":8000")
}

func GetUsersController(c echo.Context) error{
	var users = []User{
		{1, "Alta"},
		{2, "Alterra"},
		{3, "Academy"},
	}
	return c.JSON(http.StatusOK, BaseResponse{
		Status: true,
		Message: "Success",
		Data: users,
	})
}
