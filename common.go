package amiibo

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"path/filepath"
)

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
func readJSONFile(dir, filename string, v interface{}) error {
	return unmarshal(dir, filename, &v, json.Unmarshal)
}

// readXMLFile reads a XML file from disc and unmarshals its contents using xml.Unmarshal.
func readXMLFile(dir, filename string, v interface{}) error {
	return unmarshal(dir, filename, &v, xml.Unmarshal)
}

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

func writeFile(dir, filename string, b []byte) (fullpath string, err error) {
	fullpath = filepath.Join(dir, filename)
	err = ioutil.WriteFile(fullpath, b, 0644)
	return
}

func writeJSONFile(dir, filename string, v interface{}) (fullpath string, err error) {
	var b ([]byte)
	b, err = marshal(&v, json.Marshal)
	if err != nil {
		return
	}
	fullpath, err = writeFile(dir, filename, b)
	return
}

func writeXMLFile(dir, filename string, v interface{}) (fullpath string, err error) {
	var b ([]byte)
	b, err = marshal(&v, xml.Marshal)
	if err != nil {
		return
	}
	fullpath, err = writeFile(dir, filename, b)
	return
}
