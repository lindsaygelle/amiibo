package amiibo

import (
	"fmt"
	"reflect"
	"time"
)

// JPNAmiibo is a formatted JPN Nintendo Amiibo.
type JPNAmiibo struct {
	Chart                  bool                `json:"chart"`
	ID                     string              `json:"id"`
	Large                  bool                `json:"large"`
	Limited                bool                `json:"limited"`
	Name                   string              `json:"name"`
	NameAlternative        string              `json:"name_alternative"`
	New                    bool                `json:"new"`
	Price                  string              `json:"price"`
	Priority               int64               `json:"priority"`
	ReleaseDate            time.Time           `json:"release_date"`
	ReleaseDateAlternative time.Time           `json:"release_data_alternative"`
	Series                 string              `json:"series"`
	Software               []JPNAmiiboSoftware `json:"software"`
	URL                    string              `json:"url"`
}

// AddJPNChartItem adds a JPNChartItem to the JPNAmiibo.
func (j *JPNAmiibo) AddJPNChartItem(v JPNChartItem) (err error) {
	j.ID = v.Code
	if !reflect.ValueOf(j.Name).IsZero() {
		j.Name = v.Name
	}
	if !reflect.ValueOf(j.Series).IsZero() {
		j.Series = v.Series
	}
	for _, v := range v.Softwares {
		var JPNAmiiboSoftware JPNAmiiboSoftware
		JPNAmiiboSoftware, err = NewJPNAmiiboSoftware(v)
		if err != nil {
			continue
		}
		j.Software = append(j.Software, JPNAmiiboSoftware)
	}
	return
}

// AddJPNLineupItem adds a JPNLineupItem to the JPNAmiibo.
func (j *JPNAmiibo) AddJPNLineupItem(v JPNLineupItem) (err error) {
	j.Chart = v.Chart != 0
	j.Large = v.BigSize != 0
	j.Limited = v.Limited != 0
	if !reflect.ValueOf(j.Name).IsZero() {
		j.Name = v.Name
	}
	j.NameAlternative = v.NameKana
	j.New = v.New != 0
	j.Price = v.Price
	j.Priority = v.Priority
	var releaseDate time.Time
	releaseDate, _ = time.Parse("2006-01-02", v.Date)
	j.ReleaseDate = releaseDate
	var releaseDateAlternative time.Time
	releaseDateAlternative, _ = time.Parse("2006-01-02", v.DisplayDate)
	j.ReleaseDateAlternative = releaseDateAlternative
	j.Series = v.Series
	return
}

func NewJPNAmiibo(JPNChartItem JPNChartItem, JPNLineupItem JPNLineupItem) (v JPNAmiibo, err error) {
	var ok bool
	ok = JPNChartItem.GetID() == JPNLineupItem.GetID()
	if !ok {
		err = fmt.Errorf("JPNChartItem != JPNLineupItem")
	}
	if err != nil {
		return
	}
	return
}
