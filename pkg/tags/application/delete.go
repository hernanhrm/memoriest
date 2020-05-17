package application

type ServicerDeleteByID interface {
	DeleteByID(id uint) error
}

type RepositoryDeleteByID interface {
	DeleteByID(id uint) error
}

type ServiceDeleteByID struct {
	repo RepositoryDeleteByID
}

func NewDeleteByIDService(repo RepositoryDeleteByID) *ServiceDeleteByID {
	return &ServiceDeleteByID{repo: repo}
}

func (s ServiceDeleteByID) DeleteByID(id uint) error {
	return s.repo.DeleteByID(id)
}
