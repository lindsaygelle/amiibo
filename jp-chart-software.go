package amiibo

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

const JPNChartSoftwareURL string = "https://www.nintendo.co.jp/hardware/amiibo/chart/data/software.xml"

// JPNChartSoftware is the unfettered Japanese language Nintendo software information.
type JPNChartSoftware struct {
	XMLName xml.Name `xml:"items"`

	Items []JPNChartSoftwareItem `xml:"items>item"`
}

// GetJPNChartSoftware gets the JPNChartSoftware from nintendo.co.jp.
func GetJPNChartSoftware() (req *http.Request, res *http.Response, v JPNChartSoftware, err error) {
	var b ([]byte)
	req, res, err = getRemoteFile(JPNChartSoftwareURL)
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

// ReadJPNChartSoftware reads a JPNChartSoftware from disc.
func ReadJPNChartSoftware(dir string, filename string) (v JPNChartSoftware, err error) {
	err = readXMLFile(dir, filename, &v)
	return v, err
}

// WriteJPNChartSoftware writes a JPNChartSoftware to disc.
func WriteJPNChartSoftware(dir string, filename string, v *JPNChartSoftware) (fullpath string, err error) {
	fullpath, err = writeXMLFile(dir, filename, &v)
	return
}
