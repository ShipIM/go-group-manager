package repository

import (
	"github.com/ShipIM/go-group-manager/internal/domain/entity"
)

type GroupRepository interface {
	CreateGroup(group entity.Group) (entity.Group, error)

	GetGroupByName(name string) (entity.Group, error)
	FindAllGroups() ([]entity.Group, error)

	UpdateGroup(group entity.Group) error

	DeleteGroupByName(name string) error
}

type StudentRepository interface {
	CreateStudent(student entity.Student) (entity.Student, error)

	GetStudentById(id int) (entity.Student, error)
	FindAllStudents() ([]entity.Student, error)

	UpdateStudent(student entity.Student) error

	DeleteStudentById(id int) error
}
