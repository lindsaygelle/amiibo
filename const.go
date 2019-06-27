// Package amiibo fetches the current Nintendo Amiibo statues, cards and more that are available or in development.
// Uses a found XHR HTTP response on the Nintendo Amiibo URI.
package amiibo

const (
	// URL is the endpoint used to fetch the raw Amiibo JSON.
	URL string = "https://www.nintendo.com/content/noa/en_US/amiibo/line-up/jcr:content/root/responsivegrid/lineup.model.json?linkItems=true"
)
const (
	timeLayout string = "2006-01-02T15:04:05.000Z"
)
