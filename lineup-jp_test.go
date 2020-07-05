package amiibo

import (
	"fmt"
	"testing"
)

func TestGetLineupJPN(t *testing.T) {

	var _, _, err = getLineupJPN()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetLineupJPNXML(t *testing.T) {

	var v, err = getLineupJPNXML()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(v)
}
