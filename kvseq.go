package loz

import (
	"iter"
	"maps"
)

// KVSeq is an alias to [iter.Seq2] that provides additional methods for
// filtering, transforming, and collecting the elements. Though the name and
// several doc comments imply it contains key/value pairs, that is only the most
// common use case and the relationship between the two values is arbitrary.
type KVSeq[K, V any] iter.Seq2[K, V]

// IterMap creates a Seq over the key/value pairs of a map.
func IterMap[K comparable, V any](input map[K]V) KVSeq[K, V] {
	return KVSeq[K, V](maps.All(input))
}

// ToKeys converts a KVSeq[K, V] to a Seq[K], continuing the iteration with only
// the keys.
func (s KVSeq[K, V]) Keys() Seq[K] {
	return func(yield yielder[K]) {
		s(func(k K, _ V) bool {
			return yield(k)
		})
	}
}

// Values converts a KVSeq[K, V] to a Seq[V], continuing the iteration with only
// the values.
func (s KVSeq[K, V]) Values() Seq[V] {
	return func(yield yielder[V]) {
		s(func(_ K, v V) bool {
			return yield(v)
		})
	}
}

// ForEach consumes the iterator and calls the provided function with each of
// the key/value pairs.
func (s KVSeq[K, V]) ForEach(process func(K, V)) {
	s(func(k K, v V) bool {
		process(k, v)
		return true
	})
}

// TryForEach is identical to [KVSeq.ForEach], except it will recover any panic
// caused by [PanicHaltIteration] and return the wrapped error.
func (s KVSeq[K, V]) TryForEach(process func(K, V)) (err error) {
	defer recoverHaltIteration(&err)
	s.ForEach(process)
	return nil
}

// Map transforms the key/value pairs within the iterator using the provided
// mapper function. Due to limitations of the Go type system, the mapped keys
// and values must be the same types as the input. To perform mapping
// operations that change types, see [KVMap1], [KVMap2], etc.
func (s KVSeq[K, V]) Map(mapper func(K, V) (K, V)) KVSeq[K, V] {
	return func(yield yielder2[K, V]) {
		s(func(k K, v V) bool {
			return yield(mapper(k, v))
		})
	}
}

// FilterMap is a combination of [KVSeq.Filter] and [KVSeq.Map]. If the
// provided mapper function returns an error, then the current key/value pair
// of the iteration will be skipped. If no error is returned, then the mapped
// key/value pair is passed to the next iteration stage as normal.
func (s KVSeq[K, V]) FilterMap(mapper func(K, V) (K, V, error)) KVSeq[K, V] {
	return func(yield yielder2[K, V]) {
		s(func(k K, v V) bool {
			mk, mv, err := mapper(k, v)
			if err != nil {
				return true
			}
			return yield(mk, mv)
		})
	}
}

// Reduce reduces the iterator to a single key/value pair by iteratively
// combining its elements using the provided function. If the iterator is empty
// then zero values will be returned along with an error.
func (s KVSeq[K, V]) Reduce(combine reducer2[K, V]) (K, V, error) {
	var keyResult K
	var valResult V
	isFirst := true
	s(func(k K, v V) bool {
		if isFirst {
			keyResult = k
			valResult = v
			isFirst = false
			return true
		}
		keyResult, valResult = combine(keyResult, valResult, k, v)
		return true
	})
	if isFirst {
		return keyResult, valResult, EmptySeqErr
	}
	return keyResult, valResult, nil
}

// TryReduce is identical to [KVSeq.ForReduce], except it will recover any
// panic caused by [PanicHaltIteration] and return the wrapped error.
func (s KVSeq[K, V]) TryReduce(combine reducer2[K, V]) (_ K, _ V, err error) {
	defer recoverHaltIteration(&err)
	return s.Reduce(combine)
}

// Fold reduces the iterator to a single key/value pair by iteratively
// combining its elements with initial values using the provided function. If
// the iterator is empty the initial values will be returned unmodified.
func (s KVSeq[K, V]) Fold(initialKey K, initialVal V, combine reducer2[K, V]) (K, V) {
	s(func(k K, v V) bool {
		initialKey, initialVal = combine(initialKey, initialVal, k, v)
		return true
	})
	return initialKey, initialVal
}

// TryFold is identical to [KVSeq.Fold], except it will recover any panic
// caused by [PanicHaltIteration] and return the wrapped error.
func (s KVSeq[K, V]) TryFold(initialKey K, initialVal V, combine reducer2[K, V]) (_ K, _ V, err error) {
	defer recoverHaltIteration(&err)
	k, v := s.Fold(initialKey, initialVal, combine)
	return k, v, nil
}

// First consumes the iterator and returns its first key/value pair. If the
// iterator is empty then zero values will be returned along with an error.
func (s KVSeq[K, V]) First() (K, V, error) {
	var key K
	var val V
	isEmpty := true
	s(func(k K, v V) bool {
		key, val, isEmpty = k, v, false
		return false
	})
	if isEmpty {
		return key, val, EmptySeqErr
	}
	return key, val, nil
}

