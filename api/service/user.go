package service

import (
	"fmt"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/wanted-linx/linx-backend/api/domain"
	"github.com/wanted-linx/linx-backend/api/ent"
	"github.com/wanted-linx/linx-backend/api/ent/user"
	"github.com/wanted-linx/linx-backend/api/util"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo domain.UserReposiory
}

func NewUserSerivce(userRepo domain.UserReposiory) domain.UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) SignUp(reqSignUp *domain.SignUpRequest) (*domain.UserDto, error) {
	hashPassword, err := util.Hash(reqSignUp.Password)
	if err != nil {
		return nil, errors.WithMessage(err, "알 수 없는 오류가 발생했습니다.")
	}

	u := &ent.User{
		Email:    reqSignUp.Email,
		Password: hashPassword,
		Kind:     user.Kind(reqSignUp.Kind),
	}

	newUser, err := s.userRepo.Create(u)
	if err != nil {
		if errors.Is(err, err.(*ent.ConstraintError)) {
			return nil, errors.WithMessage(err, "이미 존재하는 이메일입니다.")
		}
		return nil, errors.WithMessage(err, "알 수 없는 에러가 발생했습니다.")
	}

	if newUser.Kind.String() == "student" {
		// register student table
		log.Info("회원가입(학생) 성공")
	} else {
		// register company table
		log.Info("회원가입(기업) 성공")
	}

	return domain.UserToDto(newUser), nil
}

func (s *userService) Login(reqUser *domain.LoginRequest) (*domain.UserDto, error) {
	u, err := s.userRepo.FindByEmailAndPassword(reqUser.Email, reqUser.Password)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.WithMessage(err, "존재하지 않는 유저입니다.")
		} else if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, errors.WithMessage(err, "잘못된 비밀번호입니다.")
		}

		return nil, errors.WithMessage(err, "알 수 없는 에러가 발생했습니다.")
	}

	if u.Kind.String() != reqUser.Kind {
		return nil, errors.WithStack(errors.New((fmt.Sprintf("잘못된 계정 유형입니다. %s 계정으로 로그인 해주세요", u.Kind.String()))))
	}

	log.Info("로그인 성공")
	return domain.UserToDto(u), nil
}

func (s *userService) GetUserByID(userID int) (*domain.UserDto, error) {
	u, err := s.userRepo.GetByID(userID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.WithMessage(err, "존재하지 않는 유저입니다.")
		}

		return nil, errors.WithMessage(err, "알 수 없는 에러가 발생했습니다.")
	}

	return domain.UserToDto(u), nil
}
