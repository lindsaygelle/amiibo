package amiibo

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestNet(t *testing.T) {

	client := new(http.Client)

	req, err := http.NewRequest(http.MethodOptions, amiiboURL, nil)

	if ok := err == nil; !ok {
		t.Fatalf(fmt.Sprintf("%v", err))
	}

	resp, err := client.Do(req)

	if ok := err == nil; !ok {
		t.Fatalf(fmt.Sprintf("%v", err))
	}
	if ok := resp.StatusCode == http.StatusOK; !ok {
		t.Fatalf(fmt.Sprintf("%s is not responding %s", amiiboURL, resp.Status))
	}

	methodsPermitted := strings.Replace(resp.Header.Get("Allow"), ",", " ", -1)

	if ok := strings.Contains(methodsPermitted, http.MethodGet); ok {
		t.Fatalf(fmt.Sprintf("%s does not permit required HTTP method; only supports \"%v\"", amiiboURL, methodsPermitted))
	}
}
