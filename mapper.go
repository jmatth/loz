//go:generate go run ./internal/gen/gen.go map.g
package loz

// Mapper1.Fold is identical to [Seq.Fold] except that the type of the result can
// be different than than the type of the elements in the sequence.
func (m Mapper1[T, O]) Fold(initial O, combine reducer[T, O]) O {
	for v := range m {
		initial = combine(initial, v)
	}
	return initial
}
