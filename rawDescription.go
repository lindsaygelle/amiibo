package amiibo

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"golang.org/x/net/html"
)

var (
	_ rawAmiiboDescription = (*RawAmiiboDescription)(nil)
)

type rawAmiiboDescription interface {
	parseGoQuery() *goquery.Document
	parseHTMLNode() (*html.Node, error)
	parseHTMLTextNode() string
	String() string
}

// A RawAmiiboDescription string represents the HTML content found in the overviewDescription property
// that sits inside a RawAmiibo within in the Nintendo XHR HTTP response.
type RawAmiiboDescription string

func (r *RawAmiiboDescription) parseHTMLNode() (*html.Node, error) {
	return html.Parse(strings.NewReader(string(*r)))
}

func (r *RawAmiiboDescription) parseHTMLTextNode() string {
	return r.parseGoQuery().Find("p").First().Text()
}

func (r *RawAmiiboDescription) parseGoQuery() *goquery.Document {
	root, err := r.parseHTMLNode()
	if err != nil {
		panic(err)
	}
	return goquery.NewDocumentFromNode(root)
}

func (r *RawAmiiboDescription) strip() string {
	s := regexp.MustCompile(`(\n\t)`).ReplaceAllString(r.parseHTMLTextNode(), "")
	return strings.TrimSpace(regexp.MustCompile(`\s{1,}`).ReplaceAllString(s, " "))
}

func (r *RawAmiiboDescription) String() string {
	return fmt.Sprintf("%s", r.strip())
}
