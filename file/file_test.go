package file_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gellel/amiibo/file"
)

func Test(t *testing.T) {
	const (
		ext        string      = "txt"
		filename   string      = "test"
		permission os.FileMode = 0777
	)
	var (
		current  string
		err      error
		fullpath string
		s        string
	)
	s, err = os.Executable()
	if err != nil {
		t.Fatalf(err.Error())
	}
	current = filepath.Dir(s)
	fullpath, err = file.Make(current, filename, ext, permission, []byte("test"))
	if err != nil {
		t.Fatalf(err.Error())
	}
	err = file.Del(fullpath)
	if err != nil {
		t.Fatalf(err.Error())
	}
}
