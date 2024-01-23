package sql

import "database/sql"

type ResultWrapper interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type resultWrapper struct {
	result sql.Result
}

func NewResultWrapper(result sql.Result) ResultWrapper {
	return resultWrapper{
		result,
	}
}

func (r resultWrapper) LastInsertId() (int64, error) {
	return r.result.LastInsertId()
}

func (r resultWrapper) RowsAffected() (int64, error) {
	return r.result.RowsAffected()
}
