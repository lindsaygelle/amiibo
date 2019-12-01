package write

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
	// Permission is the os.FileMode all content is written as using the amiibo package.
	Permission os.FileMode = 0777
)

func Current(name, ext string, b []byte) (string, error) {
	var (
		fullpath string
		p, err   = dir.Current()
	)
	if err != nil {
		return fullpath, err
	}
	fullpath = fmt.Sprintf(tep, filepath.Join(p, name), ext)
	err = ioutil.WriteFile(fullpath, b, Permission)
	return fullpath, err
}

func Write(directory, folder, name, ext string, b []byte) (string, error) {
	var (
		fullpath string
		p, err   = dir.Make(directory, folder)
	)
	if err != nil {
		return fullpath, err
	}
	fullpath = fmt.Sprintf(tep, filepath.Join(p, name), ext)
	err = ioutil.WriteFile(fullpath, b, Permission)
	return fullpath, err
}
