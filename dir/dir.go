package dir

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	templateErr string = "%s is not a directory" // template error invalid directory
)

const (
	// Permission is the os.FileMode all directories created by the amiibo package are written under.
	Permission os.FileMode = 0777
)

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
	err = os.Remove(p)
	ok = (err == nil)
	if !ok {
		return err
	}
	return err
}

func DelAll(path string, folder string) error {
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
	err = os.RemoveAll(p)
	ok = (err == nil)
	if !ok {
		return err
	}
	return err
}

func Has(path string) bool {
	return (Not(path) == false)
}

func Not(path string) bool {
	var (
		err error
	)
	_, err = os.Stat(path)
	return os.IsNotExist(err)
}

func Current() (string, error) {
	var (
		err error
		ok  bool
		s   string
	)
	s, err = os.Executable()
	ok = (err == nil)
	if !ok {
		return s, err
	}
	s = filepath.Dir(s)
	return s, err
}

func Is(path string) bool {
	var (
		err  error
		info os.FileInfo
		ok   bool
	)
	info, err = os.Stat(path)
	ok = (err == nil)
	if !ok {
		return false
	}
	return info.IsDir()
}

func Make(path string, folder string) (string, error) {
	var (
		err error
		ok  bool
		p   = filepath.Join(path, folder)
	)
	ok = Has(p)
	if ok {
		return p, err
	}
	err = os.MkdirAll(p, Permission)
	ok = (err == nil)
	if !ok {
		return p, err
	}
	ok = Is(p)
	if !ok {
		return p, fmt.Errorf(templateErr, p)
	}
	return p, err
}
