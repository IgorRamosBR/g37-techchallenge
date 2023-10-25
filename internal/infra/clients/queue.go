package clients

type Queue interface {
	SendMessage(data []byte) error
}
