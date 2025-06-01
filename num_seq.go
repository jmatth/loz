//go:generate go run ./internal/agen/agen.go seq Seq NumSeq num_seq.g.go
//go:generate go run ./internal/agen/agen.go --a ord_seq OrdSeq NumSeq num_seq.g.go
package loz

import (
	"errors"

	"golang.org/x/exp/constraints"
)

type number interface {
	constraints.Integer | constraints.Float
}

type NumSeq[V number] Seq[V]

func Range(to int) NumSeq[int] {
	return func(yield func(int) bool) {
		for i := range to {
			if !yield(i) {
				break
			}
		}
	}
}

func RangeFrom(from, to int) NumSeq[int] {
	return func(yield func(int) bool) {
		for i := from; i < to; i++ {
			if !yield(i) {
				break
			}
		}
	}
}

func RangeInterval[V number](from, to, interval V) NumSeq[V] {
	return func(yield func(V) bool) {
		if interval == 0 || (to > from && interval <= 0) || (to < from && interval >= 0) {
			return
		}
		var shouldContinue func(V, V) bool
		if interval > 0 {
			shouldContinue = func(i, to V) bool {
				return i < to
			}
		} else {
			shouldContinue = func(i, to V) bool {
				return i > to
			}
		}
		for i := from; shouldContinue(i, to); i += interval {
			if !yield(i) {
				break
			}
		}
	}
}

// Sum returns the sum of all the values in the iterator. It returns 0 if the
// iterator is empty.
func (s NumSeq[V]) Sum() V {
	sum, err := Seq[V](s).Reduce(func(a, b V) V { return a + b })
	if err != nil {
		if errors.Is(err, EmptySeqErr) {
			return 0
		}
		panic(err)
	}
	return sum
}
