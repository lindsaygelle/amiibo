package network

import (
	"net/http"
	"time"
)

const (
	// Timeout is the default HTTP client timeout in seconds.
	Timeout time.Duration = time.Second * 5
)

var (
	// Client is the default HTTP client used for contacting the *://*.nintendo.com/*.
	Client = (&http.Client{
		Timeout: Timeout})
)
