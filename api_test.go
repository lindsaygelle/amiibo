package amiibo_test

import (
	"testing"

	"github.com/gellel/amiibo"
)

func TestAPI(t *testing.T) {

	api := amiibo.NewAPI()

	err := api.Get()

	if err != nil {
		t.Fatalf("t.Fatal: amiibo.API.Get() returned err")
	}
}
