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

// Del deletes a file at the filepath target if the file
// is a file and if os.Remove is permitted to remove the file.
func Del(fullpath string) error {
	var (
		ok = Is(fullpath)
	)
	if !ok {
		return fmt.Errorf(templateErr, fullpath)
	}
	return os.Remove(fullpath)
}

// Is checks that the filepath target is a file.
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

// Make makes a file at the filepath target, using the name and extension
// argument to create the file name. Requires the permissions to the target
// to be accessible by Go.
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

// Open opens a file at the filepath provided. Returns an error
// if the fullpath does not point to a file or
// if ioutil.ReadFile cannot read the target destination.
func Open(fullpath string) ([]byte, error) {
	if Is(fullpath) == false {
		return nil, fmt.Errorf(templateErr, fullpath)
	}
	return ioutil.ReadFile(fullpath)
}
