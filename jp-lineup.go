package amiibo

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// JPNLineupURL is the URL for the Nintendo Japan Nintendo Amiibo lineup chart.
const JPNLineupURL string = "https://www.nintendo.co.jp/hardware/amiibo/chart/data/lineup.xml"

// JPNLineup is the unfettered Japanese language Nintendo Amiibo product and game support information.
//
// JPNLineup is provided by Nintendo of America.
//
// https://www.nintendo.co.jp/hardware/amiibo/chart/data/lineup.xml
type JPNLineup struct {

	// XMLName is the xml node.
	XMLName xml.Name `xml:"items"`

	// Item is a collection of Nintendo Amiibo products containing their product information in Japanese.
	Items []JPNLineupItem `xml:"item"`

	// SeriesItem is a collection of Nintendo Amiibo product auxiliary information.
	SeriesItems []JPLineupSeriesItem `xml:"series_item"`
}

// GetJPNLineup gets the http.Response from nintendo.co.jp for the lineup Nintendo Amiibo XML.
func GetJPNLineup() (req *http.Request, res *http.Response, v JPNLineup, err error) {
	var b ([]byte)
	req, err = http.NewRequest(http.MethodGet, JPNLineupURL, nil)
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
