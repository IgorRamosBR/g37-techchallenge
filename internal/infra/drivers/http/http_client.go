package http

import "net/http"

type HttpClient interface {
	DoPost(path string, body []byte) (*http.Response, error)
}
