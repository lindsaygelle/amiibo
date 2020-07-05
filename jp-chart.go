package amiibo

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// JPNChartURL is the URL for the Nintendo Japan Nintendo Amiibo compatibility chart.
const JPNChartURL string = "https://www.nintendo.co.jp/hardware/amiibo/chart/data/chart.xml"

// JPNChart is the unfettered Japanese language Nintendo Amiibo product and game support information.
//
// JPNChart is provided by Nintendo Japan.
//
// https://www.nintendo.co.jp/hardware/amiibo/chart/data/chart.xml
type JPNChart struct {

	// XMLName is the xml node.
	XMLName xml.Name `xml:"items"`

	// Items is a collection of Nintendo Amiibo product information.
	Items []JPChartItem `xml:"item"`
}

// GetJPNChart gets the JPNChart from nintendo.co.jp.
func GetJPNChart() (req *http.Request, res *http.Response, v JPNChart, err error) {
	var b ([]byte)
	req, err = http.NewRequest(http.MethodGet, JPNChartURL, nil)
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
