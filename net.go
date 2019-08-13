package amiibo

import (
	"errors"
	"io/ioutil"
	"net/http"
)

// Net performs a HTTP GET to the Nintendo Amiibo URL. Returns the HTTP response body after HTTP status is OK. Returns nil if a HTTP or IO error occurs.
func net() (*[]byte, error) {
	req, err := http.NewRequest(http.MethodGet, amiiboURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &body, nil
}
