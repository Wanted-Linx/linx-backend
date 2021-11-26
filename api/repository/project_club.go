package repository

import (
	"context"
	"log"

	"github.com/Wanted-Linx/linx-backend/api/domain"
	"github.com/Wanted-Linx/linx-backend/api/ent"
	"github.com/Wanted-Linx/linx-backend/api/ent/projectclub"
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

func (r *projectClubRepository) Get(projectID, clubID int) (*ent.ProjectClub, error) {
	pc, err := r.db.ProjectClub.Query().
		Where(projectclub.And(projectclub.ProjectID(projectID), projectclub.ClubID(clubID))).
		Only(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	log.Println(pc)
	return pc, err
}
