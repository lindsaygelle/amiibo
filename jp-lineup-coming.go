package amiibo

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

// JPNLineupComingURL is the URL for the Nintendo Japan Nintendo Amiibo lineup coming chart.
const JPNLineupComingURL string = "https://www.nintendo.co.jp/data/software/xml-system/amiibo-lineup-coming.xml"

// JPNLineupComing is the unfettered upcoming Nintendo Amiibo product information provided by nintendo.co.jp.
type JPNLineupComing struct {

	// XMLName is the xml node.
	XMLName xml.Name `xml:"data"`

	// Items is a collection of Nintendo Amiibo product information.
	Items []JPNLineupComingItem `xml:"items>item"`
}

// GetJPNLineupComing gets the http.Response from nintendo.co.jp for the lineup Nintendo Amiibo XML.
func GetJPNLineupComing() (req *http.Request, res *http.Response, v JPNLineupComing, err error) {
	var b ([]byte)
	req, res, err = getRemoteFile(JPNLineupComingURL)
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

// ReadJPNLineupComing reads a JPNLineupComing from disc.
func ReadJPNLineupComing(dir string, filename string) (v JPNLineupComing, err error) {
	err = readXMLFile(dir, filename, &v)
	return
}

// WriteJPNLineupComing writes a JPNLineupComing to disc.
func WriteJPNLineupComing(dir string, filename string, v JPNLineupComing) (fullpath string, err error) {
	fullpath, err = writeXMLFile(dir, filename, &v)
	return
}
