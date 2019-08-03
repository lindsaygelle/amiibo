package amiibo

import (
	"fmt"
	"testing"
)

func TestAmiiboMap(t *testing.T) {

	b, err := local()
	if err != nil {
		t.Fatalf(fmt.Sprintf("%v", err))
	}
	m := getAmiiboMap(b)
	if m == nil {
		t.Fatalf("{nil}")
	}

	fmt.Println(NewAmiiboMap(*b...))
}
