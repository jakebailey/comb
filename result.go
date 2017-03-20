package comb

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
	Tag       string
	Ignore    bool
}

// Matched returns true if Err is not nil.
func (r Result) Matched() bool {
	return r.Err == nil
}

// Failed returns a failed result with a given error.
func Failed(err error) Result {
	return Result{Err: err}
}

// Failedf returns a failed result in fmt.Errorf form.
// fmt.Errorf will not be called until the error is read to prevent
// unnecessary computation. This is important, as failed results
// can be checked without ever generating an error.
func Failedf(format string, a ...interface{}) Result {
	return Result{
		Err: errorf(format, a...),
	}
}
