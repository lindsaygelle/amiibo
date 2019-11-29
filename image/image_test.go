package image_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gellel/amiibo/image"
)

const (
	dir       string = "screenshot-gallery"
	domain    string = "nintendo"
	ext       string = "png"
	host      string = subdomain + "." + domain + "." + tld
	hostname  string = host + ":" + port
	name      string = "amiibo_Alm_FireEmblem_char"
	path      string = "content/dam/noa/en_US/amiibo/alm-amiibo-fire-emblem-series"
	port      string = "443"
	scheme    string = "http"
	subdomain string = "www"
	tld       string = "com"
)

const (
	rawurl string = scheme + "://" + subdomain + "." + domain + "." + tld + "/" + path + "/" + dir + "/" + name + "." + ext
)

const (
	templateErr string = "image.Image.%s: i.(%s) %s %s"
)

var (
	i, err = image.NewImage(rawurl)
)

func Test(t *testing.T) {
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestDir(t *testing.T) {
	if i.Dir != dir {
		t.Fatalf(templateErr, "Dir", i.Dir, "!=", dir)
	}
}

func TestExt(t *testing.T) {
	if i.Dir != dir {
		t.Fatalf(templateErr, "Ext", i.Ext, "!=", ext)
	}
}

func TestHeight(t *testing.T) {
	if i.Height == -1 {
		t.Fatalf(templateErr, "Height", fmt.Sprintf("%d", i.Height), "==", fmt.Sprintf("%d", -1))
	}
}

func TestMeasurement(t *testing.T) {
	if i.Measurement != "px" {
		t.Fatalf(templateErr, "Measurement", i.Measurement, "!=", "px")
	}
}

func TestName(t *testing.T) {
	if i.Name != name {
		t.Fatalf(templateErr, "Name", i.Name, "!=", name)
	}
}

func TestStatusCode(t *testing.T) {
	if i.StatusCode != (http.StatusOK) {
		t.Fatalf(templateErr, "StatusCode", fmt.Sprintf("%d", i.StatusCode), "!=", fmt.Sprintf("%d", http.StatusOK))
	}
}

func TestWidth(t *testing.T) {
	if i.Width == -1 {
		t.Fatalf(templateErr, "Width", fmt.Sprintf("%d", i.Width), "==", fmt.Sprintf("%d", -1))
	}
}
