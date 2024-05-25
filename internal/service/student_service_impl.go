package service

import (
	"github.com/ShipIM/go-group-manager/internal/domain/entity"
	"github.com/ShipIM/go-group-manager/internal/repository"
)

type StudentServiceImpl struct {
	studentRepository repository.StudentRepository
}

func NewStudentService(studentRepository repository.StudentRepository) *StudentServiceImpl {
	return &StudentServiceImpl{studentRepository}
}

func (s *StudentServiceImpl) CreateStudent(student entity.Student) (entity.Student, error) {
	return s.studentRepository.CreateStudent(student)
}

func (s *StudentServiceImpl) GetStudentById(id int) (entity.Student, error) {
	return s.studentRepository.GetStudentById(id)
}

func (s *StudentServiceImpl) GetStudents(filter repository.StudentFilter) ([]entity.Student, error) {
	return s.studentRepository.FindAllStudents(filter)
}

func (s *StudentServiceImpl) UpdateStudent(student entity.Student) error {
	return s.studentRepository.UpdateStudent(student)
}

func (s *StudentServiceImpl) DeleteStudentById(id int) error {
	return s.studentRepository.DeleteStudentById(id)
}
