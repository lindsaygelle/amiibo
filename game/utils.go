package game

import (
	"encoding/json"
	"os"

	"github.com/gellel/amiibo/errors"
	"github.com/gellel/amiibo/file"
)

var (
	// Extension is the file extension game.Game is written as.
	Extension string = "json"
)

var (
	// Name is the filename key used (before the .extension) when writing game.Game using game.Write.
	Name string = "name"
)

// Write writes an game.Game to the provided path using the supported file permission.
//
// Write usess the Game.Field function to select the filename that the Game will be written under.
// If the provided field cannot be found in the Game, the function will
// return an error and not write the file.
// Upon successfully writing an game.Game, the fullpath that the struct was written as is
// returned and can be used to load the newly written content from.
func Write(path string, perm os.FileMode, game *Game) (string, error) {
	var (
		b        []byte
		err      error
		fullpath string
	)
	if game == nil {
		return fullpath, errors.ErrArgGameNil
	}
	var (
		name = game.Field(Name)
	)
	if len(name) == 0 {
		return fullpath, err
	}
	b, err = json.Marshal(game)
	if err != nil {
		return fullpath, err
	}
	fullpath, err = file.Make(path, name, Extension, perm, b)
	return fullpath, err
}
