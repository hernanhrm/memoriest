package application

import (
	dsl "github.com/hernanhrm/dslparser"
	"github.com/hernanhrm/memoriest/pkg/tags/domain"
)

type ServicerGetByID interface {
	GetByID(id uint, dslModel dsl.Model) (domain.GetByID, error)
}

type RepositoryGetByID interface {
	GetByID(id uint, dslModel dsl.Model) (domain.GetByID, error)
}

type ServiceGetByID struct {
	repo RepositoryGetByID
}

func NewServiceGetByID(repo RepositoryGetByID) *ServiceGetByID {
	return &ServiceGetByID{repo: repo}
}

func (s ServiceGetByID) GetByID(id uint, dslModel dsl.Model) (domain.GetByID, error) {
	return s.repo.GetByID(id, dslModel)
}
