package main

import (
	"net/http"
)

func NewAPI() API {
	return API{
		http.Client{
			CheckRedirect: func(request *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
			Timeout: httpTimeout}}
}

type API struct {
	http.Client
}
