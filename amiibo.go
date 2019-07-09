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

type amiibo interface {
	String() string
}

// An Amiibo struct represents the organised classification of Nintendo's toys-to-life platform.
// A populated Amiibo instances contains the normalized information collected from the RawAmiibo
// data found in the Nintendo Amiibo-list XHR HTTP response.
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

func (pointer *Amiibo) String() string {
	return fmt.Sprintf("%s", pointer.Name)
}
