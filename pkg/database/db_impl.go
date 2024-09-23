package database

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/projects/sys-des/txn-routine/pkg/env"
)

type sqlDb struct {
	db *sql.DB
}

func NewSqlDb(ctx context.Context) DB {
	_db := env.FromContext(ctx).DB
	return sqlDb{
		db: _db,
	}
}

func (s sqlDb) Begin() (*sql.Tx, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (s sqlDb) BeginTx(ctx context.Context, txOptions *sql.TxOptions) (*sql.Tx, error) {
	tx, err := s.db.BeginTx(ctx, txOptions)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (s sqlDb) Close() error {
	return s.db.Close()
}

func (s sqlDb) Exec(query string, args ...interface{}) (sql.Result, error) {
	return s.db.Exec(query, args...)
}

func (s sqlDb) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return s.db.ExecContext(ctx, query, args...)
}

func (s sqlDb) PingContext(ctx context.Context) error {
	return s.db.PingContext(ctx)
}

func (s sqlDb) Prepare(query string) (*sql.Stmt, error) {
	return s.db.Prepare(query)
}

func (s sqlDb) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return s.db.PrepareContext(ctx, query)
}

func (s sqlDb) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return s.db.Query(query, args...)
}

func (s sqlDb) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return s.db.QueryContext(ctx, query, args...)
}

func (s sqlDb) QueryRow(query string, args ...interface{}) *sql.Row {
	return s.db.QueryRow(query, args...)
}

func (s sqlDb) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return s.db.QueryRowContext(ctx, query, args...)
}

func (s sqlDb) SetConnMaxLifetime(dur time.Duration) {
	s.db.SetConnMaxLifetime(dur)
}

func (s sqlDb) SetMaxIdleConns(n int) {
	s.db.SetMaxIdleConns(n)
}

func (s sqlDb) SetMaxOpenConns(n int) {
	s.db.SetMaxOpenConns(n)
}

func (s sqlDb) Stats() sql.DBStats {
	return s.db.Stats()
}
