package lineup

import (
	"net/http"

	"github.com/gellel/amiibo/resource"
)

var (
	// Request is the HTTP request made to collect the XHR content from https://www.nintendo.com/.
	Request, _ = http.NewRequest(http.MethodGet, resource.Lineup, nil)
)
