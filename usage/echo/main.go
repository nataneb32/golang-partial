package main

import (
	"partial/usage/echo/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	
	e.PATCH("/user", handler.UpdateUserHandler)
	e.POST("/user", handler.SaveUserHandler)
	e.GET("/user/:id", handler.FindUserHandler)
	
	e.Start(":8888")
}
