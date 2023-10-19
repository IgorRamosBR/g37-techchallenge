package clients

import (
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

func (c postgresClient) Find(entity interface{}, query string, values ...any) error {
	result := c.db.Where(query, values).First(entity)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (c postgresClient) FindAll(entity interface{}) error {
	result := c.db.Find(entity)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (c postgresClient) Save(entity interface{}) error {
	result := c.db.Create(entity)
	return result.Error
}
