package amiibo

import "encoding/xml"

// JPNChartSoftwareItem is the unfettered Japanese language Nintendo Amiibo software information.
type JPNChartSoftwareItem struct {

	// XMLName is the xml node.
	XMLName xml.Name `xml:"item"`

	Blank interface{} `xml:"blank"`

	// Chart is a integer representation of a boolean.
	//
	// Chart relates to the occurrence of the Nintendo Amiibo product in the chart XML.
	Chart int `xml:"chart"`

	// Chartseries is the Japanese Hiragana for the Nintendo hardware system that the software is supported for.
	//
	// Chartseries will need to be translated from Japanese to English.
	Chartseries string `xml:"chartseries"`

	// Code is the ID code for the Nintendo product.
	Code string `xml:"code"`

	// Date is the YYYYMMDD expression of the Nintendo Amiibo product release date.
	Date string `xml:"date"`

	// DisplayDate is the Japanese Hiragana expression of the Nintedo Amiibo product release date.
	//
	// DisplayDate currently (2019年3月7日発売) has a typo on the nintendo.co.jp and exists
	// as dispalydate.
	DisplayDate string `xml:"dispalydate"`

	Icon int `xml:"icon"`

	Info interface{} `xml:"info"`

	// Name is the name of the Nintendo product in Japanese Hiragana.
	Name string `xml:"name"`

	// NameAlternative is the alternative name given to the Nintendo Amiibo product.
	//
	// NameAlternative contains Japanese Hiragana.
	NameKana string `xml:"nameKana"`

	// New is a integer representation of a boolean.
	//
	// New relates to the newness of the Nintendo software product.
	New int `xml:"new"`

	// Price is the price of the Nintendo Amiibo product in Japanese Yen.
	Price string `xml:"price"`

	// Priority is the numerical rank of the Nintendo Amiibo product.
	Priority string `xml:"priority"`

	// Series is the Japanese Hiragana for the Nintendo hardware system that the software is supported for.
	//
	// Series will need to be translated from Japanese to English.
	Series string `xml:"series"`

	SoftOption interface{} `xml:"softoption"`

	Subname string `xml:"subname"`
}

// GetID returns the JPNChartSoftwareItem ID.
func (j JPNChartSoftwareItem) GetID() string {
	return j.Code
}

// ReadJPNChartSoftwareItem reads a JPNChartSoftwareItem from disc.
func ReadJPNChartSoftwareItem(dir string, filename string) (v JPNChartSoftwareItem, err error) {
	err = readXMLFile(dir, filename, &v)
	return v, err
}

// WriteJPNChartSoftwareItem writes a JPNChartSoftwareItem to disc.
func WriteJPNChartSoftwareItem(dir string, filename string, v *JPNChartSoftwareItem) (fullpath string, err error) {
	fullpath, err = writeXMLFile(dir, filename, v)
	return
}
