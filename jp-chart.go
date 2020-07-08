package amiibo

import (
	"encoding/xml"
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
	Items []JPNChartItem `xml:"item"`
}

// GetJPNChart gets the JPNChart from nintendo.co.jp.
func GetJPNChart() (req *http.Request, res *http.Response, v JPNChart, err error) {
	var b ([]byte)
	req, res, err = getRemoteFile(JPNChartURL)
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

// ReadJPNChart reads a JPNChart from disc.
func ReadJPNChart(dir string, filename string) (v JPNChart, err error) {
	err = readXMLFile(dir, filename, &v)
	return v, err
}

// WriteJPNChart writes a JPNChart to disc.
func WriteJPNChart(dir string, filename string, v *JPNChart) (fullpath string, err error) {
	fullpath, err = writeXMLFile(dir, filename, v)
	return
}
