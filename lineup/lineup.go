package lineup

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gellel/amiibo/file"
	"github.com/gellel/amiibo/network"
)

const (
	// Version is the semver of lineup.XHR.
	Version string = "1.0.0"
)

var (
	// Extension is the file extension lineup.XHR is written as.
	Extension string = "json"
)

var (
	// Name is the filename (before the .extension) lineup.XHR is written as.
	Name string = "lineup"
)

// Get performs a HTTP request to Nintendo Amiibo lineup resource and unmarshals the
// HTTP response body on http.StatusOK. Returns an error if the Nintendo server
// returns anything other than http.StatusOK. If the response content cannot be
// handled by json.Unmarshal the corresponding error message is returned. Get
// will always contact the Nintendo Amiibo lineup using the preconstructed
// lineup.Request. The lineup.Request can be modified to provide any additional
// parameters should the Nintendo endpoint change.
func Get() (*XHR, error) {
	var (
		data     []byte
		res, err = network.Client.Do(Request)
		xhr      XHR
	)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(res.Status)
	}
	defer res.Body.Close()
	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &xhr)
	if err != nil {
		return nil, err
	}
	xhr.ContentLength = res.ContentLength
	xhr.Cookies = res.Cookies()
	xhr.Headers = res.Header
	xhr.Status = res.Status
	xhr.StatusCode = res.StatusCode
	xhr.Version = Version
	return &xhr, err
}

// Load loads a written HTTP response from Nintendo Amiibo lineup resource
// and unmarshals the io content into a lineup.XHR. Returns an error
// if the fullpath does not point to a marshalled lineup.XHR or if a
// io read error occurs.
func Load(fullpath string) (*XHR, error) {
	var (
		b, err = file.Open(fullpath)
		x      XHR
	)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &x)
	if err != nil {
		return nil, err
	}
	return &x, err
}

// Write writes the HTTP response from Nintendo Amiibo lineup resource
// as a marshalled JSON file to the provided path. Writes content using the
// provided file permissions, but will always write using the lineup.Name
// and lineup.Extension. These values can be overwritten to suit the
// requirements of the integration.
// Returns an error if the file cannot be written to the
// target destination or a JSON marshalling error occurs.
func Write(path string, perm os.FileMode, xhr *XHR) (string, error) {
	var (
		b, err   = json.Marshal(xhr)
		fullpath string
	)
	if err != nil {
		return fullpath, err
	}
	return file.Make(path, Name, Extension, perm, b)
}
