package clients

import (
	"database/sql"

	"gorm.io/gorm"
)

type postgresClient struct {
	db *gorm.DB
}

type PostgresConfig struct {
	Username string
	Password string
	Engine   string
	Host     string
	Port     string
	Dbname   string
}

func NewPostgresClient(db *gorm.DB) SQLClient {
	return postgresClient{
		db: db,
	}
}

func (c postgresClient) FindAll(entity interface{}) (*sql.Rows, error) {
	result := c.db.Find(entity)
	if result.Error != nil {
		return nil, result.Error
	}

	rows, err := result.Rows()
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (c postgresClient) Save(entity interface{}) error {
	result := c.db.Create(entity)
	return result.Error
}
