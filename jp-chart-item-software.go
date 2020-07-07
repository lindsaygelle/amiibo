package amiibo

// JPNChartItemSoftware is the unfettered Japanese language Nintendo software support for a Nintendo Amiibo product.
type JPNChartItemSoftware struct {

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
