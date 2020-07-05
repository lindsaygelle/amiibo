package amiibo

import (
	"fmt"
	"testing"
)

func TestGetLineupComingJPN(t *testing.T) {

	var _, _, err = getLineupComingJPN()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetLineupComingJPNXML(t *testing.T) {

	var v, err = getLineupComingJPNXML()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(v)
}
