package main

import (
	"fmt"

	"github.com/Wanted-Linx/linx-backend/api/config"
	"github.com/Wanted-Linx/linx-backend/api/delivery/http"
	"github.com/Wanted-Linx/linx-backend/api/delivery/http/handler"
	"github.com/Wanted-Linx/linx-backend/api/repository"
	"github.com/Wanted-Linx/linx-backend/api/service"
)

func main() {
	dbClient := repository.Connect()

	userRepo := repository.NewUserRepository(dbClient)
	studentRepo := repository.NewStudentRepository(dbClient)
	companyRepo := repository.NewCompanyRepository(dbClient)

	userService := service.NewUserSerivce(userRepo, studentRepo)
	studentService := service.NewStudentService(studentRepo)
	companyService := service.NewCompanyService(companyRepo)

	userHandler := handler.NewUserHandler(userService, studentService, companyService)
	studentHandler := handler.NewStudentHandler(studentService)
	companyHandler := handler.NewCompanyHandler(companyService)

	server := http.NewServer(userHandler, studentHandler, companyHandler)

	server.Logger.Fatal(server.Start(fmt.Sprintf(":%d", config.Config.Port)))
}
