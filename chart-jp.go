package amiibo

// https://www.nintendo.co.jp/hardware/amiibo/chart/data/chart.xml

// chart is the unfettered Nintendo Amiibo chart information provided by nintendo.co.jp.
//
// chart contains the Japanese language Nintendo Amiibo software compatability.
//
// chart is assumed to be in Japanese hiragana.
type chart struct {
	Items []itemJPN `xml:"items"`
}

// itemJPN is the unfettered Nintendo Amiibo chart information provided by nintendo.co.jp.
//
// itemJPN contains the simplified Nintendo Amiibo product information.
//
// itemJPN is assumed to be in Japanese Hiragana.
//
// itemJPN is provided as XML from nintendo.co.jp.
type itemJPN struct {

	// Code is the ID code for the Nintendo product.
	Code string `xml:"code"`

	// Name is the name of the Nintendo product in Japanese Hiragana.
	Name string `xml:"name"`

	// Series is the Japanese Hiragana for the Nintendo product that the Nintendo Amiibo product is affiliated with.
	//
	// Series will need to be translated from Japanese to English.
	Series string `xml:"series"`

	// Softwares is a collection of metadata that the Nintendo Amiibo product integrates with.
	Softwares []softwareJP `xml:"softwares"`
}

// softwareJP is the software support information for a Nintendo Amiibo chart item provided by nintendo.co.jp.
//
// softwareJP is assumed to be in Japanese Hiragana.
//
// softwareJP is provided as XML from nintedo.co.jp.
type softwareJP struct {

	// Code is the ID code for the Nintendo product.
	Code string `xml:"code"`

	// Name is the name of the Nintendo product in Japanese Hiragana.
	Name string `xml:"name"`

	// More is the verbose description for the Nintendo Amiibo chart item.
	More string `xml:"more"`

	// Pickup is a provided property with an unclear purpose.
	Pickup int64 `xml:"pickup"`

	// ReadWrite is a provided property with an unclear purpose.
	ReadWrite string `xml:"readwrite"`
}
