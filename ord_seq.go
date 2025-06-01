//go:generate go run ./internal/agen/agen.go seq Seq OrdSeq ord_seq.g.go
package loz

import (
	"cmp"
)

type OrdSeq[V cmp.Ordered] Seq[V]

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