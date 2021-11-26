package domain

import (
	"log"
	"time"

	"github.com/Wanted-Linx/linx-backend/api/ent"
)

type ProjectDto struct {
	ID                int              `json:"id"`
	Name              string           `json:"name"`
	Content           string           `json:"content"`
	StartDate         string           `json:"start_date"`
	EndDate           string           `json:"end_date"`
	ApplyingStartDate string           `json:"applying_start_date"`
	AppylingEndDate   string           `json:"appyling_end_date"`
	Qualification     string           `json:"qualification"`
	SponsorFee        int              `json:"sponsor_fee"`
	TaskType          []string         `json:"task_type"`
	Company           *CompanyDto      `json:"company"`     // 프로젝트 만든 기업
	ProjectClubs      []*ProjectClub   `json:"clubs"`       // 프로젝트 참가 신청한 동아리 목록(따로 조회?)
	ProjectLogs       []*ProjectLogDto `json:"project_log"` // 이 프로젝트에 참가한 모든 동아리의 로그 목록들
	CreatedAt         time.Time        `json:"created_at"`
}

type ProjectLogDto struct {
	ID           int                         `json:"id"`
	Title        string                      `json:"title"`
	Author       string                      `json:"author"`
	Content      string                      `json:"content"`
	StartDate    string                      `json:"start_date"`
	EndDate      string                      `json:"end_date"`
	Participants []*ProjectLogParticipantDto `json:"participants"`
	Feedbacks    []*ProjectLogFeedbackDto    `json:"feedbacks"`
	// ProjectID string `json:"project_id"`
}

type ProjectLogParticipantDto struct {
	ID   int    `json:"id"`
	Name string `json:"name"` // 나중에 studentDto로 연결할까...(그럼 table에도 student_id가 FK로 있어야 함)
}

type ProjectLogFeedbackDto struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`  // 제목은 필요없을지도...
	Author  string `json:"author"` // 나중에 companyDto 연결할까...(그럼 table에도 company_id가 FK로 있어야 함)
	Content string `json:"content"`
}

type ProjectCreateRequest struct {
	Name              string   `json:"name"`
	Content           string   `json:"content"`
	StartDate         string   `json:"start_date"`
	EndDate           string   `json:"end_date"`
	ApplyingStartDate string   `json:"applying_start_date"`
	AppylingEndDate   string   `json:"applying_end_date"`
	Qualification     string   `json:"qualification"`
	SponsorFee        int      `json:"sponsor_fee"`
	TaskType          []string `json:"task_type"`
}

type ProjectLogCreateRequest struct {
	Title        string   `json:"title"`
	Author       string   `json:"author"`
	Content      string   `json:"content"`
	StartDate    string   `json:"start_date"`
	EndDate      string   `json:"end_date"`
	Participants []string `json:"participants"`
	ClubID       int      `json:"club_id"`
	ProjectID    int      `json:"project_id"`
}

type ProjectLogFeedbackRequest struct {
	ProjectLogID int    `json:"project_log_id"`
	Author       string `json:"author"` // company name
	Content      string `json:"content"`
}

type ProjectService interface {
	CreateProject(companyID int, reqProject *ProjectCreateRequest) (*ProjectDto, error)
	CreateProjectLog(reqPlog *ProjectLogCreateRequest) (*ProjectLogDto, error)
	CreateProjectLogFeedback(companyID int, reqPlfeedback *ProjectLogFeedbackRequest) (*ProjectLogFeedbackDto, error)
	GetProjectByID(projectID int) (*ProjectDto, error)
	GetAllProjects(limit, offset int) ([]*ProjectDto, error)
}

type ProjectRepository interface {
	Save(reqProject *ent.Project) (*ent.Project, error)
	SaveProjectLog(reqPlog *ent.ProjectLog) (*ent.ProjectLog, error)
	SaveProjectLogFeedback(reqPfeedback *ent.ProjectLogFeedback) (*ent.ProjectLogFeedback, error)
	GetByID(projectID int) (*ent.Project, []*ent.Club, error)
	GetAll(limit, offset int) ([]*ent.Project, [][]*ent.Club, error)
}

func ProjectToDto(src *ent.Project, srcProjectClub []*ent.Club) *ProjectDto {
	log.Println(src.Edges.ProjectLog)
	return &ProjectDto{
		ID:                src.ID,
		Name:              src.Name,
		Content:           src.Content,
		StartDate:         src.StartDate,
		EndDate:           src.EndDate,
		ApplyingStartDate: src.ApplyingStartDate,
		AppylingEndDate:   src.ApplyingEndDate,
		Qualification:     src.Qualification,
		SponsorFee:        src.SponsorFee,
		TaskType:          []string{"개발, 디자인"},
		Company:           CompanyToDto(src.Edges.Company),
		ProjectLogs:       ProjectLogsToDto(src.Edges.ProjectLog),
		ProjectClubs:      ProjectClubsToDto(srcProjectClub),
		CreatedAt:         src.CreatedAt,
	}
}

func ProjectLogToDto(src *ent.ProjectLog) *ProjectLogDto {
	return &ProjectLogDto{
		ID:           src.ID,
		Title:        src.Title,
		Author:       src.Author,
		Content:      src.Content,
		StartDate:    src.StartDate,
		EndDate:      src.EndDate,
		Participants: ProjectLogParticipantsToDto(src.Edges.ProjectLogParticipant),
		Feedbacks:    ProjectLogFeedbacksToDto(src.Edges.ProjectLogFeedback),
	}
}

func ProjectLogsToDto(src []*ent.ProjectLog) []*ProjectLogDto {
	projectLogsDto := make([]*ProjectLogDto, 0)

	for _, projectLog := range src {
		projectLogDto := ProjectLogToDto(projectLog)
		projectLogsDto = append(projectLogsDto, projectLogDto)
	}

	return projectLogsDto
}

func ProjectLogParticipantToDto(src *ent.ProjectLogParticipant) *ProjectLogParticipantDto {
	return &ProjectLogParticipantDto{
		ID:   src.ID,
		Name: src.Name,
	}
}

func ProjectLogParticipantsToDto(src []*ent.ProjectLogParticipant) []*ProjectLogParticipantDto {
	participantsDto := make([]*ProjectLogParticipantDto, 0)

	for _, participant := range src {
		participantDto := ProjectLogParticipantToDto(participant)
		participantsDto = append(participantsDto, participantDto)
	}

	return participantsDto
}

func ProjectLogFeedbackToDto(src *ent.ProjectLogFeedback) *ProjectLogFeedbackDto {
	return &ProjectLogFeedbackDto{
		ID:      src.ID,
		Author:  src.Author,
		Content: src.Content,
	}
}

func ProjectLogFeedbacksToDto(src []*ent.ProjectLogFeedback) []*ProjectLogFeedbackDto {
	feedbacksDto := make([]*ProjectLogFeedbackDto, 0)

	for _, feedback := range src {
		feedbackDto := ProjectLogFeedbackToDto(feedback)
		feedbacksDto = append(feedbacksDto, feedbackDto)
	}

	return feedbacksDto
}
