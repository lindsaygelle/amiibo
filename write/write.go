package write

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gellel/amiibo/dir"
)

const (
	tep string = "%s.%s" // tep is the template for formatting filename.extension.
)

const (
	// Permission is the os.FileMode all content is written as using the amiibo package.
	Permission os.FileMode = 0777
)

// Current writes the byte content to the current directory the program calling
// the amiibo package is being exected from. Translate the byte sequence into a corresponding
// file type based on the extension string provided to the function.
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

// Write writes the byte content to the specified directory using the argument
// byte sequence as the content to be sent to the file writer.
// Translate the byte sequence into a corresponding
// file type based on the extension string provided to the function.
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
