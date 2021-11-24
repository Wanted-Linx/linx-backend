package repository

import (
	"context"

	"github.com/Wanted-Linx/linx-backend/api/domain"
	"github.com/Wanted-Linx/linx-backend/api/ent"
	"github.com/pkg/errors"
)

type clubMemberRepository struct {
	db *ent.Client
}

func NewClubMemberRepository(db *ent.Client) domain.ClubMemberRepository {
	return &clubMemberRepository{db: db}
}

func (r *clubMemberRepository) Register(studentID, clubID int) (*ent.ClubMember, error) {
	cm, err := r.db.ClubMember.Create().
		SetStudentID(studentID).
		SetClubID(clubID).
		Save(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return cm, nil
}

// func (r *clubMemberRepository) FindbyClubID(clubID int) ([]*ent.ClubMember, error) {
// 	cm, err := r.db.ClubMember.Query().
// 		Where(clubmember.HasClubWith(club.ID(clubID))).
// 		WithStudent().
// 		WithClub().
// 		All(context.Background())
// 	if err != nil {
// 		return nil, err
// 	}

// 	return cm, nil
// }
// func (r *clubMemberRepository) FindByStudentIDAndClubID(studentID, clubID int) (*ent.ClubMember, error) {
// 	cm, err := r.db.ClubMember.Query().
// 		Where(clubmember.And(clubmember.HasStudentWith(student.ID(studentID)), clubmember.HasClubWith(club.ID(clubID)))).
// 		WithStudent().
// 		WithClub().
// 		Only(context.Background())
// 	if err != nil {
// 		return nil, err
// 	}

// 	return cm, nil
// }
