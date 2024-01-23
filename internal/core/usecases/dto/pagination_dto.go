package dto

const DEFAULT_LIMIT = 100

type PageParams struct {
	offset int
	limit  int
}

func NewPageParams(offset, limit int) PageParams {
	return PageParams{
		offset: offset,
		limit:  limit,
	}
}

func (p PageParams) GetLimit() int {
	if p.limit < 1 || p.limit > DEFAULT_LIMIT {
		return DEFAULT_LIMIT
	}
	return p.limit
}

func (p PageParams) GetOffset() int {
	if p.offset < 0 {
		return 1
	}
	return p.offset
}

type Page[T any] struct {
	Result []T  `json:"results"`
	Next   *int `json:"next,omitempty"`
}

func BuildPage[T any](list []T, params PageParams) Page[T] {
	if len(list) > 0 && len(list) == params.limit {
		next := params.GetOffset() + params.GetLimit()
		return Page[T]{
			Result: list,
			Next:   &next,
		}
	}

	return Page[T]{
		Result: list,
	}
}
