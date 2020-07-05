package amiibo

import "testing"

func TestGetLineupComingJPN(t *testing.T) {

	var _, _, err = getLineupComingJPN()
	if err != nil {
		t.Fatal(err)
	}
}
