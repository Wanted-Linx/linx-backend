package http

import "github.com/Wanted-Linx/linx-backend/api/delivery/http/handler"

func initRouter(userHandler *handler.UserHandler) {
	userRouter(userHandler)
}

func userRouter(userHandler *handler.UserHandler) {
	group := e.Group("/users")
	group.POST("/signup", userHandler.SignUp)
	group.POST("/login", userHandler.Login)
	group.GET("/me", userHandler.GetUserByID)
}
