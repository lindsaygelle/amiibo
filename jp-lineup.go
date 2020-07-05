package amiibo

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// JPNLineup is the unfettered Nintendo Amiibo lineup information provided by nintendo.co.jp.
//
// JPNLineup contains the product properties related to Nintendo Amiibo products.
//
// JPNLineup is provided in Japanese Hiragana.
type JPNLineup struct {

	// XMLName is the xml node.
	XMLName xml.Name `xml:"items"`

	// Item is a collection of Nintendo Amiibo products containing their product information in Japanese.
	Items []JPNLineupItem `xml:"item"`

	// SeriesItem is a collection of Nintendo Amiibo product auxiliary information.
	SeriesItems []JPLineupSeriesItem `xml:"series_item"`
}

// getJPNLineup gets the http.Response from nintendo.co.jp for the lineup Nintendo Amiibo XML.
func getJPNLineup() (req *http.Request, res *http.Response, v JPNLineup, err error) {
	const URL = "https://www.nintendo.co.jp/hardware/amiibo/chart/data/lineup.xml"
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
