package compatability

import (
	"net/http"

	"github.com/gellel/amiibo/resource"
)

var (
	// Request is a preconstructed the HTTP request to collect the XHR content from resource.Compatability.
	Request, _ = http.NewRequest(http.MethodGet, resource.Compatability, nil)
)
