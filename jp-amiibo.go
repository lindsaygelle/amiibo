package amiibo

import (
	"fmt"
	"reflect"
	"time"
)

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
		j.URL = "https://www.nintendo.co.jp/hardware/amiibo/lineup/" + j.ID
	}
	return
}

// GetID returns the JPNAmiibo ID.
func (j JPNAmiibo) GetID() string {
	return j.ID
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
		j.URL = "https://www.nintendo.co.jp/hardware/amiibo/lineup/" + j.ID
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
