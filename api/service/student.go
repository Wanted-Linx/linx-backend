package service

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Wanted-Linx/linx-backend/api/domain"
	"github.com/Wanted-Linx/linx-backend/api/ent"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type studentService struct {
	studentRepo    domain.StudentRepository
	clubMemberRepo domain.ClubMemberRepository
	taskTypeRepo   domain.TaskTypeRepository
}

func NewStudentService(studentRepo domain.StudentRepository,
	clubMeberRepo domain.ClubMemberRepository,
	taskTypeRepo domain.TaskTypeRepository) domain.StudentService {
	return &studentService{
		studentRepo:    studentRepo,
		clubMemberRepo: clubMeberRepo,
		taskTypeRepo:   taskTypeRepo,
	}
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

	// 가입할 때는 따로 존재하지 않는다...
	// tasks, err := s.studentRepo.GetAllTasks(newStudent.ID)
	// if err != nil {
	// 	return nil, errors.WithMessage(err, "알 수 없는 오류가 발생했습니다.")
	// }

	// newStudent.Edges.TaskType = tasks

	log.Info("회원가입(학생) 완료", newStudent)
	return domain.StudentToDto(newStudent, nil), nil
}

func (s *studentService) GetStudentByID(studentID int) (*domain.StudentDto, error) {
	student := &ent.Student{
		ID: studentID,
	}

	getStudent, err := s.studentRepo.GetByID(studentID, student)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.WithMessage(err, "해당하는 학생을 찾을 수 없습니다.")
		}
		return nil, errors.WithMessage(err, "알 수 없는 오류가 발생했습니다.")
	}

	tasks, err := s.studentRepo.GetAllTasks(getStudent.ID)
	if err != nil {
		return nil, errors.WithMessage(err, "알 수 없는 오류가 발생했습니다.")
	}

	getStudent.Edges.TaskType = tasks
	log.Info("해당하는 학생 조회 완료", getStudent)
	return domain.StudentToDto(getStudent, nil), nil
}

func (s *studentService) UpdateProfile(studentID int, reqStudent *domain.StudentProfileUpdate) (*domain.StudentDto, error) {
	student := &ent.Student{
		ID:          studentID,
		ProfileLink: &reqStudent.ProfileLink,
	}

	updatedStudent, err := s.studentRepo.UpdateProfile(student)
	if err != nil {
		return nil, errors.WithMessage(err, "알 수 없는 에러가 발생했습니다.")
	}

	tasks := make([]*ent.TaskType, len(reqStudent.InterestedType))
	for i := 0; i < len(reqStudent.InterestedType); i++ {
		interestedType := &ent.TaskType{
			Type: reqStudent.InterestedType[i],
		}

		task, err := s.studentRepo.SaveTasks(updatedStudent, interestedType)
		if err != nil {
			return nil, errors.WithMessage(err, "알 수 없는 오류가 발생했습니다.")
		}
		tasks[i] = task
	}

	updatedStudent.Edges.TaskType = tasks
	// TODO: interested_type 추가하고 eager loading...
	// InterestedType의 요청이 있다면
	// if len(reqStudent.InterestedType) > 0 {
	// 	// s.studentRepo.SaveInterestedType(interestedType *ent.StudentInterestedType)
	// }

	return domain.StudentToDto(updatedStudent, nil), nil
}

func (s *studentService) UploadProfileImage(studentID int, reqImage *domain.ProfileImageRequest) ([]byte, error) {
	// Destination
	dir := fmt.Sprintf("./students/profile/%d/image", studentID)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0700) // Create your file
		if err != nil {
			return nil, errors.Wrap(err, "이미지 저장용 디렉토리 생성 실패")
		}
	}

	key := uuid.New().String()
	fileBytes, err := UploadImage(studentID, dir, key, reqImage)
	if err != nil {
		return nil, errors.Wrap(err, "프로필 이미지 업로드 실패")
	}

	student := &ent.Student{
		ID:           studentID,
		ProfileImage: &key,
	}

	ups, err := s.studentRepo.UploadProfileImage(student)
	if err != nil {
		return nil, errors.WithMessage(err, "프로필 이미지 업로드 실패")
	}

	log.Info("프로필 이미지 업로드 성공", ups)
	return fileBytes, nil
}

func (s *studentService) GetProfileImage(studentID int) ([]byte, error) {
	student, err := s.GetStudentByID(studentID)
	if err != nil {
		return nil, err
	}

	if student.ProfileImage == nil {
		return nil, errors.New("프로필 이미지가 존재하지 않습니다.")
	}

	fileBytes, err := ioutil.ReadFile(fmt.Sprintf("./students/profile/%d/image/%s", studentID, *student.ProfileImage))
	if err != nil {
		return nil, errors.Wrap(err, "프로필 이미지 조회 실패")
	}

	return fileBytes, nil
}
