package http

import "github.com/Wanted-Linx/linx-backend/api/delivery/http/handler"

func initRouter(
	userHandler *handler.UserHandler,
	studentHandler *handler.StudentHandler,
	companyHandler *handler.CompanyHandler,
	clubHandler *handler.ClubHandler,
	clubMemberHandler *handler.ClubMemberHandler) {
	userRouter(userHandler)
	studentRouter(studentHandler)
	companyRouter(companyHandler)
	clubRouter(clubHandler)
	clubMemberRouter(clubMemberHandler)
}

func userRouter(userHandler *handler.UserHandler) {
	group := e.Group("/users")
	group.POST("/signup", userHandler.SignUp)
	group.POST("/login", userHandler.Login)
	group.GET("/me", userHandler.GetUserByID)
}

func studentRouter(studentHandler *handler.StudentHandler) {
	group := e.Group("/students")
	group.GET("", studentHandler.GetStudent)
	// 학생 프로필 업로드
	group.POST("/profile/images", studentHandler.UploadProfileImage)
	// 학생 프로필 수정
	group.PUT("/profile", studentHandler.UpdateProfile)
	group.GET("/profile/images", studentHandler.GetProfileImage)
}

func companyRouter(companyHandler *handler.CompanyHandler) {
	_ = e.Group("/companies")
}

func clubRouter(clubHandler *handler.ClubHandler) {
	group := e.Group("/clubs")
	group.POST("", clubHandler.CreateClub)
	group.GET("", clubHandler.GetAllClubs)
	group.GET("/:club_id", clubHandler.GetClubByID)
}

func clubMemberRouter(clubMemberHandler *handler.ClubMemberHandler) {
	// 동아리에 가입 API
	group := e.Group("/clubs/members")
	group.POST("", clubMemberHandler.RegisterClubMember)
}
