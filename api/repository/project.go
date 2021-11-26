package repository

import (
	"context"
	"log"
	"time"

	"github.com/Wanted-Linx/linx-backend/api/domain"
	"github.com/Wanted-Linx/linx-backend/api/ent"
	"github.com/Wanted-Linx/linx-backend/api/ent/project"
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

	// projecLogs, err := p.QueryProjectLog().All(context.Background())
	// if err != nil {
	// 	return nil, errors.WithStack(err)
	// }

	company, err := p.QueryCompany().First(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// clubs, err := p.QueryProjectClub().QueryClub().WithLeader().All(context.Background())
	// if err != nil {
	// 	return nil, errors.WithStack(err)
	// }

	// p.Edges.ProjectLog = projecLogs
	p.Edges.Company = company
	// p.Edges.Club = clubs
	return p, nil
}

func (r *projectRepository) GetByID(projectID int) (*ent.Project, []*ent.Club, error) {
	p, err := r.db.Project.Query().
		Where(project.ID(projectID)).
		WithCompany().
		// WithProjectLog().
		Only(context.Background())
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	clubs, err := p.QueryProjectClub().QueryClub().WithLeader().All(context.Background())
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	projectLogs, err := p.QueryProjectLog().All(context.Background())
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	p.Edges.ProjectLog = projectLogs

	return p, clubs, nil
}

func (r *projectRepository) GetAll(limit, offset int) ([]*ent.Project, [][]*ent.Club, error) {
	p, err := r.db.Project.Query().
		Offset(offset).Limit(limit).
		WithCompany().
		// WithProjectLog().
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
		projectLogs, err := project.QueryProjectLog().All(context.TODO())
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

	log.Println(pl)
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
