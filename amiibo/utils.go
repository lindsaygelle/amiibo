package amiibo

import (
	"encoding/json"
	"os"

	"github.com/gellel/amiibo/errors"
	"github.com/gellel/amiibo/file"
)

var (
	// Extension is the file extension amiibo.Amiibo is written as.
	Extension string = "json"
)

var (
	// Name is the filename key used (before the .extension) when writing amiibo.Amiibo using amiibo.Write.
	Name string = "name"
)

// Load loads an amiibo.Amiibo from the provided fullpath using the last substring after the
// trailing slash as the file name to open.
//
// Load assumes the fullpath points to a valid json file. If the function cannot parse
// or cannot reach the file, a corresponding error is returned.
func Load(fullpath string) (*Amiibo, error) {
	var (
		amiibo Amiibo
		b, err = file.Open(fullpath)
	)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &amiibo)
	if err != nil {
		return nil, err
	}
	return &amiibo, err
}

// Write writes an amiibo.Amiibo to the provided path using the supported file permission.
//
// Write usess the Amiibo.Field function to select the filename that the Amiibo will be written under.
// If the provided field cannot be found in the Amiibo, the function will
// return an error and not write the file.
// Upon successfully writing an amiibo.Amiibo, the fullpath that the struct was written as is
// returned and can be used to load the newly written content from.
func Write(path string, perm os.FileMode, amiibo *Amiibo) (string, error) {
	var (
		b        []byte
		err      error
		fullpath string
	)
	if amiibo == nil {
		return fullpath, errors.ErrArgAmiiboNil
	}
	var (
		name = amiibo.Field(Name)
	)
	if len(name) == 0 {
		return fullpath, err
	}
	b, err = json.Marshal(amiibo)
	if err != nil {
		return fullpath, err
	}
	fullpath, err = file.Make(path, name, Extension, perm, b)
	return fullpath, err
}
