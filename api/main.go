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
	clubRepo := repository.NewClubRepository(dbClient)
	clubMemberRepo := repository.NewClubMemberRepository(dbClient)
	projectRepo := repository.NewProjectRepository(dbClient)
	projectClubRepo := repository.NewProjectClubRepository(dbClient)

	userService := service.NewUserSerivce(userRepo, studentRepo)
	studentService := service.NewStudentService(studentRepo, clubMemberRepo)
	companyService := service.NewCompanyService(companyRepo)
	clubService := service.NewClubService(clubRepo, clubMemberRepo)
	clubMemberService := service.NewClubMemberService(clubMemberRepo)
	projectService := service.NewProjectService(projectRepo, projectClubRepo)
	projectClubService := service.NewProjectClubService(projectClubRepo)

	userHandler := handler.NewUserHandler(userService, studentService, companyService)
	studentHandler := handler.NewStudentHandler(studentService)
	companyHandler := handler.NewCompanyHandler(companyService)
	clubHandler := handler.NewClubHandler(clubService)
	clubMemberHandler := handler.NewClubMemberHandler(clubMemberService)
	projectHandler := handler.NewProjectHandler(projectService)
	projectClubHandler := handler.NewProjectClubHandler(projectClubService)

	server := http.NewServer(userHandler, studentHandler, companyHandler,
		clubHandler, clubMemberHandler, projectHandler, projectClubHandler)

	server.Logger.Fatal(server.Start(fmt.Sprintf(":%d", config.Config.Port)))
}
