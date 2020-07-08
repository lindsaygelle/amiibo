package amiibo

import (
	"encoding/xml"
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
	SeriesItems []JPNLineupSeriesItem `xml:"series_item"`
}

// GetJPNLineup gets the http.Response from nintendo.co.jp for the lineup Nintendo Amiibo XML.
func GetJPNLineup() (req *http.Request, res *http.Response, v JPNLineup, err error) {
	var b ([]byte)
	req, res, err = getRemoteFile(JPNLineupURL)
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

// ReadJPNLineup reads a JPNLineup from disc.
func ReadJPNLineup(dir string, filename string) (v JPNLineup, err error) {
	err = readXMLFile(dir, filename, &v)
	return
}

// WriteJPNLineup writes a JPNLineup to disc.
func WriteJPNLineup(dir string, filename string, v *JPNLineup) (fullpath string, err error) {
	fullpath, err = writeXMLFile(dir, filename, v)
	return
}
