package service

import (
	"github.com/Wanted-Linx/linx-backend/api/domain"
	"github.com/Wanted-Linx/linx-backend/api/ent"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type studentService struct {
	studentRepo domain.StudentRepository
}

func NewStudentService(studentRepo domain.StudentRepository) domain.StudentService {
	return &studentService{studentRepo: studentRepo}
}

func (s *studentService) Save(userID int, reqSignup *domain.SignUpRequest) (*domain.StudentDto, error) {
	student := &ent.Student{
		ID:         userID,
		Name:       reqSignup.Name,
		University: reqSignup.University,
		Edges: ent.StudentEdges{
			User: &ent.User{ID: userID},
		},
	}

	newStudent, err := s.studentRepo.Save(student)
	if err != nil {
		return nil, errors.WithMessage(err, "알 수 없는 에러가 발생했습니다.")
	}

	log.Info("회원가입(학생) 완료", newStudent)
	return domain.StudentToDto(newStudent), nil
}

func (s *studentService) GetStudentByID(studentID int) (*domain.StudentDto, error) {

	return nil, nil
}
