// Package mapping provides additional wrapper types for loz.Seq and loz.KVSeq
// so that methods can change the type of the elements in the iterator. This
// package is almost entirely generated. Refer to the documentation for
// [github.com/jmatth/loz] for usage examples.
//
//go:generate go run ../internal/map_gen/map_gen.go ../ map.g
package mapping

import (
	. "github.com/jmatth/loz/internal"
)

// Fold is identical to [Seq.Fold] except that the type of the result can
// be different than than the type of the elements in the sequence.
func (m Map1[T, O]) Fold(initial O, combine Reducer[T, O]) O {
	for v := range m {
		initial = combine(initial, v)
	}
	return initial
}

func (m Map1[T, O]) TryFold(initial O, combine Reducer[T, O]) (result O, err error) {
	defer RecoverHaltIteration(&err)
	return m.Fold(initial, combine), nil
}
