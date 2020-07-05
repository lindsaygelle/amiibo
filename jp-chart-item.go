package amiibo

import "encoding/xml"

// JPChartItem is the unfettered Japanese language Nintendo Amiibo product information.
type JPChartItem struct {

	// XMLName is the xml node.
	XMLName xml.Name `xml:"item"`

	// Code is the ID code for the Nintendo product.
	Code string `xml:"code"`

	// Name is the name of the Nintendo product in Japanese Hiragana.
	Name string `xml:"name"`

	// Series is the Japanese Hiragana for the Nintendo product that the Nintendo Amiibo product is affiliated with.
	//
	// Series will need to be translated from Japanese to English.
	Series string `xml:"series"`

	// Softwares is a collection of metadata that the Nintendo Amiibo product integrates with.
	Softwares []JPNChartItemSoftware `xml:"softwares"`
}
