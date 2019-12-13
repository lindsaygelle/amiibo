package game

import (
	"encoding/json"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/mux"

	"github.com/gellel/amiibo/errors"
	"github.com/gellel/amiibo/file"
	"github.com/gellel/amiibo/mix"
)

var (
	// Extension is the file extension game.Game is written as.
	Extension string = "json"
)

var (
	// Name is the filename key used (before the .extension) when writing game.Game using game.Write.
	Name string = "name"
)

// Get gets all the HTTP content from the various Nintendo resources and organizes
// the response content into a slice of normalized game.Game.
//
// Get consumes the reponse content from the compatability.XHR and lineup.XHR.
// Should either one of these resources not be able to be parsed,
// the collection will return an error.
func Get() ([]*Game, error) {
	var (
		m, err = mix.Get()
	)
	if err != nil {
		return nil, err
	}
	var (
		s  = []*Game{}
		wg sync.WaitGroup
	)
	for _, m := range m.Games {
		wg.Add(1)
		go func(m *mix.Game) {
			defer wg.Done()
			var (
				a, err = NewGame(m.Game, m.Item)
			)
			if err != nil {
				return
			}
			s = append(s, a)
		}(m)
	}
	wg.Wait()
	return s, nil
}

// Load loads an game.Game from the provided fullpath using the last substring after the
// trailing slash as the file name to open.
//
// Load assumes the fullpath points to a valid json file. If the function cannot parse
// or cannot reach the file, a corresponding error is returned.
func Load(fullpath string) (*Game, error) {
	var (
		game   Game
		b, err = file.Open(fullpath)
	)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &game)
	if err != nil {
		return nil, err
	}
	return &game, err
}

// Server sets up a basic http.Handler interface to be used with http.ListenAndServe.
//
// Server is built on the map strategy used when creating the argument game.Map
// passed into the function. Each game.Game in the game.Map is returned
// when a HTTP connection matches the provided map key path used when
// creating the initial game.Map.
//
// When the key cannot be used to match to a value in the game.Map
// the HTTP request is treated as http.StatusNotFound.
//
// If the game.Game cannot be marshaled, the HTTP request is
// handle as http.StatusServiceUnavailable.
//
// Server only accepts connections via http.MethodGet.
func Server(m *Map) http.Handler {
	const (
		contentTypeKey string = "Content-Type"
	)
	const (
		contentTypeValue string = "application/json; charset=utf-8"
	)
	var r = mux.NewRouter().StrictSlash(true)
	var handleSlash = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(contentTypeKey, contentTypeValue)
		var b, err = json.Marshal(m)
		if err == nil {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
		w.Write(b)
	}
	var handleID = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(contentTypeKey, contentTypeValue)
		var vars = mux.Vars(r)
		var ID = vars["ID"]
		var (
			amiibo, ok = m.Get(ID)
		)
		if ok {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
		var b, err = json.Marshal(amiibo)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
		w.Write(b)
	}
	r.HandleFunc("/", handleSlash).Methods(http.MethodGet)
	r.HandleFunc("/{ID}", handleID).Methods(http.MethodGet)
	return r
}

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
