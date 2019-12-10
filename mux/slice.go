package mux

import "net/http"

var (
	// SliceID is the default namespace Slice uses with router.Handler(path, http.Handler).
	SliceID string = ""
)

// Slice is a slice-like struct of bytes that implements a http.Handler.
type Slice []byte

// ServeHTTP handles incoming HTTP requests in compliance with http.Handler.
//
// ServeHTTP assumes the byte sequence is a RFC compliant JSON string.
func (s Slice) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set(contentType, contentTypeValue)
	w.Write(s)
}
