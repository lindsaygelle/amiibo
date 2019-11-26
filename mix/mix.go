package mix

import "fmt"

const (
	tep string = "*%s is %s"
)

var (
	errCNil = fmt.Errorf(tep, "c", "nil")
	errGNil = fmt.Errorf(tep, "g", "nil")
	errINil = fmt.Errorf(tep, "i", "nil")
	errLNil = fmt.Errorf(tep, "l", "nil")
)
