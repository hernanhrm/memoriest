package repository

import (
	"database/sql"
	"fmt"
	"github.com/hernanhrm/dbexec"
	"github.com/hernanhrm/dslparser"
	"github.com/hernanhrm/memoriest/pkg/tags/domain"
)

const psqlGetAll = "SELECT %s FROM tags %s %s LIMIT $1 OFFSET $2"

var (
	getAllFields  = []string{"id", "name", "is_active", "created_at", "updated_at"}
	getAllFilters = []string{"id:int", "is_active:bool"}
	getAllSorts   = []string{"id", "created_at", "updated_at"}
)

type psqlGetAllRepo struct {
	db       *sql.DB
	executor dbexec.SelectExecutor
}

func newPsqlGetAll(db *sql.DB) *psqlGetAllRepo {
	e := dbexec.NewSelectExecutor(postgres, db)
	return &psqlGetAllRepo{db: db, executor: e}
}

func (p psqlGetAllRepo) GetAll(dslModel dsl.Model) (domain.GetAlls, error) {
	dslfield := dsl.NewFieldParser(postgres, dslModel.Fields, getAllFields)
	dslfilter := dsl.NewFilterParser(postgres, dslModel.Filters, getAllFilters)
	dslsort := dsl.NewSortParser(postgres, dslModel.Sorts, getAllSorts)

	q := fmt.Sprintf(psqlGetAll, dslfield.Parse(dsl.SelectQuery), dslfilter.Parse(true), dslsort.Parse(true))

	err := p.executor.Select(q, dslModel.GetLimit(false), dslModel.GetOffset(false)).Err()
	if err != nil {
		return nil, err
	}
	fmt.Println(p.executor.GetRows().ColumnTypes())
	models := domain.GetAlls{}
	for p.executor.GetRows().Next() {
		m := domain.GetAll{}
		updatedAt := sql.NullTime{}
		err = p.executor.GetRows().Scan(&m.ID, &m.Name, &m.IsActive, &m.CreatedAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		m.UpdatedAt = updatedAt.Time

		models = append(models, &m)
	}

	return models, nil
}
