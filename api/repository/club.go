package repository

import (
	"context"
	"log"

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

func (r *clubRepository) GetByID(clubID int) (*ent.Club, []*ent.Student, error) {
	c, err := r.db.Club.Query().
		Where(club.ID(clubID)).
		WithLeader().
		Only(context.Background())
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	members, err := c.QueryClubMember().QueryStudent().
		All(context.Background())
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	projects, err := c.QueryProjectClub().QueryProject().
		WithProjectLog().All(context.Background())
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	for _, project := range projects {
		project.Edges.Company, err = project.QueryCompany().First(context.TODO())
		if err != nil {
			return nil, nil, errors.WithStack(err)
		}
	}

	c.Edges.Project = projects
	log.Println(c.Edges, members)
	return c, members, nil
}

func (r *clubRepository) GetAll(limit, offset int) ([]*ent.Club, [][]*ent.Student, error) {
	c, err := r.db.Club.Query().
		Offset(offset).Limit(limit).
		WithLeader().
		All(context.Background())
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	allMembers := make([][]*ent.Student, len(c))
	// allProjects
	for idx, club := range c {
		members, err := club.QueryClubMember().
			QueryStudent().All(context.Background())
		if err != nil {
			return nil, nil, errors.WithStack(err)
		}

		club.Edges.Project, err = club.QueryProjectClub().QueryProject().
			WithCompany().WithProjectLog().All(context.Background())
		if err != nil {
			return nil, nil, errors.WithStack(err)
		}

		allMembers[idx] = members
	}

	return c, allMembers, nil
}

func (r *clubRepository) UploadProfileImage(reqClub *ent.Club) (*ent.Club, error) {
	c, err := r.db.Club.UpdateOneID(reqClub.ID).
		SetNillableProfileImage(reqClub.ProfileImage).Save(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return c, nil
}
