package html

import (
	"html"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

const (
	noa string = "noa:publisher/" // prefix for Nintendo of America.
)
const (
	rep string = " " // rep string
)

var (
	regexpHTML    = regexp.MustCompile(`(<[^>]*>|\n(\s{1,})?)`)  // match all HTML tokens.
	regexpHyphens = regexp.MustCompile(`\-{2,}`)                 // match all repeating hyphens.
	regexpName    = regexp.MustCompile(`(\&\#[0-9]+\;|™|\(|\))`) // match all unwanted characters for a Amiibo name.
	regexpSpaces  = regexp.MustCompile(`\s{2,}`)                 // match all repeating spaces.
	regexpURI     = regexp.MustCompile(`[^a-zA-Z0-9&]+`)         // match all unwanted characters in a URI.
)
var (
	replacerURI = strings.NewReplacer([]string{"&", "and", "'", "", "é", "e"}...) // replacer for names like Pokémon.
)

var (
	transformer = transform.Chain(norm.NFD, transform.RemoveFunc(func(r rune) bool { return unicode.Is(unicode.Mn, r) }), norm.NFC)
)

// PresentedBy removes all unwanted substrings from a presented by string.
func PresentedBy(s string) string {
	return strings.TrimPrefix(s, noa)
}

// Name removes all unwanted characters to create a name.
func Name(s string) string {
	return regexpName.ReplaceAllString(s, "")
}

// Untokenize removes all HTML tokens from a string.
func Untokenize(s string) string {
	s = regexpSpaces.ReplaceAllString(regexpHTML.ReplaceAllString(s, rep), rep)
	s = html.UnescapeString(strings.TrimSpace(s))
	return s
}

// URI formats a string to be an expected REST URI.
func URI(s string) string {
	s = replacerURI.Replace(s)
	s = regexpURI.ReplaceAllString(s, "-")
	s = regexpHyphens.ReplaceAllString(s, "")
	s = strings.TrimSuffix(s, "-")
	s = strings.ToLower(s)
	return s
}
