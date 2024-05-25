package service

import (
	"github.com/ShipIM/go-group-manager/internal/domain/entity"
	"github.com/ShipIM/go-group-manager/internal/repository"
)

type StudentService struct {
	studentRepository repository.StudentRepository
}

func NewStudentService(studentRepository repository.StudentRepository) *StudentService {
	return &StudentService{studentRepository}
}

func (s *StudentService) CreateStudent(student entity.Student) (entity.Student, error) {
	return s.studentRepository.CreateStudent(student)
}

func (s *StudentService) GetStudentById(id int) (entity.Student, error) {
	return s.studentRepository.GetStudentById(id)
}

func (s *StudentService) GetStudents() ([]entity.Student, error) {
	return s.studentRepository.FindAllStudents()
}

func (s *StudentService) UpdateStudent(student entity.Student) error {
	return s.studentRepository.UpdateStudent(student)
}

func (s *StudentService) DeleteStudentById(id int) error {
	return s.studentRepository.DeleteStudentById(id)
}
