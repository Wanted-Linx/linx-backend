package http

import "github.com/Wanted-Linx/linx-backend/api/delivery/http/handler"

func initRouter(
	userHandler *handler.UserHandler,
	studentHandler *handler.StudentHandler,
	companyHandler *handler.CompanyHandler,
	clubHandler *handler.ClubHandler,
	clubMemberHandler *handler.ClubMemberHandler,
	projectHandler *handler.ProjectHandler,
	projectClubHandler *handler.ProjectClubHandler) {
	userRouter(userHandler)
	studentRouter(studentHandler)
	companyRouter(companyHandler)
	clubRouter(clubHandler)
	clubMemberRouter(clubMemberHandler)
	projectRouter(projectHandler)
	projectClubRouter(projectClubHandler)
}

func userRouter(userHandler *handler.UserHandler) {
	group := e.Group("/users")
	group.POST("/signup", userHandler.SignUp)
	group.POST("/login", userHandler.Login)
	group.GET("/me", userHandler.GetUserByID)
}

func studentRouter(studentHandler *handler.StudentHandler) {
	group := e.Group("/students")
	// ?owner={{student_id}}
	group.GET("", studentHandler.GetStudent)
	// 학생 프로필 업로드
	group.POST("/profile/images", studentHandler.UploadProfileImage)
	// 학생 프로필 수정
	group.PUT("/profile", studentHandler.UpdateProfile)
	group.GET("/profile/images", studentHandler.GetProfileImage)
}

func companyRouter(companyHandler *handler.CompanyHandler) {
	group := e.Group("/companies")
	// ?owner={{company_id}}
	group.GET("/:company_id", companyHandler.GetCompany)
	group.GET("", companyHandler.GetAllCompanies)
	// 기업 프로필 이미지 업로드
	group.POST("/profile/images", companyHandler.UploadProfileImage)
	// 기업 프로필 수정
	group.PUT("/profile", companyHandler.UpdateProfile)
	group.GET("/profile/images", companyHandler.GetProfileImage)
}

func clubRouter(clubHandler *handler.ClubHandler) {
	group := e.Group("/clubs")
	group.POST("", clubHandler.CreateClub)
	group.GET("", clubHandler.GetAllClubs)
	group.GET("/:club_id", clubHandler.GetClubByID)
	group.POST("/profile/images", clubHandler.UploadProfileImage)
	group.GET("/profile/images", clubHandler.GetProfileImage)
}

func clubMemberRouter(clubMemberHandler *handler.ClubMemberHandler) {
	// 동아리에 가입 API
	group := e.Group("/clubs/members")
	group.POST("", clubMemberHandler.RegisterClubMember)
}

func projectRouter(projectHandler *handler.ProjectHandler) {
	group := e.Group("/projects")
	group.POST("", projectHandler.CreateProject)
	group.GET("", projectHandler.GetAllProjects)
	group.GET("/:project_id", projectHandler.GetProjectByID)
	group.POST("/logs", projectHandler.CreateProjectLog)
	group.POST("/feedbacks", projectHandler.CreateProjectLogFeedback)
	group.POST("/profile/images", projectHandler.UploadProfileImage)
	group.GET("/profile/images", projectHandler.GetProfileImage)
}

// TODO: 그냥 projectRouter에 합쳐도 됨
func projectClubRouter(projectClubHandler *handler.ProjectClubHandler) {
	group := e.Group("/projects/clubs")
	group.POST("", projectClubHandler.ApplyProjectClub)
}
