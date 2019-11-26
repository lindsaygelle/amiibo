package address

import (
	"fmt"
	"net/url"
	"strings"
)

const (
	protocol string = "http" // HTTP protocol prefix
)

const (
	protocolHTTP  string = protocol + "://"       // HTTP prefix
	protocolHTTPS string = protocol + "s" + "://" // HTTPS prefix
)

const (
	sep string = "." // sep string for parsing raw url
	rep string = ""  // rep string for parsing raw url
)

const (
	templateErr string = "%s does not start with either %s or %s" // template error invalid raw url
)

const (
	// Version is the semver of address.Address.
	Version string = "1.0.0"
)

// Address is a destructured url.URL provided by Go. Addresses are
// to give verbosity to all data pulled from the various Nintendo web resources
// that populate the content of the amiibo package.
type Address struct {
	Domain    string `json:"domain"`
	Fragment  string `json:"fragment"`
	Host      string `json:"host"`
	Hostname  string `json:"hostname"`
	Path      string `json:"path"`
	Port      string `json:"port"`
	Scheme    string `json:"scheme"`
	Subdomain string `json:"subdomain"`
	TLD       string `json:"tld"`
	URL       string `json:"url"`
	Version   string `json:"version"`
}

// NewAddress creates a new instance of the address.Address based on the
// argument raw url string provided to the function. Returns an error
// if the argument raw url does not contain a http(s)://(subdomain|www) prefix
// or if url.Parse(rawurl) cannot parse the raw url. All address.Address's
// are created in reference to a remote Nintendo source.
func NewAddress(rawurl string) (*Address, error) {
	var (
		a  *Address
		ok = (strings.HasPrefix(rawurl, protocolHTTP) || strings.HasPrefix(rawurl, protocolHTTPS))
	)
	if !ok {
		return a, fmt.Errorf(templateErr, rawurl, protocolHTTP, protocolHTTPS)
	}
	var (
		u, err = url.Parse(rawurl)
	)
	if err != nil {
		return a, err
	}
	var (
		domain    = parseDomain(u)
		fragment  = parseFragment(u)
		host      = parseHost(u)
		hostname  = parseHostname(u)
		path      = parsePath(u)
		port      = parsePort(u)
		scheme    = parseScheme(u)
		subdomain = parseSubdomain(u)
		tld       = parseTLD(subdomain, domain, hostname)
	)
	a = &Address{
		Domain:    domain,
		Fragment:  fragment,
		Host:      host,
		Hostname:  hostname,
		Path:      path,
		Port:      port,
		Scheme:    scheme,
		Subdomain: subdomain,
		TLD:       tld,
		URL:       rawurl,
		Version:   Version}
	return a, err
}

// parseDomain pares the domain from url.URL.
func parseDomain(u *url.URL) string {
	var (
		n          int
		ok         bool
		substrings = strings.Split(u.Host, sep)
	)
	n = len(substrings)
	ok = (n < 4)
	if ok {
		return substrings[n-2]
	}
	return substrings[n-3]
}

// parseFragment parses the fragment substring from url.URL.
func parseFragment(u *url.URL) string {
	return u.Fragment
}

// parseHost parses the host substring from url.URL.
func parseHost(u *url.URL) string {
	return u.Host
}

// parseHostname parses the hostname substring from url.URL.
func parseHostname(u *url.URL) string {
	return u.Hostname()
}

// parsePath parses the path substring from the url.URL.
func parsePath(u *url.URL) string {
	return u.Path
}

// parsePort parses the port substring from the url.URL.
func parsePort(u *url.URL) string {
	return u.Port()
}

// parseScheme parses the HTTP scheme substring from the url.URL.
func parseScheme(u *url.URL) string {
	return u.Scheme
}

// parseSubdomain parses the subdomain from the raw url.
func parseSubdomain(u *url.URL) string {
	var (
		n          int
		ok         bool
		subdomain  string
		substrings = strings.Split(u.Host, sep)
	)
	n = len(substrings)
	ok = (n < 4)
	if ok {
		n = (n - 3)
	} else {
		n = (n - 4)
	}
	subdomain = substrings[n]
	return subdomain
}

// parseTLD parses the top-level domain substring from the raw url.
func parseTLD(subdomain, domain, hostname string) string {
	var (
		TLD string
	)
	TLD = strings.Replace(hostname, (subdomain + sep), rep, 1)
	TLD = strings.Replace(TLD, (domain + sep), rep, 1)
	return TLD
}
