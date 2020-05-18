package repository

import (
	"database/sql"
	"github.com/hernanhrm/dbexec"
)

const psqlDeleteByID = "DELETE FROM tags WHERE id = $1"

type psqlDeleteByIDRepo struct {
	db       *sql.DB
	executor dbexec.DeleteExecutor
}

func newPsqlDeleteByIDRepo(db *sql.DB) *psqlDeleteByIDRepo {
	e := dbexec.NewDeleteExecutor(postgres, db)
	return &psqlDeleteByIDRepo{db: db, executor: e}
}

func (p psqlDeleteByIDRepo) DeleteByID(id uint) error {
	return p.executor.Delete(psqlDeleteByID, id).Err()
}