// TryFirst is identical to [KVSeq.First], except it will recover any panic
// caused by [PanicHaltIteration] and return the wrapped error.
func (s KVSeq[K, V]) TryFirst() (_ K, _ V, err error) {
	defer recoverHaltIteration(&err)
	return s.First()
}

// Last consumes the iterator and returns its last key/value pair. If the
// iterator is empty then zero values will be returned along with an error.
func (s KVSeq[K, V]) Last() (K, V, error) {
	var key K
	var val V
	isEmpty := true
	s(func(k K, v V) bool {
		if isEmpty {
			isEmpty = false
		}
		key, val = k, v
		return true
	})
	if isEmpty {
		return key, val, EmptySeqErr
	}
	return key, val, nil
}

// TryLast is identical to [KVSeq.Last], except it will recover any panic
// caused by [PanicHaltIteration] and return the wrapped error.
func (s KVSeq[K, V]) TryLast() (_ K, _ V, err error) {
	defer recoverHaltIteration(&err)
	return s.Last()
}

// Any returns true if test returns true for at least one key/value pair in the
// iterator, and false otherwise. Returns false for an empty iterator.
func (s KVSeq[K, V]) Any(test yielder2[K, V]) bool {
	result := false
	s(func(k K, v V) bool {
		if test(k, v) {
			result = true
			return false
		}
		return true
	})
	return result
}

// TryAny is identical to [KVSeq.Any], except it will recover any panic caused
// by [PanicHaltIteration] and return the wrapped error.
func (s KVSeq[K, V]) TryAny(test yielder2[K, V]) (_ bool, err error) {
	defer recoverHaltIteration(&err)
	return s.Any(test), nil
}

// Every returns true if test returns false for every key/value pair of the
// iterator, and false otherwise. Returns true for an empty iterator.
func (s KVSeq[K, V]) None(test yielder2[K, V]) bool {
	result := true
	s(func(k K, v V) bool {
		if test(k, v) {
			result = false
			return false
		}
		return true
	})
	return result
}

// TryNone is identical to [KVSeq.None], except it will recover any panic
// caused by [PanicHaltIteration] and return the wrapped error.
func (s KVSeq[K, V]) TryNone(test yielder2[K, V]) (_ bool, err error) {
	return s.None(test), nil
}

// Every returns true if test returns true for every key/value pair of the
// iterator, and false otherwise. Returns true for an empty iterator.
func (s KVSeq[K, V]) Every(test yielder2[K, V]) bool {
	result := true
	s(func(k K, v V) bool {
		if !test(k, v) {
			result = false
			return false
		}
		return true
	})
	return result
}

// TryEvery is identical to [KVSeq.Every], except it will recover any panic
// caused by [PanicHaltIteration] and return the wrapped error.
func (s KVSeq[K, V]) TryEvery(test yielder2[K, V]) (_ bool, err error) {
	defer recoverHaltIteration(&err)
	return s.Every(test), nil
}

func (s KVSeq[K, V]) Filter(filter yielder2[K, V]) KVSeq[K, V] {
	return func(yield yielder2[K, V]) {
		s(func(k K, v V) bool {
			if !filter(k, v) {
				return true
			}
			return yield(k, v)
		})
	}
}

// Skip skips the first toSkip key/value pairs of the iterator. If toSkip is
// greater than or equal to the number of elements in the iterator the result
// will be an empty iterator.
func (s KVSeq[K, V]) Skip(toSkip int) KVSeq[K, V] {
	return func(yield yielder2[K, V]) {
		var skipped int
		s(func(k K, v V) bool {
			if skipped < toSkip {
				skipped++
				return true
			}
			return yield(k, v)
		})
	}
}

// SkipWhile skips the leading key/value pairs for which test returns true.
func (s KVSeq[K, V]) SkipWhile(test yielder2[K, V]) KVSeq[K, V] {
	return func(yield yielder2[K, V]) {
		skipping := true
		s(func(k K, v V) bool {
			if skipping {
				if test(k, v) {
					return true
				}
				skipping = false
			}
			return yield(k, v)
		})
	}
}

// Take restricts the iterator to at most the first toTake key/value pairs.
func (s KVSeq[K, V]) Take(toTake int) KVSeq[K, V] {
	return func(yield yielder2[K, V]) {
		var took int
		s(func(k K, v V) bool {
			if took >= toTake {
				return false
			}
			took++
			return yield(k, v)
		})
	}
}

// TakeWhile restricts the iterator to the leading key/value pairs for which
// test returns true.
func (s KVSeq[K, V]) TakeWhile(test yielder2[K, V]) KVSeq[K, V] {
	return func(yield yielder2[K, V]) {
		s(func(k K, v V) bool {
			return test(k, v) && yield(k, v)
		})
	}
}
