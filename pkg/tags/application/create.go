package application

import (
	"github.com/hernanhrm/memoriest/pkg/tags/domain"
)

type ServicerCreate interface {
	Create(model *domain.Command) error
}

type RepositoryCreate interface {
	Create(model *domain.Command) error
}

type ServiceCreate struct {
	repo RepositoryCreate
}

func NewCreateService(repo RepositoryCreate) *ServiceCreate {
	return &ServiceCreate{repo: repo}
}

func (s ServiceCreate) Create(model *domain.Command) error {
	return s.repo.Create(model)
}
