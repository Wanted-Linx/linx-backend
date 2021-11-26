package domain

import (
	"github.com/Wanted-Linx/linx-backend/api/ent"
)

type ProjectClubApplyRequest struct {
	ProjectID int    `json:"project_id"`
	ClubID    int    `json:"club_Id"`
	StartDate string `json:"start_date"`
	//ApplyingDate string `json:"applying_date"` 지원한 날짜가 모집기간보다 지났을 경우(그냥 time.Now()로 비교해도 될 듯)
}

type ProjectClub struct {
	ClubID     int    `json:"club_id"`
	ClubName   string `json:"club_name"`
	LeaderID   int    `json:"leader_id"`
	LeaderName string `json:"leader_name"`
}

type ClubProject struct {
	ProjectID   int              `json:"project_id"`
	ProjectName string           `json:"project_name"`
	CompanyName string           `json:"company_name"`
	ProjectLogs []*ProjectLogDto `json:"project_logs"`
}

type ProjectClubService interface {
	ApplyProjectClub(reqPc *ProjectClubApplyRequest) error
	// GetClubMemberByClubID(clubID int) (ClubMembersDto, error)
	// GetAllClubMembers(studentID, clubID int) (ClubMembersDto, error)
}

type ProjectClubRepository interface {
	Apply(clubID, projectID int, startDate string) (*ent.ProjectClub, error)
	// FindbyClubID(clubID int) ([]*ent.ClubMember, error)
	// FindByStudentIDAndClubID(studentID, clubID int) (*ent.ClubMember, error)
}

func ProjectClubsToDto(srcProjcetClubs []*ent.Club) []*ProjectClub {
	clubs := make([]*ProjectClub, 0)

	for _, projectClub := range srcProjcetClubs {
		var club ProjectClub
		club.ClubID = projectClub.ID
		club.ClubName = projectClub.Name
		club.LeaderID = projectClub.Edges.Leader.ID
		club.LeaderName = projectClub.Edges.Leader.Name

		clubs = append(clubs, &club)
	}

	return clubs
}

func ClubProjectsToDto(srcClubProject []*ent.Project) []*ClubProject {
	projects := make([]*ClubProject, 0)

	for _, clubProject := range srcClubProject {
		var project *ClubProject
		project.ProjectID = clubProject.ID
		project.ProjectName = clubProject.Name
		project.CompanyName = clubProject.Edges.Company.Name
		project.ProjectLogs = ProjectLogsToDto(clubProject.Edges.ProjectLog)

		projects = append(projects, project)
	}

	return projects
}
