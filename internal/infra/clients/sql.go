package clients

type SQLClient interface {
	Save(entity interface{}) error
	Find(entity interface{}, query string, values ...any) error
	FindAll(entity interface{}) error
}
