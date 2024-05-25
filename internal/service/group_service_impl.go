package service

import (
	"github.com/ShipIM/go-group-manager/internal/domain/entity"
	"github.com/ShipIM/go-group-manager/internal/repository"
)

type GroupServiceImpl struct {
	groupRepository repository.GroupRepository
}

func NewGroupService(groupRepository repository.GroupRepository) *GroupServiceImpl {
	return &GroupServiceImpl{groupRepository}
}

func (s *GroupServiceImpl) CreateGroup(group entity.Group) (entity.Group, error) {
	return s.groupRepository.CreateGroup(group)
}

func (s *GroupServiceImpl) GetGroupByName(name string) (entity.Group, error) {
	return s.groupRepository.GetGroupByName(name)
}

func (s *GroupServiceImpl) GetGroups(filter repository.GroupFilter) ([]entity.Group, error) {
	return s.groupRepository.FindAllGroups(filter)
}

func (s *GroupServiceImpl) UpdateGroup(group entity.Group) error {
	return s.groupRepository.UpdateGroup(group)
}

func (s *GroupServiceImpl) DeleteGroupByName(name string) error {
	return s.groupRepository.DeleteGroupByName(name)
}
