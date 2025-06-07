//go:generate go run ./internal/map_gen/map_gen.go map.g
package loz

// Fold is identical to [Seq.Fold] except that the type of the result can
// be different than than the type of the elements in the sequence.
func (m Map1[T, O]) Fold(initial O, combine reducer[T, O]) O {
	for v := range m {
		initial = combine(initial, v)
	}
	return initial
}

func (m Map1[T, O]) TryFold(initial O, combine reducer[T, O]) (result O, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = recoverSeqError(r)
		}
	}()
	return m.Fold(initial, combine), nil
}
