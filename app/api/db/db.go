package db

import (
	"api/config"
	"os"
	"path"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

func ConnectDatabase(cfg *config.Config) *sqlx.DB {
	db, err := sqlx.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		panic(err)
	}

	executeMigrations(db, cfg.MigatrionsDirPath)
	return db
}

func executeMigrations(db *sqlx.DB, migrationsPath string) {
	entries, err := os.ReadDir(migrationsPath)
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		migration, err := os.ReadFile(path.Join(migrationsPath, entry.Name()))
		if err != nil {
			panic(err)
		}

		tx := db.MustBegin()
		tx.MustExec(string(migration))
		tx.Commit()
	}
}
