package amiibo

import (
	"encoding/json"
	"os"

	"github.com/gellel/amiibo/errors"
	"github.com/gellel/amiibo/file"
)

// Write writes an amiibo.Amiibo to the provided path using the supported file permission.
//
// Write usess the Amiibo.Field function to select the filename that the Amiibo will be written under.
// If the provided field cannot be found in the Amiibo, the function will
// return an error and not write the file.
// Upon successfully writing an amiibo.Amiibo, the fullpath that the struct was written as is
// returned and can be used to load the newly written content from.
// All amiibo.Amiibo are written as json.
func Write(path, key string, perm os.FileMode, amiibo *Amiibo) (string, error) {
	const (
		ext string = "json"
	)
	var (
		b        []byte
		err      error
		fullpath string
	)
	if amiibo == nil {
		return fullpath, errors.ErrArgAmiiboNil
	}
	var (
		name = amiibo.Field(key)
	)
	if len(name) == 0 {
		return fullpath, err
	}
	b, err = json.Marshal(amiibo)
	if err != nil {
		return fullpath, err
	}
	fullpath, err = file.Make(path, name, ext, perm, b)
	return fullpath, err
}

// WriteAll writes a collection of amiibo.Amiibo to the provided path using the supported file permissions.
//
// On unsuccessful writes the function will exit and return an error.
func WriteAll(path, key string, perm os.FileMode, amiibo ...*Amiibo) {
	var (
		err error
	)
	for _, amiibo := range amiibo {
		_, err = Write(path, key, per, amiibo)
		if err != nil {
			return err
		}
	}
	return err
}
