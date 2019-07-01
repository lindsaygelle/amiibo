package amiibo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	_ rawItem = (*RawItem)(nil)
)

func deleteRawItem(rawItem *RawItem) error {
	return os.Remove(filepath.Join(storepathRawItem(), fmt.Sprintf("r-%s.json", rawItem.Title)))
}

func getRawItem(ID string) *RawItem {
	if ok := strings.HasSuffix(ID, ".json"); !ok {
		ID = fmt.Sprintf("%s.json", ID)
	}
	if ok := strings.HasPrefix(ID, "r-"); !ok {
		ID = fmt.Sprintf("r-%s", ID)
	}
	b, err := openRawItem(ID)
	if err != nil {
		return nil
	}
	rawItem, err := unmarshallRawItem(b)
	if err != nil {
		return nil
	}
	return rawItem
}

func marshallRawItem(rawItem *RawItem) ([]byte, error) {
	content, err := json.Marshal(rawItem)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func openRawItem(name string) (*[]byte, error) {
	filepath := filepath.Join(storepathRawItem(), name)
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

func storepathRawItem() string {
	return filepath.Join(rootpath, "item")
}

func unmarshallRawItem(content *[]byte) (*RawItem, error) {
	r := &RawItem{}
	err := json.Unmarshal(*content, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func writeRawItem(rawItem *RawItem) error {
	err := os.MkdirAll(storepathRawItem(), 0644)
	if err != nil {
		return err
	}
	content, err := marshallRawItem(rawItem)
	if err != nil {
		return err
	}
	filepath := filepath.Join(storepathRawItem(), fmt.Sprintf("r-%s.json", rawItem.Title))
	return ioutil.WriteFile(filepath, content, 0644)
}

type rawItem interface {
	String() string
}

type RawItem struct {
	Description  string `json:"description"`  // null
	LastModified int64  `json:"lastModified"` // 1554418285473
	Path         string `json:"path"`         // "/content/noa/en_US/amiibo/detail/wolf-link-amiibo"
	Title        string `json:"title"`        // "Wolf Link"
	URL          string `json:"url"`          // "/amiibo/detail/wolf-link-amiibo"
}

func (pointer *RawItem) String() string {
	return fmt.Sprintf("%s", pointer.Title)
}
