package amiibo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// ENGLineupURL is the URL to the Nintendo of America Nintendo Amiibo lineup information.
const ENGLineupURL string = "https://www.nintendo.com/content/noa/en_US/amiibo/line-up/jcr:content/root/responsivegrid/lineup.model.json"

// ENGLineup is the unfettered Nintendo Amiibo lineup information provided by nintendo.com.
//
// ENGLineup contains the product information for the Nintendo Amiibo product as well as some related metadata.
//
// ENGLineup is provided by Nintendo of America.
//
// https://www.nintendo.com/content/noa/en_US/amiibo/line-up/jcr:content/root/responsivegrid/lineup.model.json
type ENGLineup struct {

	// AmiiboList is a collection of Nintendo Amiibo products containing their product information.
	AmiiboList []ENGLineupAmiibo `json:"amiiboList"`

	// ComponentPath is the relative path to the Nintendo resource file.
	ComponentPath string `json:"componentPath"`

	// Items is a collection of metadata related to Nintendo Amiibo products.
	Items []ENGLineupItem `json:"items"`
}

// GetENGLineup gets the http.Response from nintendo.com for the lineup Nintendo Amiibo JSON.
func GetENGLineup() (req *http.Request, res *http.Response, v ENGLineup, err error) {
	var b ([]byte)
	req, res, err = getRemoteFile(ENGLineupURL)
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

// ReadENGLineup reads a ENGLineup from disc.
func ReadENGLineup(dir string, filename string) (v ENGLineup, err error) {
	err = readJSONFile(dir, filename, &v)
	return
}

// WriteENGLineup writes a ENGLineup to disc.
func WriteENGLineup(dir string, filename string, v ENGLineup) (fullpath string, err error) {
	fullpath, err = writeJSONFile(dir, filename, &v)
	return
}
