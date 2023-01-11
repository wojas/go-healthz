package healthz

import (
	"errors"
	"fmt"
)

// Warning is a type of error that is considered a warning instead of a failure.
// It will not cause the health check to fail, but the warning will appear
// in the JSON.
type Warning struct {
	Err error
}

func (w Warning) Error() string {
	return w.Err.Error()
}

func (w Warning) Unwrap() error {
	return w.Err
}

// Warn returns a Warning with given message
func Warn(msg string) error {
	return Warning{errors.New(msg)}
}

// Warnf formats a Warning
func Warnf(format string, args ...interface{}) error {
	return Warning{fmt.Errorf(format, args...)}
}

// IsWarning returns true if the error is a Warning instead of a failure.
// It will recursively unwrap the error to look for a Warning.
func IsWarning(err error) bool {
	for err != nil {
		if _, ok := err.(Warning); ok {
			return true
		}
		err = errors.Unwrap(err)
	}
	return false
}
