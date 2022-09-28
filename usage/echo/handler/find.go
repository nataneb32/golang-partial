package handler

import (
	"partial/usage"
	"strconv"

	"github.com/labstack/echo/v4"
)


func FindUserHandler(e echo.Context) error {
	var uid uint64
	uid, err := strconv.ParseUint(e.Param("id"), 10, 64)

	user := usage.FindUser(uint(uid))
	if err != nil {
		return err
	}

	e.JSON(200, user)
	return nil
}
