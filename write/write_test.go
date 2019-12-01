package write_test

import (
	"fmt"
	"testing"

	"github.com/gellel/amiibo/write"
)

func Test(t *testing.T) {

	var (
		fullpath, err = write.Current("test", "txt", []byte("test"))
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	fmt.Println(fullpath)
}
