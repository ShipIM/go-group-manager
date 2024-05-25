package service

import (
	"github.com/ShipIM/go-group-manager/internal/domain/entity"
	"github.com/ShipIM/go-group-manager/internal/repository"
)

type GroupService interface {
	CreateGroup(group entity.Group) (entity.Group, error)

	GetGroupByName(name string) (entity.Group, error)
	GetGroups(filter repository.GroupFilter) ([]entity.Group, error)

	UpdateGroup(group entity.Group) error

	DeleteGroupByName(name string) error
}

type StudentService interface {
	CreateStudent(student entity.Student) (entity.Student, error)

	GetStudentById(id int) (entity.Student, error)
	GetStudents(filter repository.StudentFilter) ([]entity.Student, error)

	UpdateStudent(student entity.Student) error

	DeleteStudentById(id int) error
}

type Service struct {
	GroupService
	StudentService
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		GroupService:   NewGroupService(repository.GroupRepository),
		StudentService: NewStudentService(repository.StudentRepository),
	}
}
