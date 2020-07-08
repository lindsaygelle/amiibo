package amiibo

import "encoding/xml"

// JPNLineupItem is the unfettered Nintendo Amiibo product information from nintendo.co.jp.
//
// JPNLineupItem is provided in Japanese Hiragana.
type JPNLineupItem struct {

	// XMLName is the xml node.
	XMLName xml.Name `xml:"item"`

	// BigSize is a integer representation of a boolean.
	//
	// BigSize relates to the scale of the Nintendo Amiibo product.
	BigSize int `xml:"bigsize"`

	// Chart is a integer representation of a boolean.
	//
	// Chart relates to the occurrence of the Nintendo Amiibo product in the chart XML.
	Chart int64 `xml:"chart"`

	// Code is the English ID for the Nintendo Amiibo product from the Japanese CDN.
	Code string `xml:"code"`

	// Date is the YYYYMMDD expression of the Nintendo Amiibo product release date.
	Date string `xml:"date"`

	// DisplayDate is the Japanese Hiragana expression of the Nintedo Amiibo product release date.
	//
	// DisplayDate currently (5/07/2020 (DD-MM-YYYY)) has a typo on the nintendo.co.jp and exists
	// as dispalydate.
	DisplayDate string `xml:"displayDate"`

	// Limited is a integer representation of a boolean.
	//
	// Limited relates to the rareness of the Nintendo Amiibo product.
	Limited int `xml:"limited"`

	// Name is the name of the Nintendo product in Japanese Hiragana.
	Name string `xml:"name"`

	// NameKana is the name of the Nintendo Amiibo product in Japanese Hiragana.
	NameKana string `xml:"nameKana"`

	// New is a integer representation of a boolean.
	//
	// New relates to the newness of the Nintendo Amiibo product.
	New int `xml:"new"`

	// Price is the price of the Nintendo Amiibo product in Japanese Yen.
	Price string `xml:"price"`

	// Priority is the numerical rank of the Nintendo Amiibo product.
	Priority int64 `xml:"priority"`

	// Series is the Japanese Hiragana for the Nintendo product that the Nintendo Amiibo product is affiliated with.
	//
	// Series will need to be translated from Japanese to English.
	Series string `xml:"series"`
}

// GetID returns the JPNLineItem ID.
func (j JPNLineupItem) GetID() string {
	return j.Code
}

// ReadJPNLineupItem reads a JPNLineupItem from disc.
func ReadJPNLineupItem(dir string, filename string) (v JPNLineupItem, err error) {
	err = readXMLFile(dir, filename, &v)
	return v, err
}

// WriteJPNLineupItem writes a JPNLineupItem to disc.
func WriteJPNLineupItem(dir string, filename string, v *JPNLineupItem) (fullpath string, err error) {
	fullpath, err = writeXMLFile(dir, filename, v)
	return
}
