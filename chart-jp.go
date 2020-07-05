package amiibo

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// chart is the unfettered Nintendo Amiibo chart information provided by nintendo.co.jp.
//
// chart contains the Japanese language Nintendo Amiibo software compatability.
//
// chart is provided in Japanese Hiragana.
type chartJPN struct {

	// XMLName is the xml node.
	XMLName xml.Name `xml:"items"`

	Items []itemJPN `xml:"item"`
}

// itemJPN is the unfettered Nintendo Amiibo chart information provided by nintendo.co.jp.
//
// itemJPN contains the simplified Nintendo Amiibo product information.
//
// itemJPN is in Japanese Hiragana.
//
// itemJPN is provided as XML from nintendo.co.jp.
type itemJPN struct {

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
	Softwares []softwareJP `xml:"softwares"`
}

// softwareJP is the software support information for a Nintendo Amiibo chart item provided by nintendo.co.jp.
//
// softwareJP is in Japanese Hiragana.
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

// getChartJPN gets the http.Response from nintendo.co.jp for the chart Nintendo Amiibo XML.
func getChartJPN() (req *http.Request, res *http.Response, v chartJPN, err error) {
	const URL = "https://www.nintendo.co.jp/hardware/amiibo/chart/data/chart.xml"
	var b ([]byte)
	req, err = http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return
	}
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf(("http: %d"), res.StatusCode)
	}
	if err != nil {
		return
	}
	b, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	err = xml.Unmarshal(b, &v)
	if err != nil {
		return
	}
	return
}
