package internal

import "fmt"

type WrappedSeqError struct {
	wrapped error
}

func (e WrappedSeqError) Error() string {
	return fmt.Sprintf("error during iteration: %v", e.wrapped)
}

func (e WrappedSeqError) Unwrap() error {
	return e.wrapped
}

func NewWrappedSeqError(err error) error {
	if err == nil {
		return nil
	}
	return WrappedSeqError{wrapped: err}
}

func RecoverHaltIteration(err *error) {
	if r := recover(); r != nil {
		if r, ok := r.(WrappedSeqError); ok {
			*err = r.wrapped
			return
		}
		panic(r)
	}
}
