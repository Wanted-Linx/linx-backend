package service

import (
	"github.com/Wanted-Linx/linx-backend/api/domain"
	"github.com/Wanted-Linx/linx-backend/api/ent"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type projectService struct {
	projectRepo     domain.ProjectRepository
	projcetClubRepo domain.ProjectClubRepository
}

func NewProjectService(projectRepo domain.ProjectRepository, projcetClubRepo domain.ProjectClubRepository) domain.ProjectService {
	return &projectService{
		projectRepo:     projectRepo,
		projcetClubRepo: projcetClubRepo,
	}
}

func (s *projectService) CreateProject(companyID int, reqProject *domain.ProjectCreateRequest) (*domain.ProjectDto, error) {

	// TODO: profile image 먼저 upload 후 저장된 path 값을 ProfileImage에 집어넣는다.
	// 우선은 default로 '/profile/club/clubLeaderID/
	// profile_key := fmt.Sprintf("/companies/%d/project/profile/key")

	project := &ent.Project{
		Name:              reqProject.Name,
		Content:           reqProject.Content,
		StartDate:         reqProject.StartDate,
		EndDate:           reqProject.EndDate,
		ApplyingStartDate: reqProject.ApplyingStartDate,
		ApplyingEndDate:   reqProject.AppylingEndDate,
		Qualification:     reqProject.Qualification,
		Edges: ent.ProjectEdges{
			Company: &ent.Company{ID: companyID},
		},
		SponsorFee: reqProject.SponsorFee,
		// ProfileImage: reqProject.ProfileImage,
	}

	newProject, err := s.projectRepo.Save(project)
	if err != nil {
		return nil, errors.WithMessage(err, "알 수 없는 에러가 발생했습니다.")
	}

	// log.Println(newClub, clubLeaderID, newClub.Edges.Leader)

	log.Info("프로젝트 생성 완료", newProject)
	return domain.ProjectToDto(newProject, nil), nil
}

func (s *projectService) GetProjectByID(projectID int) (*domain.ProjectDto, error) {
	project, clubs, err := s.projectRepo.GetByID(projectID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.WithMessage(err, "해당하는 프로젝트를 찾지 못했습니다.")
		}
		return nil, errors.WithMessage(err, "알 수 없는 오류가 발생했습니다.")
	}

	log.Info("프로젝트 조회 완료", project, clubs)
	return domain.ProjectToDto(project, clubs), nil
}

func (s *projectService) GetAllProjects(limit, offset int) ([]*domain.ProjectDto, error) {
	projects, allClubs, err := s.projectRepo.GetAll(limit, offset)
	if err != nil {
		return nil, errors.WithMessage(err, "알 수 없는 오류가 발생했습니다.")
	}

	projectsDto := make([]*domain.ProjectDto, len(projects))

	for idx, project := range projects {
		projectDto := domain.ProjectToDto(project, allClubs[idx])
		projectsDto[idx] = projectDto
	}

	return projectsDto, nil
}

func (s *projectService) CreateProjectLog(reqPlog *domain.ProjectLogCreateRequest) (*domain.ProjectLogDto, error) {
	projectClub, err := s.projcetClubRepo.Get(reqPlog.ProjectID, reqPlog.ClubID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.WithMessage(err, "리소스를 찾지 못했습니다.")
		}
		return nil, errors.WithMessage(err, "알 수 없는 오류가 발생했습니다.")
	}

	pl := &ent.ProjectLog{
		Title:     reqPlog.Title,
		Author:    reqPlog.Author,
		Content:   reqPlog.Content,
		StartDate: reqPlog.StartDate,
		EndDate:   reqPlog.EndDate,
		Edges: ent.ProjectLogEdges{
			Project: &ent.Project{ID: reqPlog.ProjectID},
			ProjectClub: &ent.ProjectClub{
				ID:        projectClub.ID,
				ProjectID: reqPlog.ProjectID,
				ClubID:    reqPlog.ClubID,
			},
		},
	}

	projectLog, err := s.projectRepo.SaveProjectLog(pl)
	if err != nil {
		return nil, errors.WithMessage(err, "알 수 없는 오류가 발생했습니다.")
	}

	return domain.ProjectLogToDto(projectLog), nil
}

func (s *projectService) CreateProjectLogFeedback(companyID int, reqPlfeedback *domain.ProjectLogFeedbackRequest) (*domain.ProjectLogFeedbackDto, error) {
	plf := &ent.ProjectLogFeedback{
		// Title:   reqPlfeedback.Title,
		Author:  reqPlfeedback.Author,
		Content: reqPlfeedback.Content,
		Edges: ent.ProjectLogFeedbackEdges{
			ProjectLog: &ent.ProjectLog{ID: reqPlfeedback.ProjectLogID},
		},
	}

	projectLogFeedback, err := s.projectRepo.SaveProjectLogFeedback(plf)
	if err != nil {
		return nil, errors.WithMessage(err, "알 수 없는 오류가 발생했습니다.")
	}

	log.Info("피드백 작성 완료", projectLogFeedback)
	return domain.ProjectLogFeedbackToDto(projectLogFeedback), nil
}
