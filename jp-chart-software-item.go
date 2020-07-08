package amiibo

import "encoding/xml"

// JPNChartSoftwareItem is the unfettered Japanese language Nintendo Amiibo software information.
type JPNChartSoftwareItem struct {
	XMLName xml.Name `xml:"item"`

	Blank       interface{} `xml:"blank"`
	Chart       int         `xml:"chart"`
	Chartseries string      `xml:"chartseries"`
	Code        string      `xml:"code"`
	Date        string      `xml:"date"`
	DisplayDate string      `xml:"dispalydate"`
	Icon        int         `xml:"icon"`
	Info        interface{} `xml:"info"`
	Name        string      `xml:"name"`
	NameKana    string      `xml:"nameKana"`
	New         int         `xml:"new"`
	Price       string      `xml:"price"`
	Priority    string      `xml:"priority"`
	Series      string      `xml:"series"`
	Softoption  interface{} `xml:"softoption"`
	Subname     string      `xml:"subname"`
}

// GetID returns the JPNChartSoftwareItem ID.
func (j JPNChartSoftwareItem) GetID() string {
	return j.Code
}
