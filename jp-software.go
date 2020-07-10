package amiibo

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/text/language"
)

var _ Software = (JPNSoftware{})

// JPNSoftware is a formatted JPNChartSoftwareItem.
type JPNSoftware struct {
	// ID is the fully qualified ID for the Nintendo software product given by Nintendo Japan.
	ID string `json:"id"`
	// Name is the official name of the Nintendo software product.
	//
	// Name contains Japanese Hiragana.
	Name string `json:"name"`
	// NameAlternative is the alternative name given to the Nintendo software product.
	//
	// NameAlternative contains Japanese Hiragana.
	NameAlternative string `json:"name_alternative"`
	// Platform is the Nintendo hardware platform the software is available on.
	Platform string `json:"platform"`
	// Price is the price of the Nintendo software product in JPY.
	//
	// Price can be empty.
	Price    string `json:"price"`
	Priority string `json:"priority"`
	// ReleaseDate is the formatted timestamp of the Nintendo software products release date.
	ReleaseDate time.Time `json:"release_date"`
	// URL is the direct URL to the Nintendo software product page.
	URL string `json:"url"`
}

// AddJPNChartSoftwareItem adds a JPNChartSoftwareItem to the JPNSoftware.
func (j *JPNSoftware) AddJPNChartSoftwareItem(v *JPNChartSoftwareItem) (err error) {
	j.ID = v.Code
	j.Name = v.Name
	j.NameAlternative = v.NameKana
	j.Platform = v.Series
	j.Price = v.Price
	j.Priority = v.Priority
	var l = len(v.Date)
	var date = fmt.Sprintf("%s-%s-%s", v.Date[:4], v.Date[l-4:l-2], v.Date[l-2:])
	var releaseDate time.Time
	releaseDate, err = time.Parse("2006-01-02", date)
	if err == nil {
		j.ReleaseDate = releaseDate
	}
	j.URL = NintendoURLJPN + "/hardware/amiibo/game/" + j.ID
	return
}

// GetAvailable returns the JPNSoftware availability.
func (j JPNSoftware) GetAvailable() bool {
	return time.Now().After(j.ReleaseDate)
}

// GetID returns the JPNSoftware ID.
func (j JPNSoftware) GetID() string {
	return strings.TrimSuffix(filepath.Base(j.URL), ".html")
}

// GetLanguage returns the JPNSoftware language.
func (j JPNSoftware) GetLanguage() language.Tag {
	return language.Japanese
}

// GetName returns the JPNSoftware name.
func (j JPNSoftware) GetName() string {
	return j.Name
}

// GetNameAlternative returns the JPNSoftware name alternative.
func (j JPNSoftware) GetNameAlternative() string {
	return j.Name
}

// GetMD5 returns the JPNSoftware MD5.
func (j JPNSoftware) GetMD5() (MD5 string, err error) {
	var b ([]byte)
	b, err = marshal(&j, json.Marshal)
	if err != nil {
		return
	}
	MD5 = fmt.Sprintf("%x", md5.Sum(b))
	return
}

// GetReleaseDate returns the JPNSoftware release date.
func (j JPNSoftware) GetReleaseDate() time.Time {
	return j.ReleaseDate
}

// GetURL returns the JPNSoftware URL.
func (j JPNSoftware) GetURL() string {
	return j.URL
}

// ReadJPNSoftware reads a JPNSoftware from disc.
func ReadJPNSoftware(dir string, filename string) (v JPNSoftware, err error) {
	err = readJSONFile(dir, filename, &v)
	return v, err
}

// WriteJPNSoftware writes a JPNSoftware to disc.
func WriteJPNSoftware(dir string, filename string, v *JPNSoftware) (fullpath string, err error) {
	fullpath, err = writeJSONFile(dir, filename, v)
	return
}
