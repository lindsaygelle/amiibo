package amiibo

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// Local attempts to read a previously cached Nintendo Amiibo XHR HTTP response.
func local() (*[]byte, error) {
	filepath := filepath.Join(rootpath, localFile)
	reader, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(reader)
	defer reader.Close()
	if err != nil {
		return nil, err
	}
	return &content, nil
}
