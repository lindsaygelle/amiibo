package amiibo

import "encoding/xml"

// JPNLineupSeriesItem is the unfettered Nintendo Amiibo product additional information from nintendo.co.jp.
//
// JPNLineupSeriesItem is provided in Japanese Hiragana.
type JPNLineupSeriesItem struct {

	// XMLName is the xml node.
	XMLName xml.Name `xml:"series_item"`

	// BGColor is the hexidecimal code for the Nintendo Amiibo product.
	BGColor string `xml:"bgcolor"`

	// Color is the hexidecimal code for the Nintendo Amiibo product.
	Color string `xml:"color"`

	// Name is the name of the Nintendo product in Japanese Hiragana.
	Name string `xml:"name"`
}
