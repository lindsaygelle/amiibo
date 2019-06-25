package amiibo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetAmiiboMap(URI ...string) (*Map, error) {
	r, err := GetRawAmiibo(URI...)
	if err != nil {
		return nil, err
	}
	return NewMapFromRaw(r.Amiibo), nil
}

func GetAmiiboSlice(URI ...string) (*Slice, error) {
	r, err := GetRawAmiibo(URI...)
	if err != nil {
		return nil, err
	}
	return NewSliceFromRaw(r.Amiibo), nil
}

func GetAmiiboSet(URI ...string) (*Set, error) {
	r, err := GetRawAmiibo(URI...)
	if err != nil {
		return nil, err
	}
	return NewSetFromRaw(r.Amiibo), nil
}

func GetRawAmiibo(URI ...string) (*RawResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s", URL), nil)
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
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var rawResponse RawResponse
	err = json.Unmarshal(body, &rawResponse)
	if err != nil {
		return nil, err
	}
	return &rawResponse, nil
}
