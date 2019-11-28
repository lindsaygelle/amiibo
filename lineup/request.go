package lineup

import (
	"net/http"

	"github.com/gellel/amiibo/resource"
)

var (
	// Request is a preconstructed the HTTP request to collect the XHR content from resouce.Lineup.
	Request, _ = http.NewRequest(http.MethodGet, resource.Lineup, nil)
)
