package service

import "github.com/wanted-linx/linx-backend/api/domain"

type userService struct {
	userRepo domain.UserReposiory
}

func NewUserSerivce(userRepo domain.UserReposiory) domain.UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) SignUp(reqSignUp *domain.SignUpRequest) (*domain.UserDto, error) {
	return nil, nil
}

func (s *userService) GetUserByID(userID int) (*domain.UserDto, error) {
	return nil, nil
}
