package migrations

import (
	"context"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	migrateSql "github.com/golang-migrate/migrate/database/mysql"
	"github.com/projects/sys-des/txn-routine/pkg/env"

	_ "github.com/golang-migrate/migrate/source/file"
)

func Run(ctx context.Context) {
	db := env.FromContext(ctx).DB
	migrationDriver, err := migrateSql.WithInstance(db, &migrateSql.Config{})
	if err != nil {
		panic("error while creating migration driver |  err: " + err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance(os.Getenv("MIGRATION_SRC"), "mysql", migrationDriver)
	if err != nil {
		panic("error while creating migrate instance |  err: " + err.Error())
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		panic("error while running migrations |  err: " + err.Error())
	}
}
