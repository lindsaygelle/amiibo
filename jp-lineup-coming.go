package amiibo

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// JPNLineupComing is the unfettered upcoming Nintendo Amiibo product information provided by nintendo.co.jp.
type JPNLineupComing struct {

	// XMLName is the xml node.
	XMLName xml.Name `xml:"data"`

	Items []JPNLineupComingItem `xml:"items>item"`
}

// GetJPNLineupComing gets the http.Response from nintendo.co.jp for the lineup Nintendo Amiibo XML.
func GetJPNLineupComing() (req *http.Request, res *http.Response, v JPNLineupComing, err error) {
	const URL = "https://www.nintendo.co.jp/data/software/xml-system/amiibo-lineup-coming.xml"
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
