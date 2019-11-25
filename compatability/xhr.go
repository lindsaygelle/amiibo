package compatability

import "net/http"

// XHR is the content HTTP response that is requested from https://www.nintendo.com/amiibo/compatability/.
// XHR contains the structured information provided from Nintendo as-is and is updated
// with the release of new Nintendo Amiibo products. All content within the XHR
// response body describes the Nintendo Amiibo products in the context of their compatability with
// other Nintendo products, meaning information such as the build information or meta data of
// the product is not provided.
type XHR struct {
	Amiibo           []*Amiibo      `json:"amiiboList"`
	AuthorMode       bool           `json:"authorMode"`
	ComponentPath    string         `json:"componentPath"`
	ContentLength    int64          `json:"contentLength"`
	Cookies          []*http.Cookie `json:"cookies"`
	DateFormatString string         `json:"dataFormatString"`
	Headers          http.Header    `json:"headers"`
	Games            []*Game        `json:"gameList"`
	Items            []*Item        `json:"items"`
	Language         string         `json:"language"`
	Mode             string         `json:"mode"`
	Status           string         `json:"status"`
	StatusCode       int            `json:"statusCode"`
}
