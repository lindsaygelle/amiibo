package resource_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gellel/amiibo/resource"
)

const (
	templateErr string = "resource: resource.(%s) %s %s %s"
)

func r(name, rawurl string, t *testing.T) {
	const (
		template string = "http.StatusCode (%d)"
	)
	var (
		res, err = http.Get(rawurl)
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Fatalf(templateErr, name, fmt.Sprintf(template, res.StatusCode), "!=", fmt.Sprintf(template, http.StatusOK))
	}
}

func TestAmiibo(t *testing.T) {
	const (
		name string = "Amiibo"
	)
	r(name, resource.Amiibo, t)
}

func TestCompatability(t *testing.T) {
	const (
		name string = "Compatability"
	)
	r(name, resource.Compatability, t)
}

func TestGame(t *testing.T) {
	const (
		name string = "Game"
	)
	r(name, resource.Game, t)
}

func TestLineup(t *testing.T) {
	const (
		name string = "Lineup"
	)
	r(name, resource.Lineup, t)
}
