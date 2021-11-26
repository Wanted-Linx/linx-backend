package service

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Wanted-Linx/linx-backend/api/domain"
	"github.com/Wanted-Linx/linx-backend/api/ent"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type studentService struct {
	studentRepo    domain.StudentRepository
	clubMemberRepo domain.ClubMemberRepository
}

func NewStudentService(studentRepo domain.StudentRepository, clubMeberRepo domain.ClubMemberRepository) domain.StudentService {
	return &studentService{
		studentRepo:    studentRepo,
		clubMemberRepo: clubMeberRepo,
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

	// TODO: interested_type 추가하고 eager loading...
	// InterestedType의 요청이 있다면
	// if len(reqStudent.InterestedType) > 0 {
	// 	// s.studentRepo.SaveInterestedType(interestedType *ent.StudentInterestedType)
	// }

	// TODO: repository layer에 getClubs로 따로 빼기
	// clubs, err := updatedStudent.QueryClubMember().QueryClub().All(context.Background())
	// if err != nil {
	// 	return nil, errors.WithMessage(err, "알 수 없는 오류가 발생했습니다.")
	// }

	return domain.StudentToDto(updatedStudent, nil), nil
}

func (s *studentService) UploadProfileImage(studentID int, reqImage *domain.StudentProfileImage) ([]byte, error) {
	var fileBytes []byte
	for _, imageFile := range reqImage.Image {
		// Source
		src, err := imageFile.Open()
		if err != nil {
			return nil, errors.Wrap(err, "프로필 이미지 업로드 실패")
		}
		defer src.Close()

		fileBytes, err = ioutil.ReadAll(src)
		if err != nil {
			return nil, errors.Wrap(err, "프로필 이미지 업로드 실패")
		}

		// Destination
		dir := fmt.Sprintf("./students/profile/%d/image", studentID)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			os.MkdirAll(dir, 0700) // Create your file
		}

		key := uuid.New().String()
		dst, err := os.Create(filepath.Join(dir, filepath.Base(key)))
		if err != nil {
			return nil, errors.Wrap(err, "프로필 이미지 업로드 실패")
		}
		defer dst.Close()

		img, fileType, err := image.Decode(bytes.NewReader(fileBytes))
		if err != nil {
			return nil, errors.Wrap(err, "프로필 이미지 업로드 실패")
		}

		switch fileType {
		case "jpeg":
			log.Println(fileType)
			err = jpeg.Encode(dst, img, nil)
			if err != nil {
				return nil, errors.Wrap(err, "프로필 이미지 업로드 실패")
			}
		default:
			err = png.Encode(dst, img)
			if err != nil {
				return nil, errors.Wrap(err, "프로필 이미지 업로드 실패")
			}
		}

		student := &ent.Student{
			ID:           studentID,
			ProfileImage: &key,
		}
		s, err := s.studentRepo.UploadProfileImage(student)
		if err != nil {
			return nil, errors.WithMessage(err, "프로필 이미지 업로드 실패")
		}

		log.Info("프로필 이미지 업로드 성공", s)

		// // Copy
		// if _, err = io.Copy(dst, src); err != nil {
		// 	return nil, errors.Wrap(err, "프로필 이미지 업로드 실패")
		// }

		// body := bytes.NewReader(fileBytes)
	}

	// TODO: db에 profile iamge 주소 저장

	return fileBytes, nil
}

func (s *studentService) GetProfileImage(studentID int) ([]byte, error) {
	// TODO: key is image path from db
	// key := "sample_images_09.jpg"

	getStudent, err := s.GetStudentByID(studentID)
	if err != nil {
		return nil, err
	}

	if getStudent.ProfileImage == nil {
		return nil, errors.New("프로필 이미지가 존재하지 않습니다.")
	}
	fileBytes, err := ioutil.ReadFile(fmt.Sprintf("./students/profile/%d/image/%s", studentID, *getStudent.ProfileImage))
	if err != nil {
		return nil, errors.Wrap(err, "프로필 이미지 조회 실패")
	}

	return fileBytes, nil
}
