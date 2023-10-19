package clients

import "database/sql"

type SQLClient interface {
	Save(entity interface{}) error
	FindAll(entity interface{}) (*sql.Rows, error)
}
