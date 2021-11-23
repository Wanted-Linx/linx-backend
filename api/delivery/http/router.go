package http

import (
	"github.com/wanted-linx/linx-backend/api/delivery/http/handler"
)

func initRouter(userHandler *handler.UserHandler) {
	userRouter(userHandler)
}

func userRouter(userHandler *handler.UserHandler) {
	group := e.Group("/users")
	group.POST("/signup", userHandler.SignUp)
	group.GET("", userHandler.GetUserByID)
}
