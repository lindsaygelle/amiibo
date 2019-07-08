package amiibo

import (
	"testing"
)

func TestAmiiboSlice(t *testing.T) {
	b, err := local()
	if err != nil {
		panic(err)
	}
	slice := getAmiiboSlice(b)
	if ok := slice.Len() > 0; !ok {
		t.Fatalf("amiibo.getAmiiboSlice(content []byte) returned an empty slice")
	}
}
