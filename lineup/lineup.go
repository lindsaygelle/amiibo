package lineup

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gellel/amiibo/network"
)

// Get performs a HTTP request to Nintendo Amiibo lineup resource and unmarshal's the
// HTTP response body on http.StatusOK. Throws an error if the Nintendo server
// returns anything other than http.StatusOK. If the response content cannot be
// handled by json.Unmarshal the corresponding error message is returned.
func Get() (*XHR, error) {
	var (
		data     []byte
		res, err = network.Client.Do(Request)
		xhr      XHR
	)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(res.Status)
	}
	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &xhr)
	if err != nil {
		return nil, err
	}
	xhr.ContentLength = res.ContentLength
	xhr.Cookies = res.Cookies()
	xhr.Headers = res.Header
	return &xhr, err
}
