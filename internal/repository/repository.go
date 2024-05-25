package repository

import (
	"database/sql"

	"github.com/ShipIM/go-group-manager/internal/domain/entity"
)

type GroupFilter struct {
	Course *int
	Grade  *string
}

type GroupRepository interface {
	CreateGroup(group entity.Group) (entity.Group, error)

	GetGroupByName(name string) (entity.Group, error)
	FindAllGroups(filter GroupFilter) ([]entity.Group, error)

	UpdateGroup(group entity.Group) error

	DeleteGroupByName(name string) error
}

type StudentFilter struct {
	Name       *string
	Surname    *string
	Patronymic *string
	GroupName  *string
	Age        *int
}

type StudentRepository interface {
	CreateStudent(student entity.Student) (entity.Student, error)

	GetStudentById(id int) (entity.Student, error)
	FindAllStudents(filter StudentFilter) ([]entity.Student, error)

	UpdateStudent(student entity.Student) error

	DeleteStudentById(id int) error
}

type Repository struct {
	GroupRepository
	StudentRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		GroupRepository:   NewGroupRepository(db),
		StudentRepository: NewStudentRepository(db),
	}
}

func deref[T any](ptr *T) interface{} {
	if ptr == nil {
		return nil
	}

	return *ptr
}
