package db

import (
	"github.com/golang-migrate/migrate"
	psqlMigrate "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

func MustMigrate(psql *sqlx.DB, path string) {
	var config Config
	envconfig.MustProcess("", &config)

	driver, err := psqlMigrate.WithInstance(psql.DB, &psqlMigrate.Config{
		DatabaseName: config.Database,
	})
	if err != nil {
		panic(err)
	}

	migrationContext, err := migrate.NewWithDatabaseInstance(path, "postgres", driver)
	if err != nil {
		panic(err)
	}

	err = migrationContext.Up()
	if err != nil && err != migrate.ErrNoChange {
		panic(err)
	}
}
