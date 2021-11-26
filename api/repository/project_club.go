package repository

import (
	"context"

	"github.com/Wanted-Linx/linx-backend/api/domain"
	"github.com/Wanted-Linx/linx-backend/api/ent"
	"github.com/pkg/errors"
)

type projectClubRepository struct {
	db *ent.Client
}

func NewProjectClubRepository(db *ent.Client) domain.ProjectClubRepository {
	return &projectClubRepository{db: db}
}

func (r *projectClubRepository) Apply(clubID, projectID int, startDate string) (*ent.ProjectClub, error) {
	pc, err := r.db.ProjectClub.Create().
		SetClubID(clubID).
		SetProjectID(projectID).
		SetStartDate(startDate).
		Save(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return pc, nil
}
