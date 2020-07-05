package amiibo

import (
	"encoding/json"
	"fmt"
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
	AmiiboList []ENGChart `json:"amiiboList"`

	// GameList is a collection of Nintendo Amiibo supported games.
	GameList []ENGChartGame `json:"gameList"`

	// Items is a collection of additional Nintendo Amiibo product information.
	Items []ENGChartItem `json:"items"`
}

// GetENGChart gets the ENGChart from nintendo.com.
func GetENGChart() (req *http.Request, res *http.Response, v ENGChart, err error) {
	var b ([]byte)
	req, err = http.NewRequest(http.MethodGet, ENGChartURL, nil)
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
	err = json.Unmarshal(b, &v)
	return
}
