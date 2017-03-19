package comb

import "fmt"

// Result represents the result of a parser.
// It supports a range of common values, including a rune slice,
// integer and float types used in strconv, as well as an interface{}
// for anything not included. Err will be set if a Result is failed.
// If your result contains an error that is not a failure, then it should
// be placed into Interface.
type Result struct {
	Err       error
	Runes     []rune
	Int64     int64
	Float64   float64
	Interface interface{}
}

// Matched returns true if Err is not nil.
func (r Result) Matched() bool {
	return r.Err == nil
}

// Failed returns a failed result with a given error.
func Failed(err error) Result {
	return Result{Err: err}
}

type errWrap struct {
	format string
	a      []interface{}
	err    error
}

func (e errWrap) Error() string {
	if e.err != nil {
		return e.err.Error()
	}

	e.err = fmt.Errorf(e.format, e.a...)
	return e.err.Error()
}

// Failedf returns a failed result in fmt.Errorf form.
// fmt.Errorf will not be called until the error is read to prevent
// unneccesary computation. This is important, as failed results
// can be checked without ever generating an error. After calling
// Error() on the error, the error is cached for further use.
func Failedf(format string, a ...interface{}) Result {
	return Result{
		Err: errWrap{
			format: format,
			a:      a,
		},
	}
}
