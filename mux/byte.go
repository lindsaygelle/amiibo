package mux

import (
	"net/http"

	"github.com/gorilla/mux"
)

const ()

type Byte map[string][]byte

func (b Byte) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	const (
		ID string = "ID"
	)
	var (
		vars = mux.Vars(r)
	)
	var (
		key = vars[ID]
	)
	var (
		c, ok = (b[key])
	)
	switch ok {
	case true:
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
	w.Write(c)
}
