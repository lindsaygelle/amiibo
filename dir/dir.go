package dir

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	templateErr string = "%s is not a directory" // template error invalid directory
)

// Current gets the current folder the program is being executed from.
func Current() (string, error) {
	var (
		err error
		s   string
	)
	s, err = os.Executable()
	if err != nil {
		return s, err
	}
	s = filepath.Dir(s)
	return s, err
}

// Del deletes a folder at the filepath using os.RemoveAll. Uses the
// last entry in the filepath as the folder to be removed. Returns
// an error if path does not point to a folder or if os.RemoveAll
// cannot remove the folder.
func Del(path string) error {
	var (
		ok = Is(path)
	)
	if !ok {
		return fmt.Errorf(templateErr, path)
	}
	return os.RemoveAll(path)
}

// DelAt deletes a folder from the target path by the argument name. Returns
// an error if os.RemoveAll cannot remove the folder or if the path or name
// does not point to a folder.
func DelAt(path string, folder string) error {
	var (
		err error
		ok  bool
		p   = filepath.Join(path, folder)
	)
	ok = Not(p)
	if ok {
		return err
	}
	ok = Is(p)
	if !ok {
		return fmt.Errorf(templateErr, p)
	}
	return os.RemoveAll(p)
}

// Has checks if the path is a folder.
func Has(path string) bool {
	return (Not(path) == false)
}

// Is checks if the path points to a folder.
func Is(path string) bool {
	var (
		err  error
		info os.FileInfo
		ok   bool
	)
	info, err = os.Stat(path)
	if err != nil {
		return ok
	}
	ok = info.IsDir()
	return ok
}

// Make makes a new folder at the filepath using os.MkDirAll. Uses the
// last entry in the filepath as the folder name. Does not create a new folder
// if the folder already exists. Returns an error if os.MkDirAll
// cannot create a new folder.
func Make(path string, perm os.FileMode) (string, error) {
	var (
		err error
		ok  = Has(path)
	)
	if ok {
		return path, err
	}
	err = os.MkdirAll(path, perm)
	return path, err
}

// MakeAt makes a new folder at the provided filepath, using the name argument
// as the new folder to be added to the directory. Does not create a new folder if the folder
// already exists. Returns an error if os.MkDirAll
// cannot create a new folder.
func MakeAt(path string, folder string, perm os.FileMode) (string, error) {
	var (
		err error
		ok  bool
		p   = filepath.Join(path, folder)
	)
	ok = Has(p)
	if ok {
		return p, err
	}
	err = os.MkdirAll(p, perm)
	return p, err
}

// Not checks that the path does not point to a folder.
func Not(path string) bool {
	var (
		err error
	)
	_, err = os.Stat(path)
	return os.IsNotExist(err)
}
