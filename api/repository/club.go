package repository

import (
	"context"

	"github.com/Wanted-Linx/linx-backend/api/domain"
	"github.com/Wanted-Linx/linx-backend/api/ent"
	"github.com/Wanted-Linx/linx-backend/api/ent/club"
	"github.com/pkg/errors"
)

type clubRepository struct {
	db *ent.Client
}

func NewClubRepository(db *ent.Client) domain.ClubRepository {
	return &clubRepository{db: db}
}

func (r *clubRepository) Save(reqClub *ent.Club) (*ent.Club, error) {
	c, err := r.db.Club.Create().
		SetName(reqClub.Name).
		SetLeader(reqClub.Edges.Leader).
		SetOrganization(reqClub.Organization).
		SetDescription(reqClub.Description).
		SetNillableProfileLink(reqClub.ProfileLink).
		SetNillableProfileImage(reqClub.ProfileImage).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	leader, err := c.QueryLeader().First(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	c.Edges.Leader = leader
	return c, nil
}

func (r *clubRepository) GetByID(clubID int) (*ent.Club, error) {
	c, err := r.db.Club.Query().
		Where(club.ID(clubID)).
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	return c, nil
}
