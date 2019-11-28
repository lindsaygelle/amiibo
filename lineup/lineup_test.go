package lineup_test

import (
	"fmt"
	"testing"

	"github.com/gellel/amiibo/lineup"
)

func TestXHR(t *testing.T) {
	const (
		templateErr string = "lineup.XHR: %s"
	)
	var (
		xhr, err = lineup.Get()
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if len(xhr.Amiibo) == 0 {
		t.Fatalf(fmt.Sprintf(templateErr, "x.Amiibo is empty"))
	}
	if len(xhr.Items) == 0 {
		t.Fatalf(fmt.Sprintf(templateErr, "x.Items is empty"))
	}
}
