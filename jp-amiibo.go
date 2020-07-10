package amiibo

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"golang.org/x/text/language"
)

var _ Amiibo = (JPNAmiibo{})

// JPNAmiibo is a formatted JPN Nintendo Amiibo.
type JPNAmiibo struct {
	Chart bool `json:"chart"`
	// ID is the fully qualified ID for the Nintendo Amiibo product given by Nintendo Japan.
	ID string `json:"id"`
	// Large indicates the whether the Nintendo Amiibo product is considered a large scale item.
	Large bool `json:"large"`
	// Limited indicates the whether the Nintendo Amiibo product is a limited release.
	Limited bool `json:"limited"`
	// Name is the official name of the Nintendo Amiibo product.
	//
	// Name contains Japanese Hiragana.
	Name string `json:"name"`
	// NameAlternative is the alternative name given to the Nintendo Amiibo product.
	//
	// NameAlternative contains Japanese Hiragana.
	NameAlternative string `json:"name_alternative"`
	// New indicates the whether the Nintendo Amiibo product is a new release.
	New bool `json:"new"`
	// Price is the price of the Nintendo Amiibo product in JPY.
	//
	// Price can be empty.
	Price    string `json:"price"`
	Priority int64  `json:"priority"`
	// ReleaseDate is the formatted timestamp of the Nintendo Amiibo products release date.
	ReleaseDate time.Time `json:"release_date"`
	// ReleaseDateAlternative is the Japanese formatted Nintendo Amiibo product release date.
	ReleaseDateAlternative time.Time `json:"release_data_alternative"`
	// Series is the defined series of products that the Nintendo Amiibo product is group or affiliated with.
	Series string `json:"series"`
	// Software is a collection of Nintendo software products the Nintendo Amiibo is compatible with.
	Software JPNAmiiboSoftwareMap `json:"software"`
	// URL is the direct URL to the Nintendo Amiibo product page.
	URL string `json:"url"`
}

// AddJPNChartItem adds a JPNChartItem to the JPNAmiibo.
func (j *JPNAmiibo) AddJPNChartItem(v *JPNChartItem) (err error) {
	j.ID = v.Code
	if reflect.ValueOf(j.Software).IsZero() {
		j.Software = make(JPNAmiiboSoftwareMap)
	}
	if reflect.ValueOf(j.Name).IsZero() {
		j.Name = v.Name
	}
	if reflect.ValueOf(j.Series).IsZero() {
		j.Series = v.Series
	}
	for _, JP := range v.Softwares {
		var v JPNAmiiboSoftware
		v, err = NewJPNAmiiboSoftware(&JP)
		if err != nil {
			continue
		}
		j.Software[JP.Code] = v
	}
	if reflect.ValueOf(j.URL).IsZero() {
		j.URL = NintendoURLJPN + "/hardware/amiibo/lineup/" + j.ID
	}
	return
}

// GetAvailable returns the JPNAmiibo availability.
func (j JPNAmiibo) GetAvailable() bool {
	return time.Now().After(j.ReleaseDate)
}

// GetID returns the JPNAmiibo ID.
func (j JPNAmiibo) GetID() string {
	return j.ID
}

// GetLanguage returns the JPNAmiibo language.
func (j JPNAmiibo) GetLanguage() language.Tag {
	return language.Japanese
}

// GetName returns the JPNAmiibo name.
func (j JPNAmiibo) GetName() string {
	return j.Name
}

// GetNameAlternative returns the JPNAmiibo name alternative.
func (j JPNAmiibo) GetNameAlternative() string {
	return j.NameAlternative
}

// GetMD5 returns the JPNAmiibo MD5.
func (j JPNAmiibo) GetMD5() (MD5 string, b []byte, err error) {
	b, err = marshal(&j, json.Marshal)
	if err != nil {
		return
	}
	MD5 = fmt.Sprintf("%x", md5.Sum(b))
	return
}

// GetPrice returns the JPNAmiibo price.
func (j JPNAmiibo) GetPrice() string {
	return j.Price
}

// GetReleaseDate returns the JPNAmiibo release date.
func (j JPNAmiibo) GetReleaseDate() time.Time {
	return j.ReleaseDate
}

// GetSeries returns the JPNAmiibo series.
func (j JPNAmiibo) GetSeries() string {
	return j.Series
}

// GetURL returns the JPNAmiibo URL
func (j JPNAmiibo) GetURL() string {
	return j.URL
}

// AddJPNLineupItem adds a JPNLineupItem to the JPNAmiibo.
func (j *JPNAmiibo) AddJPNLineupItem(v *JPNLineupItem) (err error) {
	j.Chart = v.Chart != 0
	if reflect.ValueOf(j.ID).IsZero() {
		j.ID = v.Code
	}
	j.Large = v.BigSize != 0
	j.Limited = v.Limited != 0
	if reflect.ValueOf(j.Name).IsZero() {
		j.Name = v.Name
	}
	j.NameAlternative = v.NameKana
	j.New = v.New != 0
	j.Price = v.Price
	j.Priority = v.Priority
	var releaseDate time.Time
	var l = len(v.Date)
	var date = fmt.Sprintf("%s-%s-%s", v.Date[:4], v.Date[l-4:l-2], v.Date[l-2:])
	releaseDate, err = time.Parse("2006-01-02", date)
	if err == nil {
		j.ReleaseDate = releaseDate
	}
	var releaseDateAlternative time.Time
	releaseDateAlternative, err = time.Parse("2006-01-02", v.DisplayDate)
	if err == nil {
		j.ReleaseDateAlternative = releaseDateAlternative
	}
	err = nil
	j.Series = v.Series
	if reflect.ValueOf(j.URL).IsZero() {
		j.URL = NintendoURLJPN + "/hardware/amiibo/lineup/" + j.ID
	}
	return
}

// NewJPNAmiibo returns a new NewJPNAmiibo.
func NewJPNAmiibo(JPNChartItem *JPNChartItem, JPNLineupItem *JPNLineupItem) (v JPNAmiibo, err error) {
	var ok bool
	ok = JPNChartItem.GetID() == JPNLineupItem.GetID()
	if !ok {
		err = fmt.Errorf("JPNChartItem != JPNLineupItem")
	}
	if err != nil {
		return
	}
	err = (&v).AddJPNChartItem(JPNChartItem)
	if err != nil {
		return
	}
	if err != nil {
		return
	}
	err = (&v).AddJPNLineupItem(JPNLineupItem)
	return
}

// ReadJPNAmiibo reads a JPNAmiibo from disc.
func ReadJPNAmiibo(dir string, filename string) (v JPNAmiibo, err error) {
	err = readJSONFile(dir, filename, &v)
	return v, err
}

// WriteJPNAmiibo writes a JPNAmiibo to disc.
func WriteJPNAmiibo(dir string, filename string, v *JPNAmiibo) (fullpath string, err error) {
	fullpath, err = writeJSONFile(dir, filename, v)
	return
}
