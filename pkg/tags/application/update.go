package application

import (
	"github.com/hernanhrm/memoriest/pkg/tags/domain"
)

type ServiceUpdateByID struct {
	repo RepositoryUpdateByID
}

func NewServiceUpdateByID(repo RepositoryUpdateByID) *ServiceUpdateByID {
	return &ServiceUpdateByID{repo: repo}
}

func (s ServiceUpdateByID) UpdateByID(model domain.Command) error {
	return s.repo.UpdateByID(model)
}

type ServicerUpdateByID interface {
	UpdateByID(model domain.Command) error
}

type RepositoryUpdateByID interface {
	UpdateByID(model domain.Command) error
}
