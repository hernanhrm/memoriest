package repository

import (
	"database/sql"
	"fmt"
	"github.com/hernanhrm/dbexec"
	dsl "github.com/hernanhrm/dslparser"
	"github.com/hernanhrm/memoriest/pkg/tags/domain"
)

const psqlUpdateByID = "UPDATE tags SET %s WHERE id = $%d"

var updateFields = []string{"name", "is_active"}

type psqlUpdateRepo struct {
	db       *sql.DB
	executor dbexec.UpdateExecutor
	dslfield dsl.FieldParser
}

func newPsqlUpdateRepo(db *sql.DB) *psqlUpdateRepo {
	e := dbexec.NewUpdateExecutor(postgres, db)
	d := dsl.NewFieldParser(postgres, nil, updateFields)
	return &psqlUpdateRepo{db: db, executor: e, dslfield: d}
}

func (p psqlUpdateRepo) UpdateByID(model domain.Command) error {
	query := fmt.Sprintf(psqlUpdateByID, p.dslfield.Parse(dsl.UpdateQuery), len(updateFields)+1)
	return p.executor.Update(query, model.Name, model.IsActive, model.ID).Err()
}
