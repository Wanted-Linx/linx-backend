package repository

import (
	"context"

	"github.com/Wanted-Linx/linx-backend/api/domain"
	"github.com/Wanted-Linx/linx-backend/api/ent"
	"github.com/Wanted-Linx/linx-backend/api/ent/student"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type studentRepository struct {
	db *ent.Client
}

func NewStudentRepository(db *ent.Client) domain.StudentRepository {
	return &studentRepository{db: db}
}

func (r *studentRepository) Save(reqStudent *ent.Student) (*ent.Student, error) {
	s, err := r.db.Student.Create().
		SetID(reqStudent.ID).
		SetName(reqStudent.Name).
		SetUniversity(reqStudent.University).
		SetUser(reqStudent.Edges.User).
		SetNillableProfileLink(reqStudent.ProfileLink).
		SetNillableProfileImage(reqStudent.ProfileImage).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	user, err := s.QueryUser().First(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	s.Edges.User = user
	return s, nil
}

func (r *studentRepository) GetByID(studentID int, reqStudent *ent.Student) (*ent.Student, error) {
	s, err := r.db.Student.Query().
		Where(student.ID(studentID)).
		WithClubMember().
		WithUser().
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	clubs, err := s.QueryClubMember().QueryClub().WithLeader().All(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	log.Info(clubs)
	s.Edges.Club = clubs
	log.Info(s.Edges.Club)
	return s, nil
}

func (r *studentRepository) UploadProfileImage(reqStudent *ent.Student) (*ent.Student, error) {
	s, err := r.db.Student.UpdateOneID(reqStudent.ID).
		SetNillableProfileImage(reqStudent.ProfileImage).Save(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return s, nil
}

func (r *studentRepository) UpdateProfile(reqStudent *ent.Student) (*ent.Student, error) {
	s, err := r.db.Student.UpdateOneID(reqStudent.ID).
		SetNillableProfileLink(reqStudent.ProfileLink).Save(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	user, err := s.QueryUser().First(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	clubs, err := s.QueryClubMember().QueryClub().WithLeader().All(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	s.Edges.Club = clubs
	s.Edges.User = user
	return s, nil
}
