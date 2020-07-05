package amiibo

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// lineupComingJPN is the unfettered upcoming Nintendo Amiibo product information provided by nintendo.co.jp.
type lineupComingJPN struct {

	// XMLName is the xml node.
	XMLName xml.Name `xml:"data"`

	Items []lineupComingItemJPN `xml:"items>item"`
}

// lineupComingItemJPN is the unfettered upcoming Nintendo Amiibo product information provided by nintendo.co.jp.
type lineupComingItemJPN struct {

	// XMLName is the xml node.
	XMLName xml.Name `xml:"item"`

	// AmiiboLabel is the label for the Nintendo Amiibo product.
	AmiiboLabel string `xml:"amiibo_label"`

	// AmiiboLink is the URL for the Nintendo Amiibo product.
	AmiiboLink string `xml:"amiibo_link"`

	// AmiiboSeries is the Japanese Hiragana for the Nintendo product the Amiibo is affiliated with.
	AmiiboSeries string `xml:"amiibo_series"`

	// D is the YYYY-MM-DD datestamp for the Nintendo Amiibo product.
	D string `xml:"d"`

	// Link is the URL to the Nintendo Amiibo product.
	Link string `xml:"link"`

	// LinkTarget is the relative URL to the Nintendo Amiibo product.
	LinkTarget string `xml:"link_target"`

	// Memo is the verbose title for the Nintendo Amiibo product.
	Memo string `xml:"memo"`

	// Price is the price of the Nintendo Amiibo in Japanese Hiragana.
	Price string `xml:"price"`

	// ReleaseDateStr is the datestamp for the Nintendo Amiibo product release date in Japanese Hiragana.
	ReleaseDateStr string `xml:"release_date_str"`

	// ThumbVariation is the alternative namespace for the Nintendo Amiibo product.
	ThumbVariation string `xml:"thumb_variation"`

	// Title is the Japanese title in either Hiragana or Kanji for Nintendo Amiibo product.
	Title string `xml:"title"`

	// TitleRuby is the Japanese Hiragana for the Nintendo Amiibo product.
	TitleRuby string `xml:"title_ruby"`
}

// getLineupComingJPN gets the http.Response from nintendo.co.jp for the lineup Nintendo Amiibo XML.
func getLineupComingJPN() (req *http.Request, res *http.Response, v lineupComingJPN, err error) {
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
