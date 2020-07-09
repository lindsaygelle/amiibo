package amiibo

import (
	"fmt"
	"time"
)

// JPNSoftware is a formatted JPNChartSoftwareItem.
type JPNSoftware struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	NameAlternative string    `json:"name_alternative"`
	Platform        string    `json:"platform"`
	Price           string    `json:"price"`
	Priority        string    `json:"priority"`
	ReleaseDate     time.Time `json:"release_date"`
	URL             string    `json:"url"`
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
	j.URL = "https://www.nintendo.co.jp/hardware/amiibo/game/" + j.ID
	return
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
