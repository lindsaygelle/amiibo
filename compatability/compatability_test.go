package compatability_test

import (
	"fmt"
	"testing"

	"github.com/gellel/amiibo/compatability"
)

func Test(t *testing.T) {
	const (
		templateErr string = "compatability.XHR: %s"
	)
	var (
		xhr, err = compatability.Get()
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if len(xhr.Games) == 0 {
		t.Fatalf(fmt.Sprintf(templateErr, "x.Games is empty"))
	}
	if len(xhr.Items) == 0 {
		t.Fatalf(fmt.Sprintf(templateErr, "x.Items is empty"))
	}
}
