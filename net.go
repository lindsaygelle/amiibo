package amiibo

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func net() (*[]byte, error) {
	req, err := http.NewRequest("GET", amiiboURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &body, nil
}
