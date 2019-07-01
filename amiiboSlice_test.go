package amiibo

import (
	"fmt"
	"testing"
)

func TestAmiiboSlice(t *testing.T) {

	b, err := local()
	if err != nil {
		t.Fatalf(fmt.Sprintf("%v", err))
	}
	s := getAmiiboSlice(b)
	if s == nil {
		t.Fatalf(fmt.Sprintf("{%v}", s))
	}

}
