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

// DeleteRawItem deletes the raw Item from the operating system if it is writtens. Returns an error if the raw Item is unable to be deleted or another file system issue occurs.
func DeleteRawItem(rawItem *RawItem) error {
	return os.Remove(filepath.Join(StorepathRawItem(), fmt.Sprintf("r-%s.json", rawItem.Title)))
}

// GetRawItem unmarshalls a raw Item struct from the operating system if it written to the disc. Returns nil if no corresponding raw Item is found or a unmarshalling error occurs.
func GetRawItem(ID string) *RawItem {
	if ok := strings.HasSuffix(ID, ".json"); !ok {
		ID = fmt.Sprintf("%s.json", ID)
	}
	if ok := strings.HasPrefix(ID, "r-"); !ok {
		ID = fmt.Sprintf("r-%s", ID)
	}
	b, err := OpenRawItem(ID)
	if err != nil {
		return nil
	}
	rawItem, err := UnmarshallRawItem(b)
	if err != nil {
		return nil
	}
	return rawItem
}

// MarshallRawItem marshalls a raw Item pointer into a byte slice and returns the byte slice value.
func MarshallRawItem(rawItem *RawItem) ([]byte, error) {
	content, err := json.Marshal(rawItem)
	if err != nil {
		return nil, err
	}
	return content, nil
}

// NewRawItem instantiates a new raw Item struct pointer from a valid byte slice pointer. Assumes the byte slice contains valid Raw Item byte indexes.
func NewRawItem(b *[]byte) *RawItem {
	r := new(RawItem)
	err := json.Unmarshal(*b, r)
	if err != nil {
		panic(err)
	}
	return r
}

// NewRawItemFromRawMessage instantiates a new raw Item pointer from a valid JSON raw message.
func NewRawItemFromRawMessage(r *json.RawMessage) *RawItem {
	rawItem := &RawItem{}
	err := json.Unmarshal(*r, rawItem)
	if err != nil {
		return nil
	}
	return rawItem
}

// OpenRawItem returns the byte pointer for a written raw Item struct by its storage name.
func OpenRawItem(name string) (*[]byte, error) {
	filepath := filepath.Join(StorepathRawItem(), name)
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

// StorepathRawItem returns the default absolute path for raw Item struct being written to the operating system.
func StorepathRawItem() string {
	return filepath.Join(rootpath, "item")
}

// UnmarshallRawItem attempts to read and unmarshall a byte slice to a raw Item. Returns a new raw Item pointer if the byte sequence is successfully deconstructed, otherwise returns nil and a corresponding error.
func UnmarshallRawItem(content *[]byte) (*RawItem, error) {
	r := &RawItem{}
	err := json.Unmarshal(*content, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// WriteRawItem writes a single raw Item pointer to a nominated destination on the running operating system. Returns nil if raw Item is successfully marshalled to JSON, otherwise returns a corresponding error.
func WriteRawItem(rawItem *RawItem) error {
	err := os.MkdirAll(StorepathRawItem(), 0644)
	if err != nil {
		return err
	}
	content, err := MarshallRawItem(rawItem)
	if err != nil {
		return err
	}
	filepath := filepath.Join(StorepathRawItem(), fmt.Sprintf("r-%s.json", rawItem.Title))
	return ioutil.WriteFile(filepath, content, 0644)
}

// rawItem defines the interface for the Raw item struct.
type rawItem interface {
	String() string
}

// A RawItem struct represents the marshalled contents for a Nintendo Amiibo Item.
// A raw Item instance contains the raw data package that is parsed within the Nintendo Amiibo XHR HTTP response.
// Raw Item's may contain references to Amiibo's in development those that contain unfinished literaray content.
type RawItem struct {
	Description  string `json:"description"`  // null
	LastModified int64  `json:"lastModified"` // 1554418285473
	Path         string `json:"path"`         // "/content/noa/en_US/amiibo/detail/wolf-link-amiibo"
	Title        string `json:"title"`        // "Wolf Link"
	URL          string `json:"url"`          // "/amiibo/detail/wolf-link-amiibo"
}

// String returns the string value for a raw Item pointer.
func (pointer *RawItem) String() string {
	return fmt.Sprintf("%s", pointer.Title)
}
