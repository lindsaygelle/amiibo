package dir

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	templateErr string = "%s is not a directory" // template error invalid directory
)

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

func Del(path string, folder string) error {
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

func DelAt(path string) error {
	var (
		err error
		ok  bool
	)
	ok = Not(path)
	if ok {
		return err
	}
	ok = Is(path)
	if !ok {
		return fmt.Errorf(templateErr, path)
	}
	return os.RemoveAll(path)
}

func Has(path string) bool {
	return (Not(path) == false)
}

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

func Not(path string) bool {
	var (
		err error
	)
	_, err = os.Stat(path)
	return os.IsNotExist(err)
}
