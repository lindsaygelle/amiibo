package url_test

import (
	"fmt"
	"testing"

	"github.com/gellel/amiibo/url"
)

func Test(t *testing.T) {

	URL := url.NewURL("https://www.google.com/dir/1/2/search.html?arg=0-a&arg1=1-b&arg3-c#hash")

	for i, k := range URL.Unpack() {
		fmt.Println(i, k)
	}

	fmt.Println(URL.Keyset())
}
