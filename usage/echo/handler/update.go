package handler

import (
	"partial"
	"partial/usage"
	"partial/usage/entity"

	"github.com/labstack/echo/v4"
)

func UpdateUserHandler(e echo.Context) error {
	var puser partial.PartialJSON[entity.User]
	err := e.Bind(&puser)
	if err != nil {
		return err
	}

	err = usage.UpdateUser(puser)
	if err != nil {
		return err
	}

	return nil
}
