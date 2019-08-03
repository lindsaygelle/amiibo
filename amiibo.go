package amiibo

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/text/currency"
)

var (
	_ amiibo = (*Amiibo)(nil)
)

func deleteAmiibo(amiibo *Amiibo) error {
	return os.Remove(filepath.Join(storepathAmiibo(), fmt.Sprintf("%s.json", amiibo.Hex)))
}

func getAmiibo(ID string) *Amiibo {
	if ok := strings.HasSuffix(ID, ".json"); !ok {
		ID = fmt.Sprintf("%s.json", ID)
	}
	b, err := openAmiibo(ID)
	if err != nil {
		return nil
	}
	amiibo, err := unmarshallAmiibo(b)
	if err != nil {
		return nil
	}
	return amiibo
}

func marshallAmiibo(amiibo *Amiibo) ([]byte, error) {
	content, err := json.Marshal(amiibo)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func openAmiibo(name string) (*[]byte, error) {
	filepath := filepath.Join(storepathAmiibo(), name)
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

func storepathAmiibo() string {
	return filepath.Join(rootpath, "amiibo")
}

func unmarshallAmiibo(content *[]byte) (*Amiibo, error) {
	r := &Amiibo{}
	err := json.Unmarshal(*content, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func writeAmiibo(fullpath string, amiibo *Amiibo) error {
	err := os.MkdirAll(storepathAmiibo(), 0644)
	if err != nil {
		return err
	}
	content, err := marshallAmiibo(amiibo)
	if err != nil {
		return err
	}
	filepath := filepath.Join(storepathAmiibo(), fmt.Sprintf("%s.json", amiibo.Hex))
	return ioutil.WriteFile(filepath, content, 0644)
}

// NewAmiibo returns a new Amiibo pointer from a raw Amiibo pointer. Normalizes the raw Amiibo fields into
// predictable patterns as well as cleans the input data. Does not modify the raw Amiibo allowing
// the original content to be accessed. Assumes that the argument raw Amiibo pointer has been
// successfully marshalled and contains all provided fields.
func newAmiibo(r *RawAmiibo) *Amiibo {
	var (
		t, _ = time.Parse(timeLayoutRelease, r.ReleaseDateMask)
		desc = reStripSpaces.ReplaceAllString(reStripHTML.ReplaceAllString(r.OverviewDescription, " "), " ")
		h    = md5.Sum([]byte(r.AmiiboName))
	)
	return &Amiibo{
		Available:   r.IsReleased,
		Box:         (nintendoURL + r.BoxArtURL),
		Code:        r.GameCode,
		Description: html.UnescapeString(strings.TrimSpace(desc)),
		Franchise:   r.Franchise,
		Figure:      (nintendoURL + r.FigureURL),
		Hex:         strings.ToUpper(strings.Replace(r.HexCode, "#", "", 1)),
		ID:          fmt.Sprintf("%x", h),
		Name:        (reStripName.ReplaceAllString(r.AmiiboName, "")),
		Page:        (nintendoURL + r.DetailsURL),
		Path:        r.DetailsPath,
		Presenter:   (strings.Replace(r.PresentedBy, "noa:publisher/", "", -1)),
		Price:       new(currency.Amount),
		Release:     t,
		Series:      r.Series,
		Slug:        r.Slug,
		Timestamp:   (time.Unix(r.UnixTimestamp, 0).UTC()),
		Type:        r.Type,
		UPC:         r.UPC,
		URL:         (nintendoURL + r.AmiiboPage)}
}

// amiibo defines the interface for the Amiibo struct pointer.
type amiibo interface {
	String() string
}

// An Amiibo struct represents the character classification of Nintendo's toys-to-life platform, Amiibo.
// A normalized Amiibo instance contains the formatted and cleaned information collected from the raw Amiibo
// data that exists in the Nintendo Amiibo list XHR HTTP response.
type Amiibo struct {
	Available   bool             `json:"available"`   // RawAmiibo.IsReleased
	Box         string           `json:"box"`         // RawAmiibo.BoxArtURL
	Code        string           `json:"code"`        // RawAmiibo.GameCode
	Description string           `json:"description"` // RawAmiibo.OverviewDescription
	Figure      string           `json:"figure"`      // RawAmiibo.FigureURL
	Franchise   string           `json:"franchise"`   // RawAmiibo.Franchise
	Hex         string           `json:"hex"`         // RawAmiibo.HexCode
	ID          string           `json:"id"`          // Hash.md5
	Name        string           `json:"name"`        // RawAmiibo.Name
	Page        string           `json:"page"`        // RawAmiibo.DetailsURL
	Path        string           `json:"path"`        // RawAmiibo.DetailsPath
	Presenter   string           `json:"presenter"`   // RawAmiibo.PresentedBy
	Price       *currency.Amount `json:"price"`       // RawAmiibo.Price
	Release     time.Time        `json:"release"`     // RawAmiibo.ReleaseDateMask
	Series      string           `json:"series"`      // RawAmiibo.Series
	Slug        string           `json:"slug"`        // RawAmiibo.Slug
	Timestamp   time.Time        `json:"timestamp"`   // RawAmiibo.UnixTimestamp
	Type        string           `json:"type"`        // RawAmiibo.Type
	UPC         string           `json:"upc"`         // RawAmiibo.UPC
	URL         string           `json:"url"`         // RawAmiibo.AmiiboPage
}

// Strign returns the string value of the Amiibo pointer.
func (pointer *Amiibo) String() string {
	return fmt.Sprintf("%s", pointer.Name)
}
