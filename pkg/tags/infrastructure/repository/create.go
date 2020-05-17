package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/hernanhrm/dbexec"
	dsl "github.com/hernanhrm/dslparser"
	"github.com/hernanhrm/memoriest/pkg/tags/domain"
)

const psqlInsert = `INSERT INTO tags %s RETURNING id, created_at`

var insertFields = []string{"name", "is_active"}

type psqlCreateRepo struct {
	db       *sql.DB
	executor dbexec.InsertExecutor
	dslfield dsl.FieldParser
}

func newPsqlCreateRepo(db *sql.DB) *psqlCreateRepo {
	ex := dbexec.NewInsertExecutor(postgres, db)
	d := dsl.NewFieldParser(postgres, nil, insertFields)
	return &psqlCreateRepo{db: db, executor: ex, dslfield: d}
}

func (p *psqlCreateRepo) Create(model *domain.Command) error {
	if p == nil {
		return errors.New("dsl is nil")
	}

	query := fmt.Sprintf(psqlInsert, p.dslfield.Parse(dsl.InsertQuery))
	if p.executor == nil {
		return errors.New("executor is nil")
	}

	if err := p.executor.Insert(query, model.Name,  model.IsActive).Err(); err !=nil {
		return err
	}

	return p.executor.GetRow().Scan(&model.ID, &model.CreatedAt)
}
