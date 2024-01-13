package sql

import "errors"

var ErrNotFound = errors.New("entity not found")

type SQLClient interface {
	Find(query string, args ...any) (RowsWrapper, error)
	FindOne(query string, args ...any) RowWrapper
	Exec(query string, args ...any) (ResultWrapper, error)
	Begin() (TransactionWrapper, error)
	Ping() error
}
