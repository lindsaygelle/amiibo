package lineup

import "net/http"

// XHR is the content HTTP response that is requested from https://www.nintendo.com/amiibo/line-up/.
// XHR contains the structured information provided from Nintendo as-is and is updated
// with the release of new Nintendo Amiibo products. All content within the XHR
// response body describes the Amiibo content in the context of the lineup, meaning information
// such as the compatibility of the figuring is not offered.
type XHR struct {
	Amiibo           []*Amiibo      `json:"amiiboList"`
	ComponentPath    string         `json:"componentPath"`
	ContentLength    int64          `json:"contentLength"`
	Cookies          []*http.Cookie `json:"cookies"`
	DateFormatString string         `json:"dataFormatString"`
	Headers          http.Header    `json:"headers"`
	Items            []*Item        `json:"items"`
	Status           string         `json:"status"`
	StatusCode       string         `json:"statusCode"`
}
