package loz

type Collector[R, V any] interface {
	Initial() R
	Collect(R, V) R
}

type KVCollector[R, K, V any] interface {
	Initial() R
	Collect(R, K, V) R
}

type toSliceCollector[R []V, V any] struct {
	startingSlice R
}

// Collect implements [Collector].
func (t toSliceCollector[R, V]) Collect(acc R, iter V) R {
	return append(acc, iter)
}

// Initial implements [Collector].
func (t toSliceCollector[R, V]) Initial() R {
	return t.startingSlice
}

func ToSlice[V any]() Collector[[]V, V] {
	return toSliceCollector[[]V, V]{}
}

func ToSliceAppend[V any](initial []V) Collector[[]V, V] {
	return toSliceCollector[[]V, V]{
		startingSlice: initial,
	}
}

type mapCollector[R map[K]V, K comparable, V any] struct {
	startingMap R
}

// Collect implements [KVCollector].
func (m mapCollector[R, K, V]) Collect(acc R, key K, val V) R {
	acc[key] = val
	return acc
}

// Initial implements [KVCollector].
func (m mapCollector[R, K, V]) Initial() R {
	if m.startingMap != nil {
		return m.startingMap
	}
	return make(R)
}

func ToMap[K comparable, V any]() KVCollector[map[K]V, K, V] {
	return mapCollector[map[K]V, K, V]{}
}

// // ToMap collects a [KVSeq] to a `map[K]V` when passed to [KVSeq.Collect].
// // WARNING: when collecting an empty iterator, a nil map will be returned.
// func ToMap[K comparable, V any](acc map[K]V, iterK K, iterV V) map[K]V {
// 	if acc == nil {
// 		acc = map[K]V{}
// 	}
// 	acc[iterK] = iterV
// 	return acc
// }
