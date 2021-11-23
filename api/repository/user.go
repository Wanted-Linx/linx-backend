package repository

import (
	"github.com/wanted-linx/linx-backend/api/domain"
	"github.com/wanted-linx/linx-backend/api/ent"
)

type userRepository struct {
	db *ent.Client
}

func NewUserRepository(db *ent.Client) domain.UserReposiory {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *ent.User) (*ent.User, error) {

	return nil, nil
}

func (r *userRepository) GetByID(userID int) (*ent.User, error) {
	return nil, nil
}
