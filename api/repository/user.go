package repository

import (
	"context"

	"github.com/wanted-linx/linx-backend/api/domain"
	"github.com/wanted-linx/linx-backend/api/ent"
	"github.com/wanted-linx/linx-backend/api/ent/user"
	"github.com/wanted-linx/linx-backend/api/util"
)

type userRepository struct {
	db *ent.Client
}

func NewUserRepository(db *ent.Client) domain.UserReposiory {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *ent.User) (*ent.User, error) {
	u, err := r.db.User.Create().
		SetEmail(user.Email).
		SetPassword(user.Password).
		SetKind(user.Kind).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *userRepository) GetByID(userID int) (*ent.User, error) {
	u, err := r.db.User.Query().
		Where(user.ID(userID)).
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *userRepository) FindByEmailAndPassword(email, password string) (*ent.User, error) {
	u, err := r.db.User.Query().
		Where(user.Email(email)).
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	if err := util.VerifyPassword(u.Password, password); err != nil {
		return nil, err
	}
	return u, nil
}
