package db

import (
	"os"
	"path"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

func ConnectDatabase(databaseURL string) *sqlx.DB {
	db, err := sqlx.Open("postgres", databaseURL)
	if err != nil {
		panic(err)
	}

	executeMigrations(db)
	return db
}

func executeMigrations(db *sqlx.DB) {
	entries, err := os.ReadDir("db/migrations")
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		migration, err := os.ReadFile(path.Join("db", "migrations", entry.Name()))
		if err != nil {
			panic(err)
		}

		tx := db.MustBegin()
		tx.MustExec(string(migration))
		tx.Commit()
	}
}
