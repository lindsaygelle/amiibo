package http_test

import (
	"fmt"
	"testing"

	"github.com/gellel/amiibo/http"
)

func TestHTTP(t *testing.T) {

	fmt.Println(http.NewHTTP().New("GET", "https://amiiboapi.com/"))
}
