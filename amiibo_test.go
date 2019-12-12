package amiibo_test

import (
	"net/http"
	"testing"

	"github.com/gellel/amiibo/amiibo"
)

func Test(t *testing.T) {
	a, err := amiibo.Get()
	if err != nil {
		panic(err)
	}
	m, err := amiibo.NewMap("Name", a...)
	if err != nil {
		panic(err)
	}
	http.ListenAndServe(":80", amiibo.Server(m))
}
