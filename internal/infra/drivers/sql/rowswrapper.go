package sql

import "database/sql"

type RowsWrapper interface {
	Next() bool
	NextResultSet() bool
	Columns() ([]string, error)
	ColumnTypes() ([]*sql.ColumnType, error)
	Scan(dest ...any) error
	Close() error
	Err() error
}

type rowsWrapper struct {
	rows *sql.Rows
}

func NewRowsWrapper(rows *sql.Rows) RowsWrapper {
	return rowsWrapper{
		rows: rows,
	}
}

func (r rowsWrapper) Next() bool {
	return r.rows.Next()
}

func (r rowsWrapper) NextResultSet() bool {
	return r.rows.NextResultSet()
}

func (r rowsWrapper) Columns() ([]string, error) {
	return r.rows.Columns()
}

func (r rowsWrapper) ColumnTypes() ([]*sql.ColumnType, error) {
	return r.rows.ColumnTypes()
}

func (r rowsWrapper) Scan(dest ...any) error {
	return r.rows.Scan(dest...)
}

func (r rowsWrapper) Close() error {
	return r.rows.Close()
}

func (r rowsWrapper) Err() error {
	return r.rows.Err()
}
