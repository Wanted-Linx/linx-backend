package handler

import "github.com/Wanted-Linx/linx-backend/api/domain"

type StudentHandler struct {
	studentService domain.StudentService
}

func NewStudentHandler(studentService domain.StudentService) *StudentHandler {
	return &StudentHandler{studentService: studentService}
}
