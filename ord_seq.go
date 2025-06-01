//go:generate go run ./internal/agen/agen.go seq Seq OrdSeq ord_seq.g.go
package loz

import (
	"cmp"
)

// OrdSeq is a [Seq] where the elements satisfy [cmp.Ordered].
type OrdSeq[V cmp.Ordered] Seq[V]

// Max consumes the iterator and returns the largest element, as determined by
// the > operator. If the iterator is empty then a zero value is returned along
// with an error.
func (s OrdSeq[V]) Max() (V, error) {
	var result *V
	for v := range s {
		if result == nil {
			result = &v
			continue
		}
		if v > *result {
			result = &v
		}
	}
	if result == nil {
		var zero V
		return zero, EmptySeqErr
	}
	return *result, nil
}

// Min consumes the iterator and returns the smallest element, as determined by
// the < operator. If the iterator is empty then a zero value is returned along
// with an error.
func (s OrdSeq[V]) Min() (V, error) {
	var result *V
	for v := range s {
		if result == nil {
			result = &v
			continue
		}
		if v < *result {
			result = &v
		}
	}
	if result == nil {
		var zero V
		return zero, EmptySeqErr
	}
	return *result, nil
}
