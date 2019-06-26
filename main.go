package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getNintendoXHR() (*RawPayload, error) {
	req, err := http.NewRequest("GET", URL, nil)
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
	r := &RawPayload{}
	err = json.Unmarshal(body, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func main() {
	r, err := getNintendoXHR()
	if err != nil {
		panic(err)
	}
	r.AmiiboList.Each(func(i int, r *RawAmiibo) {
		fmt.Println(r.BoxArtURL.String())
	})
}
