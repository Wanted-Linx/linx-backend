package domain

import (
	"mime/multipart"
)

type ProfileImageRequest struct {
	Image []*multipart.FileHeader
}

type ProfileImageDto struct {
	ProfileImage *string `json:"profile_image"`
}
