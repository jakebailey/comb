package comb

import "fmt"

type errorFunc func() string

func (e errorFunc) Error() string {
	return e()
}

func errorf(format string, a ...interface{}) error {
	return errorFunc(func() string {
		return fmt.Sprintf(format, a...)
	})
}
