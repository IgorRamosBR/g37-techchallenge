package sql

import "database/sql"

type TransactionWrapper interface {
	Find(query string, args ...any) (RowsWrapper, error)
	FindOne(query string, args ...any) RowWrapper
	Exec(query string, args ...any) (sql.Result, error)
	ExecWithReturn(query string, args ...any) RowWrapper
	Commit() error
	Rollback() error
}

type transactionWrapper struct {
	tx *sql.Tx
}

func NewTransactionWrapper(tx *sql.Tx) TransactionWrapper {
	return transactionWrapper{
		tx,
	}
}

func (t transactionWrapper) Find(query string, args ...any) (RowsWrapper, error) {
	rows, err := t.tx.Query(query, args...)
	return NewRowsWrapper(rows), err
}

func (t transactionWrapper) FindOne(query string, args ...any) RowWrapper {
	row := t.tx.QueryRow(query, args...)
	return NewRowWrapper(row)
}

func (t transactionWrapper) Exec(query string, args ...any) (sql.Result, error) {
	result, err := t.tx.Exec(query, args...)
	return result, err
}

func (t transactionWrapper) ExecWithReturn(query string, args ...any) RowWrapper {
	row := t.tx.QueryRow(query, args...)
	return NewRowWrapper(row)
}

func (t transactionWrapper) Commit() error {
	err := t.tx.Commit()
	return err
}

func (t transactionWrapper) Rollback() error {
	err := t.tx.Rollback()
	return err
}
