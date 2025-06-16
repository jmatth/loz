package loz

import (
	"fmt"
)

type SeqError int

const (
	EmptySeqErr SeqError = iota
)

type wrappedSeqError struct {
	wrapped error
}

func (e wrappedSeqError) Error() string {
	return fmt.Sprintf("error during iteration: %v", e.wrapped)
}

func (e wrappedSeqError) Unwrap() error {
	return e.wrapped
}

func (e SeqError) Error() string {
	switch e {
	case EmptySeqErr:
		return "empty iterator"
	}
	return "unknown iteration error"
}

func recoverHaltIteration(err *error) {
	if r := recover(); r != nil {
		if r, ok := r.(wrappedSeqError); ok {
			*err = r.wrapped
			return
		}
		panic(r)
	}
}

// PanicHaltIteration causes any iteration to end early by wrapping the
// provided error and panicking. To easily recover from this panic and return
// the error normally, use a terminal method prefixed with "Try", such as
// [Seq.TryCollectSlice]. These methods automatically recover from panics
// caused by this function and return the wrapped value as their final return
// value. Calling this method with nil is a noop.
func PanicHaltIteration(err error) {
	if err == nil {
		return
	}
	panic(wrappedSeqError{wrapped: err})
}
