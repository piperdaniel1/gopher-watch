package db

import "database/sql"

type DBInterface interface {
	QueryRow(query string, args ...interface{}) *sql.Row
	Query(query string, args ...any) (*sql.Rows, error)
	Ping() error
}
