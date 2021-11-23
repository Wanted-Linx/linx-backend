package http

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/wanted-linx/linx-backend/api/ent"
	restErr "github.com/wanted-linx/linx-backend/api/errors"
)

func httpErrorHandler(err error, c echo.Context) {
	var (
		code      = http.StatusInternalServerError
		errorType = "internal_error"
		msg       string
	)

	log.Errorf("%+v", err)
	if httpError, ok := err.(*restErr.RestErr); ok {
		code = httpError.Code
		errorType = httpError.ErrorType
		msg = httpError.Message

	} else if ent.IsNotFound(err) {
		code = 404
		errorType = "not_found_error"
		msg = "리소스를 찾을 수 없습니다."
	} else if errors.As(err, &restErr.BadRequestErr) {
		log.Println("dd")
		code = 400
		errorType = "bad_request_error"
		msg = err.Error()
	} else if e.Debug {
		msg = err.Error()
	} else {
		msg = http.StatusText(code)
	}

	if !c.Response().Committed {
		if c.Request().Method == echo.HEAD {
			err := c.NoContent(code)
			if err != nil {
				c.Logger().Error(err)
			}
		} else {
			err := c.JSON(code, restErr.NewRestError(code, msg, errorType))
			if err != nil {
				c.Logger().Error(err)
			}
		}
	}
}
