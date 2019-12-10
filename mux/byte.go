package mux

import (
	"net/http"

	"github.com/gorilla/mux"
)

var (
	// MapID is the default namespace Map uses with router.Handler(path, http.Handler).
	MapID string = "ID"
)

// Map is a map-like struct of bytes that implements a http.Handler.
type Map map[string][]byte

// ServeHTTP handles incoming HTTP requests in compliance with http.Handler.
//
// ServeHTTP will attempt to fetch a byte set from the map
// to respond to the connecting client with the appropriate byte data.
// On successful map lookups the Map will write to the
// http.ResponseWriter http.StatusOK along with the found byte sequence.
//
// ServeHTTP will respond http.StatusNotFound if the map cannot find
// a corresponding key in the dataset.
//
// The ServeHTTP function uses the mux.Vars parser to
// extract the map key to retrieve a byte set when implemented by
// an interface that handles router.Handle("/{ID}", http.Handler)
//
// To change the router.Handle set mux.MapID to an appropriate key.
//
// Map assumes each byte set in the map is a json sequence.
func (m Map) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var ()
	var (
		vars = mux.Vars(r)
	)
	var (
		key = vars[MapID]
	)
	var (
		c, ok = (m[key])
	)
	switch ok {
	case true:
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
	w.Header().Set(contentType, contentTypeValue)
	w.Write(c)
}
