package errors

import "fmt"

const (
	errArgsNil string = "all arguments are nil"
)
const (
	errArgsNoRel string = "all arguments do not have a common relationship"
)
const (
	errArgCNil string = "*c is nil"
	errArgGNil string = "*g is nil"
	errArgINil string = "*i is nil"
	errArgLNil string = "*l is nil"
	errArgSNil string = "*s is nil"
)

const (
	errGoQuerySelection string = "*goquery.Selection (*s)"
	errGoQueryEmpty     string = errGoQuerySelection + " " + "is empty"
	errGoQueryNoHref    string = errGoQuerySelection + " " + "has no attribute href"
	errGoQueryNoSrc     string = errGoQuerySelection + " " + "has no attribute src"
	errGoQueryNoText    string = errGoQuerySelection + " " + "has no text nodes"
	errGoQueryNoTitle   string = errGoQuerySelection + " " + "has no attribute title"
)

var (
	ErrGoQueryEmpty = fmt.Errorf(errGoQueryEmpty)
)
var (
	ErrGoQueryNoHref = fmt.Errorf(errGoQueryNoHref)
)
var (
	ErrGoQueryNoSrc = fmt.Errorf(errGoQueryNoSrc)
)
var (
	ErrGoQueryNoText = fmt.Errorf(errGoQueryNoText)
)
var (
	ErrGoQueryNoTitle = fmt.Errorf(errGoQueryNoTitle)
)

var (
	ErrArgCNil = fmt.Errorf(errArgCNil)
)
var (
	ErrArgGNil = fmt.Errorf(errArgGNil)
)
var (
	ErrArgINil = fmt.Errorf(errArgINil)
)
var (
	ErrArgLNil = fmt.Errorf(errArgLNil)
)
var (
	ErrArgSNil = fmt.Errorf(errArgSNil)
)
var (
	ErrArgsNil = fmt.Errorf(errArgsNil)
)
var (
	ErrArgsNoRel = fmt.Errorf(errArgsNoRel)
)
