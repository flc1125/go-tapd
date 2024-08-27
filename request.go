package tapd

import (
	"net/http"
)

type RequestOption func(*http.Request) error

func WithRequestBasicAuth(username, password string) RequestOption {
	return func(req *http.Request) error {
		req.SetBasicAuth(username, password)
		return nil
	}
}

func WithRequestHeader(name, value string) RequestOption {
	return func(req *http.Request) error {
		req.Header.Set(name, value)
		return nil
	}
}

func WithRequestHeaders(headers map[string]string) RequestOption {
	return func(req *http.Request) error {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
		return nil
	}
}
