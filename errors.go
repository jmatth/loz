package loz

import (
	"fmt"
)

type SeqError int

const (
	EmptySeqErr SeqError = iota
)

type WrappedSeqError struct {
	wrapped error
}

func (e WrappedSeqError) Error() string {
	return fmt.Sprintf("error during iteration: %v", e.wrapped)
}

func (e WrappedSeqError) Unwrap() error {
	return e.wrapped
}

func (e SeqError) Error() string {
	switch e {
	case EmptySeqErr:
		return "empty iterator"
	}
	return "unknown iteration error"
}

func recoverSeqError(err any) error {
	if err, ok := err.(WrappedSeqError); ok {
		return err.wrapped
	}
	panic(err)
}

// PanicHaltIteration causes any iteration to end early by wrapping the
// provided error and passing it to panic. How this error is treated depends on
// the consuming method called on the iterator. If the consuming method returns
// an error, then it will return the error passed to this function along with
// zero values for any of its other returns. If it does not return an error
// then it will panic with a [WrappedSeqError], and any recovering code can
// access the original error by calling [WrappedSeqError.Unwrap].
func PanicHaltIteration(err error) {
	panic(WrappedSeqError{wrapped: err})
}