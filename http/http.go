package http

import "net/http"

func NewHTTP() *HTTP {
	return &HTTP{http: &http.Client{}}
}

type HTTP struct {
	http *http.Client
}

func (pointer *HTTP) New(r, URL string) (*http.Request, error) {

	req, err := http.NewRequest(r, URL, nil)

	HTTPHeaders.Each(func(key, value string) {
		req.Header.Add(key, value)
	})
	return req, err
}
