package amiibo

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// NintendoURL is the URL to the Nintendo of America website.
const NintendoURL string = "https://nintendo.com"

// NintendoURLJPN is the URL to the Nintendo of Japan website.
const NintendoURLJPN string = "https://www.nintendo.co.jp"

const (
	noa string = "noa:publisher/"
)

var (
	regexpHTML       = regexp.MustCompile(`(<[^>]*>|\n(\s{1,})?)`)  // match all HTML tokens.
	regexpHyphens    = regexp.MustCompile(`\-{2,}`)                 // match all repeating hyphens.
	regexpName       = regexp.MustCompile(`(\&\#[0-9]+\;|™|\(|\))`) // match all unwanted characters.
	regexPunctuation = regexp.MustCompile(`[\.\,\:\']`)             // match all punctuation.
	regexpSpaces     = regexp.MustCompile(`\s{2,}`)                 // match all repeating spaces.
	regexpURI        = regexp.MustCompile(`[^a-zA-Z0-9&]+`)         // match all unwanted characters in a URI.
)
var (
	replacerURI = strings.NewReplacer([]string{"&", "and", "'", "", "é", "e"}...) // replacer for names like Pokémon.
)

var (
	// transform formats unicode.
	transformer = transform.Chain(norm.NFD, transform.RemoveFunc(func(r rune) bool { return unicode.Is(unicode.Mn, r) }), norm.NFC)
)

// getRemoteFile gets a remote file from a remote URL.
func getRemoteFile(URL string) (req *http.Request, res *http.Response, err error) {
	if len(URL) == 0 {
		err = fmt.Errorf("http: URL")
	}
	if err != nil {
		return
	}
	req, err = http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return
	}
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf(("http: %d"), res.StatusCode)
	}
	if err != nil {
		return
	}
	return
}

// getRemoteImage gets a remote image file from a remote URL.
func getRemoteImage(URL string) (i image.Image, err error) {
	var res *http.Response
	_, res, err = getRemoteFile(URL)
	if err != nil {
		return
	}
	var e = strings.TrimSuffix(URL, filepath.Ext(URL))
	e = strings.ToUpper(e)
	var fn func(io.Reader) (image.Image, error)
	switch e {
	case "GIF":
		fn = gif.Decode
	case "JPEG":
		fn = jpeg.Decode
	case "PNG":
		fn = png.Decode
	}
	if !reflect.ValueOf(fn).IsZero() {
		i, err = fn(res.Body)
		return
	}
	i, _, err = image.Decode(res.Body)
	return
}

// marshal handles a (package).Marshal operation.
func marshal(v interface{}, fn func(interface{}) ([]byte, error)) (b []byte, err error) {
	b, err = fn(v)
	return
}

// readFile handles a (package).ReadFile operation.
func readFile(dir, filename string, fn func(string) ([]byte, error)) (b []byte, err error) {
	b, err = fn(filepath.Join(dir, filename))
	return
}

// readJSONFile reads a JSON file from disc and unmarshals its contents using json.Unmarshal.
func readJSONFile(dir, filename string, v interface{}) (err error) {
	err = unmarshal(dir, filename, &v, json.Unmarshal)
	return
}

// readXMLFile reads a XML file from disc and unmarshals its contents using xml.Unmarshal.
func readXMLFile(dir, filename string, v interface{}) (err error) {
	err = unmarshal(dir, filename, &v, xml.Unmarshal)
	return
}

// unmarshal handles a (package).Unmarshal operation.
func unmarshal(dir, filename string, v interface{}, fn func([]byte, interface{}) error) (err error) {
	var b ([]byte)
	b, err = readFile(dir, filename, ioutil.ReadFile)
	if err != nil {
		return
	}
	err = fn(b, v)
	if err != nil {
		return
	}
	return
}

// writeFile writes a file to disc using ioutil.WriteFile.
func writeFile(dir, filename string, b []byte) (fullpath string, err error) {
	fullpath = filepath.Join(dir, filename)
	err = ioutil.WriteFile(fullpath, b, 0644)
	return
}

// writeFileJSON writes a JSON file to disc.
func writeJSONFile(dir, filename string, v interface{}) (fullpath string, err error) {
	var b ([]byte)
	b, err = marshal(&v, json.Marshal)
	if err != nil {
		return
	}
	fullpath, err = writeFile(dir, filename, b)
	return
}

// writeXMLFile writes a XML fille to disc.
func writeXMLFile(dir, filename string, v interface{}) (fullpath string, err error) {
	var b ([]byte)
	b, err = marshal(&v, xml.Marshal)
	if err != nil {
		return
	}
	fullpath, err = writeFile(dir, filename, b)
	return
}
