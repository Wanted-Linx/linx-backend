package main

import (
	"fmt"

	"github.com/wanted-linx/linx-backend/api/config"
	"github.com/wanted-linx/linx-backend/api/delivery/http"
	"github.com/wanted-linx/linx-backend/api/delivery/http/handler"
	"github.com/wanted-linx/linx-backend/api/repository"
	"github.com/wanted-linx/linx-backend/api/service"
)

func main() {
	dbClient := repository.Connect()

	userRepo := repository.NewUserRepository(dbClient)
	userService := service.NewUserSerivce(userRepo)
	userHandler := handler.NewUserHandler(userService)

	server := http.NewServer(userHandler)

	server.Logger.Fatal(server.Start(fmt.Sprintf(":%d", config.Config.Port)))
}
