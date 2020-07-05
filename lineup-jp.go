package amiibo

// https://www.nintendo.co.jp/hardware/amiibo/chart/data/lineup.xml

// lineupJPN is the unfettered Nintendo Amiibo lineup information provided by nintendo.co.jp.
//
// lineupJPN contains the product properties related to Nintendo Amiibo products.
//
// lineupJPN is assumed to be in Japanese Hiragana.
type lineupJPN struct {

	// Item is a collection of Nintendo Amiibo products containing their product information in Japanese.
	Item []lineupItemJPN `xml:"item"`

	// SeriesItem is a collection of Nintendo Amiibo product auxiliary information.
	SeriesItem []lineupSeriesItemJP `xml:"series_item"`
}

// lineupItemJPN is the unfettered Nintendo Amiibo product lineup information from nintendo.co.jp.
type lineupItemJPN struct {

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

// lineupSeriesItemJP is the unfettered Nintendo Amiibo auxiliary metadata provided by nintendo.co.jp.
//
// lineupSeriesItemJP is assumed to be in Japanese Hiragana.
type lineupSeriesItemJP struct {

	// BGColor is the hexidecimal code for the Nintendo Amiibo product.
	BGColor string `xml:"bgcolor"`

	// Color is the hexidecimal code for the Nintendo Amiibo product.
	Color string `xml:"color"`

	// Name is the name of the Nintendo product in Japanese Hiragana.
	Name string `xml:"name"`
}
