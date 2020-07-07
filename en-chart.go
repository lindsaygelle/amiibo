package amiibo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// ENGChartURL is the URL for the Nintendo of America Nintendo Amiibo compatibility chart.
const ENGChartURL string = "https://www.nintendo.com/content/noa/en_US/amiibo/compatibility/jcr:content/root/responsivegrid/compatibility_chart.model.json"

// ENGChart is the unfettered English language Nintendo Amiibo product and game support information.
//
// ENGChart is provided by Nintendo of America.
//
// https://www.nintendo.com/content/noa/en_US/amiibo/compatibility/jcr:content/root/responsivegrid/compatibility_chart.model.json
type ENGChart struct {

	// AmiiboList is a collection of Nintendo Amiibo product information.
	AmiiboList []ENGChartAmiibo `json:"amiiboList"`

	// GameList is a collection of Nintendo Amiibo supported games.
	GameList []ENGChartGame `json:"gameList"`

	// Items is a collection of additional Nintendo Amiibo product information.
	Items []ENGChartItem `json:"items"`
}

// GetENGChart gets the ENGChart from nintendo.com.
func GetENGChart() (req *http.Request, res *http.Response, v ENGChart, err error) {
	var b ([]byte)
	req, res, err = getRemoteFile(ENGChartURL)
	if err != nil {
		return
	}
	b, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &v)
	return
}

// ReadENGChart reads a ENGChart from disc.
func ReadENGChart(dir string, filename string) (v ENGChart, err error) {
	err = readJSONFile(dir, filename, &v)
	return v, err
}

// WriteENGChart writes a ENGChart to disc.
func WriteENGChart(dir string, filename string, v ENGChart) (fullpath string, err error) {
	fullpath, err = writeJSONFile(dir, filename, &v)
	return
}
