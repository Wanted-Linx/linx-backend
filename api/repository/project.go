package repository

import (
	"context"
	"log"
	"time"

	"github.com/Wanted-Linx/linx-backend/api/domain"
	"github.com/Wanted-Linx/linx-backend/api/ent"
	"github.com/Wanted-Linx/linx-backend/api/ent/project"
	"github.com/Wanted-Linx/linx-backend/api/ent/tasktype"
	"github.com/pkg/errors"
)

type projectRepository struct {
	db *ent.Client
}

func NewProjectRepository(db *ent.Client) domain.ProjectRepository {
	return &projectRepository{db: db}
}

func (r *projectRepository) Save(reqProject *ent.Project) (*ent.Project, error) {
	p, err := r.db.Project.Create().
		SetName(reqProject.Name).
		SetNillableTaskExperience(&reqProject.TaskExperience).
		SetContent(reqProject.Content).
		SetStartDate(reqProject.StartDate).
		SetEndDate(reqProject.EndDate).
		SetApplyingStartDate(reqProject.ApplyingStartDate).
		SetApplyingEndDate(reqProject.ApplyingEndDate).
		SetQualification(reqProject.Qualification).
		SetSponsorFee(reqProject.SponsorFee).
		SetCompany(reqProject.Edges.Company).
		Save(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	company, err := p.QueryCompany().WithTaskType().First(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	p.Edges.Company = company
	return p, nil
}

func (r *projectRepository) GetByID(projectID int) (*ent.Project, []*ent.Club, error) {
	p, err := r.db.Project.Query().
		Where(project.ID(projectID)).WithCompany(func(query *ent.CompanyQuery) {
		query.WithTaskType(func(query *ent.TaskTypeQuery) {
			query.Select().All(context.Background())
		})
	}).Only(context.Background())
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	log.Println(p.Edges.Company.Edges.TaskType)
	clubs, err := p.QueryProjectClub().QueryClub().
		WithLeader().All(context.Background())
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	projectLogs, err := p.QueryProjectLog().
		WithProjectLogFeedback().All(context.Background())
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	p.Edges.ProjectLog = projectLogs

	return p, clubs, nil
}

func (r *projectRepository) GetAll(limit, offset int) ([]*ent.Project, [][]*ent.Club, error) {
	p, err := r.db.Project.Query().
		Offset(offset).Limit(limit).WithCompany(func(query *ent.CompanyQuery) {
		query.WithTaskType(func(query *ent.TaskTypeQuery) {
			query.Select().All(context.Background())
		})
	}).
		All(context.Background())
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	allClubs := make([][]*ent.Club, len(p))
	for idx, project := range p {
		clubs, err := project.QueryProjectClub().QueryClub().WithLeader().All(context.Background())
		if err != nil {
			return nil, nil, errors.WithStack(err)
		}

		projectLogs, err := project.QueryProjectLog().
			WithProjectLogFeedback().All(context.Background())
		if err != nil {
			return nil, nil, errors.WithStack(err)
		}
		project.Edges.ProjectLog = projectLogs
		allClubs[idx] = clubs
	}

	return p, allClubs, nil
}

func (r *projectRepository) SaveProjectLog(reqPlog *ent.ProjectLog) (*ent.ProjectLog, error) {
	log.Println(reqPlog.Edges)
	pl, err := r.db.ProjectLog.Create().
		SetTitle(reqPlog.Title).
		SetAuthor(reqPlog.Author).
		SetContent(reqPlog.Content).
		SetStartDate(reqPlog.StartDate).
		SetEndDate(reqPlog.EndDate).
		SetCreatedAt(time.Now()).
		SetProject(reqPlog.Edges.Project).
		SetProjectClub(reqPlog.Edges.ProjectClub).
		Save(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	participants, err := pl.QueryProjectLogParticipant().All(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	feedbacks, err := pl.QueryProjectLogFeedback().All(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	pl.Edges.ProjectLogParticipant = participants
	pl.Edges.ProjectLogFeedback = feedbacks

	return pl, nil
}

func (r *projectRepository) SaveProjectLogFeedback(reqPfeedback *ent.ProjectLogFeedback) (*ent.ProjectLogFeedback, error) {
	plf, err := r.db.ProjectLogFeedback.Create().
		SetAuthor(reqPfeedback.Author).
		SetContent(reqPfeedback.Content).
		SetProjectLog(reqPfeedback.Edges.ProjectLog).Save(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return plf, nil
}

func (r *projectRepository) UploadProfileImage(reqProject *ent.Project) (*ent.Project, error) {
	p, err := r.db.Project.UpdateOneID(reqProject.ID).
		SetNillableProfileImage(reqProject.ProfileImage).Save(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return p, nil
}

func (r *projectRepository) GetAllTasks(projectID int) ([]*ent.TaskType, error) {
	tasks, err := r.db.TaskType.Query().
		Where(tasktype.HasProjectWith(project.ID(projectID))).All(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return tasks, nil
}

func (r *projectRepository) SaveTasks(p *ent.Project, taskType *ent.TaskType) (*ent.TaskType, error) {
	task, err := r.db.TaskType.Create().
		SetType(taskType.Type).
		SetProject(p).
		Save(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return task, nil
}
