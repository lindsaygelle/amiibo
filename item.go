package amiibo

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	_ item = (*Item)(nil)
)

func deleteItem(item *Item) error {
	return os.Remove(filepath.Join(storepathItem(), fmt.Sprintf("%s.json", item.Name)))
}

func getItem(ID string) *Item {
	if ok := strings.HasSuffix(ID, ".json"); !ok {
		ID = fmt.Sprintf("%s.json", ID)
	}
	b, err := openItem(ID)
	if err != nil {
		return nil
	}
	amiibo, err := unmarshallItem(b)
	if err != nil {
		return nil
	}
	return amiibo
}

func marshallItem(item *Item) ([]byte, error) {
	content, err := json.Marshal(item)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func openItem(name string) (*[]byte, error) {
	filepath := filepath.Join(storepathItem(), name)
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

func storepathItem() string {
	return filepath.Join(rootpath, "item")
}

func unmarshallItem(content *[]byte) (*Item, error) {
	r := &Item{}
	err := json.Unmarshal(*content, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func writeItem(item *Item) error {
	err := os.MkdirAll(storepathItem(), 0644)
	if err != nil {
		return err
	}
	content, err := marshallItem(item)
	if err != nil {
		return err
	}
	filepath := filepath.Join(storepathItem(), fmt.Sprintf("%s.json", item.Name))
	return ioutil.WriteFile(filepath, content, 0644)
}

func newItem(r *RawItem) *Item {
	return &Item{
		Description: r.Description,
		ID:          fmt.Sprintf("%x", md5.Sum([]byte(r.Title))),
		Name:        (reStripName.ReplaceAllString(r.Title, "")),
		Path:        r.Path,
		Timestamp:   (time.Unix(r.LastModified, 0).UTC()),
		URL:         (nintendoURL + r.URL)}
}

// item defines the interface for a Item struct.
type item interface {
	String() string
}

// An Item struct represents the shortform data about upcoming Amiibo items from Nintendo.
// An normalized Item instance contains the formatted and cleaned information collected from the raw Item
// data that exists in the Nintendo Amiibo XHR HTTP response.
type Item struct {
	Description string    `json:"description"` // RawItem.Description
	ID          string    `json:"id"`          // Hash.MD5
	Name        string    `json:"name"`        // RawItem.Title
	Path        string    `json:"path"`        // RawItem.Path
	Timestamp   time.Time `json:"timestamp"`   // RawItem.LastModified
	URL         string    `json:"url"`         // RawItem.URL
}

// String returns the string value for the Item pointer.
func (pointer *Item) String() string {
	return fmt.Sprintf("%s", pointer.Name)
}
