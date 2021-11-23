package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/wanted-linx/linx-backend/api/delivery/http/handler"
)

var e *echo.Echo

func NewServer(userHandler *handler.UserHandler) *echo.Echo {
	e = echo.New()
	e.HTTPErrorHandler = httpErrorHandler
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	initRouter(userHandler)
	return e
}
