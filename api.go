package main

import (
	"net/http"
	"time"
)

func NewAPI() API {
	return API{
		http.Client{
			CheckRedirect: func(request *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
			Timeout: time.Second * 10}}
}

type API struct {
	http.Client
}

func (pointer API) HTTP() *http.Response {

	response, _ := pointer.Get(URL)
	return response
}
