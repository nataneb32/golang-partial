package handler

import (
	"partial/usage"
	"partial/usage/entity"

	"github.com/labstack/echo/v4"
)


func SaveUserHandler(e echo.Context) error {
	var user entity.User
	err := e.Bind(&user)
	if err != nil {
		return err
	}

	usage.SaveUser(&user)
	if err != nil {
		return err
	}

	return nil
}
