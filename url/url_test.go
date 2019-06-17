package url_test

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/gellel/amiibo/url"
)

func namespace(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func TestNewURL(t *testing.T) {
	var (
		name string = namespace(TestNewURL)
	)
	if u = url.NewURL(URL); reflect.ValueOf(u).IsNil() {
		var (
			err = fmt.Sprintf(ERROR_FUNCTION, name, "t.Fatal", "url.NewURL(URL)", u, "*url.URL")
		)
		t.Fatalf(err)
	}
}

func TestURLHost(t *testing.T) {
	var (
		name string = namespace(TestURLHost)
	)
	if host := u.Host(); host != "www.google.com" {
		var (
			err = fmt.Sprintf(ERROR_FUNCTION, name, "t.Fatal", "URL.Host", host, "www.google.com")
		)
		t.Fatalf(err)
	}
	if host := strings.TrimPrefix(u.Host(), "www."); host != "google.com" {
		var (
			err = fmt.Sprintf(ERROR_FUNCTION, name, "t.Error", "URL.Host", host, "google.com")
		)
		t.Errorf(err)
	}
}

func TestURLHTTP(t *testing.T) {
	var (
		name string = namespace(TestURLHTTP)
	)
	if ok := u.HTTP(); ok != false {
		var (
			err = fmt.Sprintf(ERROR_FUNCTION, name, "t.Fatal", "URL.HTTP", ok, false)
		)
		t.Fatalf(err)
	}
}
func TestURLSSL(t *testing.T) {
	const (
		expect bool = false
	)
	var (
		name string = namespace(TestURLSSL)
	)
	if ok := u.HTTPS(); ok == expect {
		var (
			err = fmt.Sprintf(ERROR_FUNCTION, name, "t.Fatal", "URL.HTTPS", ok, expect)
		)
		t.Fatalf(err)
	}
}

func TestURLPath(t *testing.T) {
	const (
		expect string = "/dir/1/2/"
	)
	var (
		name string = namespace(TestURLPath)
	)
	if path := u.Path(); path != expect {
		var (
			err = fmt.Sprintf(ERROR_FUNCTION, name, "t.Fatal", "URL.Path", path, expect)
		)
		t.Fatalf(err)
	}
}
