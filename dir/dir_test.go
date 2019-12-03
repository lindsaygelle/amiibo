package dir_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gellel/amiibo/dir"
)

func Test(t *testing.T) {
	var (
		current  string
		err      error
		fullpath string
	)
	current, err = dir.Current()
	if err != nil {
		t.Fatalf(err.Error())
	}
	fullpath, err = dir.MakeAt(current, fmt.Sprintf("%d", time.Now().Unix()), 0777)
	if err != nil {
		t.Fatalf(err.Error())
	}
	err = dir.Del(fullpath)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if dir.Not(fullpath) == false {
		t.Fatalf("dir.DelAt did not remove %s", fullpath)
	}
}
