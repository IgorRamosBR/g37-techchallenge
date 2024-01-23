package http

import (
	"bytes"
	"io"
	httpClient "net/http"
)

type mockHttpClient struct {
}

func NewHttpClient() HttpClient {
	return mockHttpClient{}
}

func (c mockHttpClient) DoPost(path string, body []byte) (*httpClient.Response, error) {
	response := httpClient.Response{
		Body: io.NopCloser(bytes.NewBufferString(`{"qr_data":"00020101021243650016COM.MERCADOLIBRE02013063638f1192a-5fd1-4180-a180-8bcae3556bc35204000053039865802BR5925IZABELAAAADEMELO6007BARUERI62070503***63040B6D","in_store_order_id":"d4e8ca59-3e1d-4c03-b1f6-580e87c654ae"}`)),
	}

	return &response, nil
}
