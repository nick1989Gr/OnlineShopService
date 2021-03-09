package database

import (
	"context"
	"database/sql"
)

// IEncapsulatedSQLDb wraps arount sql.db methods so that we can easily mock them
type IEncapsulatedSQLDb interface {
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) 
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

// sqlDbImpl implements the EncapsulatedSQLDb interface
type sqlDbImpl struct {
	externalSQLDb *sql.DB
 }

 // QueryRowContext just calls the QueryRowContext from the external library
 func (s *sqlDbImpl) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return s.externalSQLDb.QueryRowContext(ctx, query, args ...)
 }

 // QueryContext just calls the QueryContext from the external library
 func (s *sqlDbImpl) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return s.externalSQLDb.QueryContext(ctx, query, args ...)
 }

 // ExecContext just calls the ExecContext from the external library
 func (s *sqlDbImpl) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return s.externalSQLDb.ExecContext(ctx, query, args ...)
 }

 // NewExternalSQLDb returns a new instance of SQLDbImpl
 func NewExternalSQLDb(db *sql.DB) IEncapsulatedSQLDb {
	return &sqlDbImpl{
		externalSQLDb: db,
	}
 }