package network

import "net/http"

var (
	// Client is the default HTTP client used for contacting the *://*.nintendo.com/*.
	Client = (&http.Client{})
)
