package amiibo

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
	// Extension is the file extension amiibo.Amiibo is written as.
	Extension string = "json"
)

var (
	// Name is the filename key used (before the .extension) when writing amiibo.Amiibo using amiibo.Write.
	Name string = "name"
)

func Get() ([]*Amiibo, error) {
	var (
		m, err = mix.Get()
	)
	if err != nil {
		return nil, err
	}
	var (
		s  = []*Amiibo{}
		wg sync.WaitGroup
	)
	for _, m := range m.Amiibo {
		wg.Add(1)
		go func(m *mix.Amiibo) {
			defer wg.Done()
			var (
				a, err = NewAmiibo(m.Compatability, m.Item, m.Lineup)
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

// Server sets up a basic http.Handler interface to be used with http.ListenAndServe.
//
// Server is built on the map strategy used when creating the argument amiibo.Map
// passed into the function.
func Server(m *Map) http.Handler {
	const (
		contentTypeKey = "Content-Type"
	)
	const (
		contentTypeValue = "application/json; charset=utf-8"
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
