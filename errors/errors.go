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

var (
	ErrSEmpty = fmt.Errorf("*s is empty")
)
var (
	ErrSNoHref = fmt.Errorf("*s has no href")
)
var (
	ErrSNoSrc = fmt.Errorf("*s has no src")
)
var (
	ErrSNoText = fmt.Errorf("*s has no text")
)
var (
	ErrSNoTitle = fmt.Errorf("*s has no title")
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
