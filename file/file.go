package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gellel/amiibo/dir"
)

const (
	tep string = "%s.%s"
)

const (
	templateErr string = "%s is not a file"
)

func Del(fullpath string) error {
	var (
		ok = Is(fullpath)
	)
	if !ok {
		return fmt.Errorf(templateErr, fullpath)
	}
	return os.Remove(fullpath)
}

func Is(fullpath string) bool {
	var (
		info, err = os.Stat(fullpath)
		ok        bool
	)
	if os.IsNotExist(err) {
		return ok
	}
	return (info.IsDir() == ok)
}

func Make(path, name, ext string, perm os.FileMode, b []byte) (string, error) {
	var (
		fullpath string
		p, err   = dir.Make(path, perm)
	)
	if err != nil {
		return fullpath, err
	}
	fullpath = fmt.Sprintf(tep, filepath.Join(p, name), ext)
	err = ioutil.WriteFile(fullpath, b, perm)
	return fullpath, err
}
