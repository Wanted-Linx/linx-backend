package domain

import "github.com/wanted-linx/linx-backend/api/ent"

type UserDto struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Kind  string `json:"kind"`
}

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Kind     string `json:"kind"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Kind     string `json:"kind"`
}

type UserService interface {
	SignUp(reqSignUp *SignUpRequest) (*UserDto, error)
	Login(reqLogin *LoginRequest) (*UserDto, error)
	GetUserByID(userID int) (*UserDto, error)
}

type UserReposiory interface {
	Create(user *ent.User) (*ent.User, error)
	GetByID(userID int) (*ent.User, error)
	FindByEmailAndPassword(email, password string) (*ent.User, error)
}

func UserToDto(src *ent.User) *UserDto {
	return &UserDto{
		ID:    src.ID,
		Email: src.Email,
		Kind:  src.Kind.String(),
	}
}
