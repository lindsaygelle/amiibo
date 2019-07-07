package amiibo

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNet(t *testing.T) {

	resp, err := http.Get(amiiboURL)
	if ok := err == nil; !ok {
		t.Fatalf(fmt.Sprintf("%v", err))
	}
	if ok := resp.StatusCode == 200; !ok {
		t.Fatalf(fmt.Sprintf("Amiibo URI is not responding %s", resp.Status))
	}
}
