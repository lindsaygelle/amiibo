# Amiibo

Go package. Handles the collection, parsing and normalization of the Amiibo JSON content found on the Nintendo Amiibo URI. Parses the raw response into a set of normalized fields that can be wrapped by any Go program. Offers two variations of the Amiibo data - the `amiibo.RawAmiibo` and the `amiibo.Amiibo`. The raw variation represents the fluid-state endpoint response that is provided as-is from the Nintendo URI and is subject to change at any time. The normalized struct attempts to handle the potential volitility.


Get it:

```
go get github.com/gellel/amiibo
```

Import it:

```
import (
	"github.com/gellel/amiibo"
)
```

## Usage

Getting the raw Amiibo payload from the Nintendo URI.

```go
package main

import (
	"fmt"

	"github.com/gellel/amiibo"
)

func main() {

    rawPayload, err := amiibo.GetRawPayload()

    if err != nil {
        panic(err) // something went wrong fetching the API response.
    } else {
        fmt.Println(rawPayload)
    }
}
```

Getting the normalized Amiibo from the Nintendo URI.

```go
package main

import (
    "fmt"

    "github.com/gellel/amiibo"
)

func main() {

    if payload, err := amiibo.GetPayload(); err == nil {
        fmt.Println(payload.Amiibo)
    }
}
```

## License

[MIT](https://github.com/gellel/slice/blob/master/LICENSE)
