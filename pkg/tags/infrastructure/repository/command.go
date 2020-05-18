package repository

import (
	"database/sql"
	"github.com/hernanhrm/memoriest/pkg/tags/application"
	"strings"
)

func GetCreateRepo(engine string, db *sql.DB) application.RepositoryCreate {
	switch strings.TrimSpace(engine) {
	case postgres:
		return newPsqlCreateRepo(db)
	default:
		return nil
	}
}

func GetUpdateRepo(engine string, db *sql.DB) application.RepositoryUpdateByID {
	switch strings.TrimSpace(engine) {
	case postgres:
		return newPsqlUpdateRepo(db)
	default:
		return nil
	}
}

func GetDeleteByIDRepo(engine string, db *sql.DB) application.RepositoryDeleteByID {
	switch strings.TrimSpace(engine) {
	case postgres:
		return newPsqlDeleteByIDRepo(db)
	default:
		return nil
	}
}
