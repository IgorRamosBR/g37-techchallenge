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

func (c postgresClient) Save(entity interface{}) error {
	result := c.db.Create(entity)
	return result.Error
}

func (c postgresClient) Update(entity interface{}) error {
	result := c.db.Save(entity)
	return result.Error
}

func (c postgresClient) Delete(entity interface{}) error {
	result := c.db.Delete(entity)
	return result.Error
}

func (c postgresClient) Find(entity any, limit, offset int, query string, values ...any) error {
	result := c.db.Scopes(paginate(limit, offset, c.db)).Where(query, values).Find(entity)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (c postgresClient) FindFirst(entity any, query string, values ...any) error {
	result := c.db.Where(query, values).First(entity)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (c postgresClient) FindById(id int, entity interface{}) error {
	result := c.db.First(entity, id)
	if result.RowsAffected == 0 {
		return ErrNotFound
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (c postgresClient) FindAll(entity any, limit, offset int, eagerFields string) error {
	scope := c.db.Scopes(paginate(limit, offset, c.db))
	if eagerFields != "" {
		scope = scope.Preload(eagerFields)
	}
	result := scope.Find(entity)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
