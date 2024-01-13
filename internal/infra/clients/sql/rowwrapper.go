package sql

import "database/sql"

type RowWrapper interface {
	Scan(dest ...any) error
	Err() error
}

type rowWrapper struct {
	row *sql.Row
}

func NewRowWrapper(row *sql.Row) RowWrapper {
	return rowWrapper{
		row: row,
	}
}

func (r rowWrapper) Scan(dest ...any) error {
	return r.row.Scan(dest...)
}

func (r rowWrapper) Err() error {
	return r.row.Err()
}
