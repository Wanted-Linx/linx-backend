package domain

import "github.com/wanted-linx/linx-backend/api/ent"

type UserDto struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Kind  string `json:"kind"`
}

type SignUpRequest struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Kind     string `json:"kind"`
}

type UserService interface {
	SignUp(reqUser *SignUpRequest) (*UserDto, error)
	GetUserByID(userID int) (*UserDto, error)
}

type UserReposiory interface {
	Create(user *ent.User) (*ent.User, error)
	GetByID(userID int) (*ent.User, error)
}
