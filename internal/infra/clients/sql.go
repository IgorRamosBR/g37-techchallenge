package clients

import "errors"

var ErrNotFound = errors.New("entity not found")

type SQLClient interface {
	Save(entity interface{}) error
	Delete(entity interface{}) error
	Find(entity interface{}, query string, values ...any) error
	FindById(id uint, entity interface{}) error
	FindAll(entity interface{}) error
}
