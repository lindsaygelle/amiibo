package mux

import (
	"net/http"

	"github.com/gorilla/mux"
)

var (
	ByteID string = "ID"
)

type Byte map[string][]byte

func (b Byte) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		vars = mux.Vars(r)
	)
	var (
		key = vars[ByteID]
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
