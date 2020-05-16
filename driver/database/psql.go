package database

import (
	"database/sql"
	"fmt"
	"github.com/hernanhrm/memoriest/config"
	"sync"
)

var (
	once sync.Once
	db   *sql.DB
	err  error
)

func GetPsqlConnection(c config.Model) error {
	once.Do(func() {
		dns := fmt.Sprintf(
			"postgres://%s:%s@%s:%d/%s?sslmode=disable",
			c.Databases["postgres"]["user"],
			c.Databases["postgres"]["password"],
			c.Databases["postgres"]["server"],
			c.Databases["postgres"]["port"],
			c.Databases["postgres"]["dbname"],
		)

		db, err = sql.Open("postgres", dns)
	})
	if err != nil {
		return err
	}

	return nil
}

// GetConnection return an unique instance of sql.DB
func GetConnection() *sql.DB {
	return db
}

// Close pool connection
func Close() {
	db.Close()
}
