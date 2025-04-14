//go:generate go run ./internal/gen/gen.go map_generated
package loz

type IndexedVal[T any] struct {
	Index uint
	Val T
}
