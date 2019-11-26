package mix_test

import (
	"fmt"
	"testing"

	"github.com/gellel/amiibo/mix"
)

func Test(t *testing.T) {
	var (
		m, err = mix.Get()
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	fmt.Println(m)
}
