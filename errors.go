package loz

import . "github.com/jmatth/loz/internal"

type SeqError int

const (
	EmptySeqErr SeqError = iota
)

func (e SeqError) Error() string {
	switch e {
	case EmptySeqErr:
		return "empty iterator"
	}
	return "unknown iteration error"
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
	panic(NewWrappedSeqError(err))
}
