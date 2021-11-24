package util

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetRequestUserID(c echo.Context) (int, error) {
	header := c.Request().Header

	return strconv.Atoi(header["User"][0])
}