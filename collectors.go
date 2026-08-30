package loz

// Collector is the interface accepted by [Seq.Collect], used to collect an
// iterator into a single value.
type Collector[R, V any] interface {
	Initial() R
	Collect(R, V) R
}

// KVCollector is the interface accepted by [KVSeq.Collect], used to collect an
// iterator of key/value pairs into a single value.
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

// ToSlice collects a [Seq] into a slice.
func ToSlice[V any]() Collector[[]V, V] {
	return toSliceCollector[[]V, V]{}
}

// ToSliceAppend collects a [Seq] into a slice by appending its contents to
// `initial`.
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

// ToMap creates a [KVCollector] that collects an iterator of key/value pairs
// into a map. In the case of multiple keys, later keys in the sequence will
// take precedence.
func ToMap[K comparable, V any]() KVCollector[map[K]V, K, V] {
	return mapCollector[map[K]V, K, V]{}
}

// ToMapMerge behaves the same as [ToMap] but inserts they key/value pairs into
// the map provided by `initial` rather than creating a new map.
func ToMapMerge[K comparable, V any](initial map[K]V) KVCollector[map[K]V, K, V] {
	return mapCollector[map[K]V, K, V]{
		startingMap: initial,
	}
}

type mergingMapCollector[K comparable, V, OV any, R map[K]OV] struct {
	merge func(OV, V) OV
}

// Collect implements [KVCollector].
func (g mergingMapCollector[K, V, _, R]) Collect(r R, k K, v V) R {
	r[k] = g.merge(r[k], v)
	return r
}

// Initial implements [KVCollector].
func (g mergingMapCollector[K, V, _, R]) Initial() R {
	return make(R)
}

func mergeToList[V any](v1 []V, v2 V) []V {
	return append(v1, v2)
}

// ToMapGroup creates a [KVCollector] that collects an iterator of key/value
// pairs into a `map[K][]V`. Values that share a key are appended to a list
// under that key in the final map.
func ToMapGroup[K comparable, V any]() KVCollector[map[K][]V, K, V] {
	return mergingMapCollector[K, V, []V, map[K][]V]{
		merge: mergeToList,
	}
}
