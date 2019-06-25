package amiibo_test

import (
	"reflect"
	"testing"

	"github.com/gellel/amiibo"
)

func TestAmiiboRaw(t *testing.T) {
	var (
		amiiboSeires = "amiiboSeries"
		character    = "character"
		gameSeries   = "gameSeries"
		head         = "head"
		image        = "image"
		name         = "name"
		tail         = "tail"
		types        = "type"
		AU           = "00-00-00"
		EU           = "00-00-00"
		JP           = "00-00-00"
		NA           = "00-00-00"
	)
	r := amiibo.NewRawAmiibo(amiiboSeires, character, gameSeries, head, image, name, tail, types, AU, EU, JP, NA)

	if ok := reflect.ValueOf(r).Kind() == reflect.Ptr; ok != true {
		t.Fatalf("reflect.ValueOf(amiibo.NewRawAmiibo(amiiboSeires, character, gameSeries, head, image, name, tail, t, AU, EU, JP, NA)) != reflect.Ptr")
	}
	if ok := r.String() == "headtail"; ok != true {
		t.Fatalf("rawAmiibo.String() did not concatenate head and tail")
	}
}
