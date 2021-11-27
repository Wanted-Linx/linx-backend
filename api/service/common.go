package service

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Wanted-Linx/linx-backend/api/domain"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// type ProfileImage struct {
// 	studentProfileImage *domain.StudentProfileImage
// 	companyProfileImage *domain.CompanyProfileImage
// 	projectProfileImage *domain.ProjectProfileImage
// 	clubProfileImage    *domain.ProfileImage
// }

func UploadImage(targetID int, dir, key string, reqImage *domain.ProfileImageRequest) ([]byte, error) {
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

		// // Copy
		// if _, err = io.Copy(dst, src); err != nil {
		// 	return nil, errors.Wrap(err, "프로필 이미지 업로드 실패")
		// }

		// body := bytes.NewReader(fileBytes)
	}

	// TODO: db에 profile iamge 주소 저장

	return fileBytes, nil
}
