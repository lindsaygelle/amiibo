package server

import (
	"encoding/json"
	"net/http"
)

// Flat is a marshalled set of structs ready as a json response.
type Flat []byte

func (f Flat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte(f))
}

// NewFlat returns a prepared http.Handler.
func NewFlat(v []interface{}) (http.Handler, error) {
	var (
		b, err = json.Marshal(v)
	)
	if err != nil {
		return nil, err
	}
	return Flat(b), err
}
