package service

import (
	"github.com/ShipIM/go-group-manager/internal/domain/entity"
	"github.com/ShipIM/go-group-manager/internal/repository"
)

type GroupService struct {
	groupRepository repository.GroupRepository
}

func NewGroupService(groupRepository repository.GroupRepository) *GroupService {
	return &GroupService{groupRepository}
}

func (s *GroupService) CreateGroup(group entity.Group) (entity.Group, error) {
	return s.groupRepository.CreateGroup(group)
}

func (s *GroupService) GetGroupByName(name string) (entity.Group, error) {
	return s.groupRepository.GetGroupByName(name)
}

func (s *GroupService) GetGroups(course *int, grade *string) ([]entity.Group, error) {
	return s.groupRepository.FindAllGroups(course, grade)
}

func (s *GroupService) UpdateGroup(group entity.Group) error {
	return s.groupRepository.UpdateGroup(group)
}

func (s *GroupService) DeleteGroupByName(name string) error {
	return s.groupRepository.DeleteGroupByName(name)
}
