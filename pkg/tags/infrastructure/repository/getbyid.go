package repository

import (
	"database/sql"
	"fmt"
	"github.com/hernanhrm/dbexec"
	dsl "github.com/hernanhrm/dslparser"
	"github.com/hernanhrm/memoriest/pkg/tags/domain"
)

const queryGetByID = "SELECT %s FROM tags WHERE id = $1"

var (
	getByIDFields = []string{"id", "name", "is_active", "created_at", "updated_at"}
)

type psqlGetByID struct {
	db       *sql.DB
	executor dbexec.SelectOneExecutor
}

func newPsqlGetByID(db *sql.DB) *psqlGetByID {
	e := dbexec.NewSelectOneExecutor(postgres, db)
	return &psqlGetByID{db: db, executor: e}
}

func (p psqlGetByID) GetByID(id uint, dslModel dsl.Model) (domain.GetByID, error) {
	dslfield := dsl.NewFieldParser(postgres, dslModel.Fields, getByIDFields)
	q := fmt.Sprintf(queryGetByID, dslfield.Parse(dsl.SelectQuery))

	if err := p.executor.SelectOne(q, id).Err(); err != nil {
		return domain.GetByID{}, err
	}

	m := domain.GetByID{}
	updatedAt := sql.NullTime{}
	if err := p.executor.ScanToFields(&m.ID, &m.Name, &m.IsActive, &m.CreatedAt, &updatedAt); err != nil {
		return domain.GetByID{}, err
	}
	m.UpdatedAt = updatedAt.Time

	return m, nil
}
