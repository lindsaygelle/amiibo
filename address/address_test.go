package address_test

import (
	"testing"

	"github.com/gellel/amiibo/address"
)

const (
	domain    string = "nintendo"
	host      string = subdomain + "." + domain + "." + tld
	hostname  string = host + ":" + port
	fragment  string = "fragment"
	path      string = "/"
	port      string = "443"
	scheme    string = "http"
	subdomain string = "www"
	tld       string = "co.jp"
)

const (
	templateErr string = "address.Address.%s: a.(%s) != %s"
)

const (
	rawurl string = scheme + "://" + subdomain + "." + domain + "." + tld + path + "#" + fragment
)

var (
	a, err = address.NewAddress(rawurl)
)

func Test(t *testing.T) {
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestDomain(t *testing.T) {
	if a.Domain != domain {
		t.Fatalf(templateErr, "Domain", a.Domain, domain)
	}
}

func TestFragment(t *testing.T) {
	if a.Fragment != fragment {
		t.Fatalf(templateErr, "Fragment", a.Fragment, fragment)
	}
}

func TestHost(t *testing.T) {
	if a.Host != host {
		t.Fatalf(templateErr, "Host", a.Host, host)
	}
}

func TestHostname(t *testing.T) {
	if a.Hostname != host {
		t.Fatalf(templateErr, "Hostname", a.Hostname, hostname)
	}
}

func TestPort(t *testing.T) {
	if a.Path != path {
		t.Fatalf(templateErr, "Path", a.Path, path)
	}
}

func TestScheme(t *testing.T) {
	if a.Scheme != scheme {
		t.Fatalf(templateErr, "Scheme", a.Scheme, scheme)
	}
}

func TestSubdomain(t *testing.T) {
	if a.Subdomain != subdomain {
		t.Fatalf(templateErr, "Subdomain", a.Subdomain, subdomain)
	}
}

func TestTLD(t *testing.T) {
	if a.TLD != tld {
		t.Fatalf(templateErr, "TLD", a.TLD, tld)
	}
}

func TestURL(t *testing.T) {
	if a.URL != rawurl {
		t.Fatalf(templateErr, "URL", a.URL, rawurl)
	}
}
