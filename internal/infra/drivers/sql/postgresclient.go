package sql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type sqlClient struct {
	db *sql.DB
}

func NewPostgresSQLClient(username, password, host, port, dbname string) (SQLClient, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return sqlClient{
		db,
	}, nil
}

func (client sqlClient) Find(query string, args ...any) (RowsWrapper, error) {
	rows, err := client.db.Query(query, args...)
	return NewRowsWrapper(rows), err
}

func (client sqlClient) FindOne(query string, args ...any) RowWrapper {
	return NewRowWrapper(client.db.QueryRow(query, args...))
}

func (client sqlClient) Exec(query string, args ...any) (ResultWrapper, error) {
	result, err := client.db.Exec(query, args...)
	return NewResultWrapper(result), err
}

func (client sqlClient) ExecWithReturn(query string, args ...any) RowWrapper {
	return NewRowWrapper(client.db.QueryRow(query, args...))
}

func (client sqlClient) Begin() (TransactionWrapper, error) {
	tx, err := client.db.Begin()
	return NewTransactionWrapper(tx), err
}

func (client sqlClient) Ping() error {
	err := client.db.Ping()
	return err
}

func (client sqlClient) GetConnection() *sql.DB {
	return client.db
}
