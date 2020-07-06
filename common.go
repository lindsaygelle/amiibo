package amiibo

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

// getRemoteFile gets a remote file from a remote URL.
func getRemoteFile(URL string) (req *http.Request, res *http.Response, err error) {
	req, err = http.NewRequest(http.MethodGet, JPNChartURL, nil)
	if err != nil {
		return
	}
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf(("http: %d"), res.StatusCode)
	}
	if err != nil {
		return
	}
	return
}

// marshal handles a (package).Marshal operation.
func marshal(v interface{}, fn func(interface{}) ([]byte, error)) (b []byte, err error) {
	b, err = fn(v)
	return
}

// readFile handles a (package).ReadFile operation.
func readFile(dir, filename string, fn func(string) ([]byte, error)) (b []byte, err error) {
	b, err = fn(filepath.Join(dir, filename))
	return
}

// readJSONFile reads a JSON file from disc and unmarshals its contents using json.Unmarshal.
func readJSONFile(dir, filename string, v interface{}) (err error) {
	err = unmarshal(dir, filename, &v, json.Unmarshal)
	return
}

// readXMLFile reads a XML file from disc and unmarshals its contents using xml.Unmarshal.
func readXMLFile(dir, filename string, v interface{}) (err error) {
	err = unmarshal(dir, filename, &v, xml.Unmarshal)
	return
}

// unmarshal handles a (package).Unmarshal operation.
func unmarshal(dir, filename string, v interface{}, fn func([]byte, interface{}) error) (err error) {
	var b ([]byte)
	b, err = readFile(dir, filename, ioutil.ReadFile)
	if err != nil {
		return
	}
	err = fn(b, v)
	if err != nil {
		return
	}
	return
}

// writeFile writes a file to disc using ioutil.WriteFile.
func writeFile(dir, filename string, b []byte) (fullpath string, err error) {
	fullpath = filepath.Join(dir, filename)
	err = ioutil.WriteFile(fullpath, b, 0644)
	return
}

// writeFileJSON writes a JSON file to disc.
func writeJSONFile(dir, filename string, v interface{}) (fullpath string, err error) {
	var b ([]byte)
	b, err = marshal(&v, json.Marshal)
	if err != nil {
		return
	}
	fullpath, err = writeFile(dir, filename, b)
	return
}

// writeXMLFile writes a XML fille to disc.
func writeXMLFile(dir, filename string, v interface{}) (fullpath string, err error) {
	var b ([]byte)
	b, err = marshal(&v, xml.Marshal)
	if err != nil {
		return
	}
	fullpath, err = writeFile(dir, filename, b)
	return
}
