package application

import (
	dsl "github.com/hernanhrm/dslparser"
	"github.com/hernanhrm/memoriest/pkg/tags/domain"
)

type ServicerGetAll interface {
	GetAll(dslModel dsl.Model) (domain.GetAlls, error)
}

type RepositoryGetAll interface {
	GetAll(dslModel dsl.Model) (domain.GetAlls, error)
}

type ServiceGetAll struct {
	repo RepositoryGetAll
}

func NewServiceGetAll(repo RepositoryGetAll) *ServiceGetAll {
	return &ServiceGetAll{repo: repo}
}

func (s ServiceGetAll) GetAll(dslModel dsl.Model) (domain.GetAlls, error) {
	return s.repo.GetAll(dslModel)
}
