package amiibo

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// JPNChart is the unfettered Japanese language Nintendo Amiibo product and game support information.
//
// JPNChart is provided by Nintendo Japan.
//
// https://www.nintendo.co.jp/hardware/amiibo/chart/data/chart.xml
type JPNChart struct {

	// XMLName is the xml node.
	XMLName xml.Name `xml:"items"`

	Items []JPChartItem `xml:"item"`
}

// getJPNChart gets the http.Response from nintendo.co.jp for the chart Nintendo Amiibo XML.
func getJPNChart() (req *http.Request, res *http.Response, v JPNChart, err error) {
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
