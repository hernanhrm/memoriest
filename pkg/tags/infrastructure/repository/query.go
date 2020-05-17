package repository

import (
	"database/sql"
	"github.com/hernanhrm/memoriest/pkg/tags/application"
	"strings"
)

func GetGetAll(engine string, db *sql.DB) application.RepositoryGetAll {
	switch strings.TrimSpace(engine) {
	case postgres:
		return newPsqlGetAll(db)
	default:
		return nil
	}
}

func GetGetByID(engine string, db *sql.DB) application.RepositoryGetByID {
	switch strings.TrimSpace(engine) {
	case postgres:
		return newPsqlGetByID(db)
	default:
		return nil
	}
}
