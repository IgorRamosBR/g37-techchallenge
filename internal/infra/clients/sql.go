package clients

import (
	"errors"

	"gorm.io/gorm"
)

var ErrNotFound = errors.New("entity not found")

type SQLClient interface {
	Save(entity any) error
	Delete(entity any) error
	Find(entity any, limit, offset int, query string, values ...any) error
	FindFirst(entity any, query string, values ...any) error
	FindById(id string, entity any) error
	FindAll(entity any, limit, offset int) error
}

type Pagination[T any] struct {
	Limit      int    `json:"limit,omitempty" query:"limit"`
	Page       int    `json:"page,omitempty" query:"page"`
	Sort       string `json:"sort,omitempty" query:"sort"`
	TotalItems int64  `json:"total_items"`
	TotalPages int    `json:"total_pages"`
	Items      []T    `json:"items"`
}

func NewPagination[T any](limit, page int, sort string) Pagination[T] {
	return Pagination[T]{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}
}

func (p Pagination[T]) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p Pagination[T]) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p Pagination[T]) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p Pagination[T]) GetSort() string {
	if p.Sort == "" {
		p.Sort = "Id desc"
	}
	return p.Sort
}

func paginate(limit, offset int, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(limit)
	}
}
