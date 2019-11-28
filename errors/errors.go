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
	//ErrGoQueryEmpty is an error for an empty goquery.Selection.
	ErrGoQueryEmpty = fmt.Errorf(errGoQueryEmpty)
)
var (
	// ErrGoQueryNoHref is an error for a goquery.Selection without a href attribute.
	ErrGoQueryNoHref = fmt.Errorf(errGoQueryNoHref)
)
var (
	// ErrGoQueryNoSrc is an error for a goquery.Selection without a src attribute.
	ErrGoQueryNoSrc = fmt.Errorf(errGoQueryNoSrc)
)
var (
	// ErrGoQueryNoText is an error for a goquery.Selection without text nodes.
	ErrGoQueryNoText = fmt.Errorf(errGoQueryNoText)
)
var (
	// ErrGoQueryNoTitle is an error for a goquery.Selection without a title attribute.
	ErrGoQueryNoTitle = fmt.Errorf(errGoQueryNoTitle)
)

var (
	//ErrArgCNil is an error when pointer to argument c is nil.
	ErrArgCNil = fmt.Errorf(errArgCNil)
)
var (
	//ErrArgGNil is an error when pointer to argument g is nil.
	ErrArgGNil = fmt.Errorf(errArgGNil)
)
var (
	//ErrArgINil is an error when pointer to argument i is nil.
	ErrArgINil = fmt.Errorf(errArgINil)
)
var (
	//ErrArgLNil is an error when pointer to argument l is nil.
	ErrArgLNil = fmt.Errorf(errArgLNil)
)
var (
	//ErrArgSNil is an error when pointer to argument s is nil.
	ErrArgSNil = fmt.Errorf(errArgSNil)
)
var (
	//ErrArgsNil is an error when all arguments that are pointers are nil.
	ErrArgsNil = fmt.Errorf(errArgsNil)
)
var (
	// ErrArgsNoRel is an error when all arguments (usually to mix) have no common struct value data point.
	ErrArgsNoRel = fmt.Errorf(errArgsNoRel)
)
