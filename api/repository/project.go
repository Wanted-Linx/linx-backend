package repository

import (
	"context"

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

	projecLogs, err := p.QueryProjectLog().All(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// clubs, err := p.QueryProjectClub().QueryClub().WithLeader().All(context.Background())
	// if err != nil {
	// 	return nil, errors.WithStack(err)
	// }

	company, err := p.QueryCompany().First(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	p.Edges.ProjectLog = projecLogs
	p.Edges.Company = company
	// log.Info(p, clubs, p.Edges.Company)
	// p.Edges.Club = clubs
	return p, nil
}

func (r *projectRepository) GetByID(projectID int) (*ent.Project, []*ent.Club, error) {
	p, err := r.db.Project.Query().
		Where(project.ID(projectID)).
		WithCompany().
		WithProjectLog().
		Only(context.Background())
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	clubs, err := p.QueryProjectClub().QueryClub().WithLeader().All(context.Background())
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	return p, clubs, nil
}

func (r *projectRepository) GetAll(limit, offset int) ([]*ent.Project, [][]*ent.Club, error) {
	p, err := r.db.Project.Query().
		Offset(offset).Limit(limit).
		WithCompany().
		WithProjectLog().
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

		allClubs[idx] = clubs
	}

	return p, allClubs, nil
}
