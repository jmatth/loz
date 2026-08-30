package loz

import (
	"iter"
	"maps"

	. "github.com/jmatth/loz/internal"
)

type void = struct{}

// KVSeq is an alias to [iter.Seq2] that provides additional methods for
// filtering, transforming, and collecting the elements. Though the name and
// several doc comments imply it contains key/value pairs, that is only the most
// common use case and the relationship between the two values is arbitrary.
type KVSeq[K, V any] iter.Seq2[K, V]

// IterMap creates a Seq over the key/value pairs of a map.
func IterMap[K comparable, V any](input map[K]V) KVSeq[K, V] {
	return KVSeq[K, V](maps.All(input))
}

// Collect consumes the iterator and returns a collection containing its
// contents. How that collection is constructed depends on the collector used.
// See [ToMap] for an example.
func (s KVSeq[K, V]) Collect[R any](collector KVCollector[R, K, V]) R {
	initial := collector.Initial()
	result, _ := s.Fold(initial, void{}, func(acc R, _ struct{}, iterK K, iterV V) (R, void) {
		result := collector.Collect(acc, iterK, iterV)
		return result, void{}
	})
	return result
}

// TryCollect is identical to [KVSeq.Collect], except it will recover any panic
// caused by [PanicHaltIteration] and return the wrapped error.
func (s KVSeq[K, V]) TryCollect[R any](collector KVCollector[R, K, V]) (R, error) {
	initial := collector.Initial()
	result, _, err := s.TryFold(initial, void{}, func(acc R, _ struct{}, iterK K, iterV V) (R, void) {
		result := collector.Collect(acc, iterK, iterV)
		return result, void{}
	})
	return result, err
}

// // CollectMap collects all the elements within the iterator into a `map[K]V`.
// func (s KVSeq[K, V]) CollectMap[RK comparable](keyMapper Mapper[K, RK]) map[RK]V {
// 	return maps.Collect(
// 		s.Map(func(k K, v V) (RK, V) {
// 			return keyMapper(k), v
// 		}).
// 			Unwrap(),
// 	)
// }

// Unwrap casts a [KVSeq] back to an [iter.Seq2].
func (s KVSeq[K, V]) Unwrap() iter.Seq2[K, V] {
	return iter.Seq2[K, V](s)
}

// Keys converts a KVSeq[K, V] to a Seq[K], continuing the iteration with only
// the keys.
func (s KVSeq[K, V]) Keys() Seq[K] {
	return func(yield Yielder[K]) {
		s(func(k K, _ V) bool {
			return yield(k)
		})
	}
}

// Values converts a KVSeq[K, V] to a Seq[V], continuing the iteration with only
// the values.
func (s KVSeq[K, V]) Values() Seq[V] {
	return func(yield Yielder[V]) {
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
	defer RecoverHaltIteration(&err)
	s.ForEach(process)
	return nil
}

// Map transforms the key/value pairs within the iterator using the provided
// mapper function. Due to limitations of the Go type system, the mapped keys
// and values must be the same types as the input.
func (s KVSeq[K, V]) Map[RK, RV any](mapper func(K, V) (RK, RV)) KVSeq[RK, RV] {
	return func(yield KVYielder[RK, RV]) {
		s(func(k K, v V) bool {
			return yield(mapper(k, v))
		})
	}
}

// FilterMap is a combination of [KVSeq.Filter] and [KVSeq.Map]. If the
// provided mapper function returns false, then the current key/value pair of
// the iteration will be skipped. If true is returned, then the mapped
// key/value pair is passed to the next iteration stage.
func (s KVSeq[K, V]) FilterMap[RK, RV any](mapper func(K, V) (RK, RV, bool)) KVSeq[RK, RV] {
	return func(yield KVYielder[RK, RV]) {
		s(func(k K, v V) bool {
			mk, mv, ok := mapper(k, v)
			if !ok {
				return true
			}
			return yield(mk, mv)
		})
	}
}

// Reduce reduces the iterator to a single key/value pair by iteratively
// combining its elements using the provided function. If the iterator is empty
// then zero values will be returned along with an error.
func (s KVSeq[K, V]) Reduce(combine KVReducer[K, V, K, V]) (K, V, error) {
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
func (s KVSeq[K, V]) TryReduce(combine KVReducer[K, V, K, V]) (_ K, _ V, err error) {
	defer RecoverHaltIteration(&err)
	return s.Reduce(combine)
}

// Fold reduces the iterator to a single key/value pair by iteratively
// combining its elements with initial values using the provided function. If
// the iterator is empty the initial values will be returned unmodified.
func (s KVSeq[K, V]) Fold[RK, RV any](initialKey RK, initialVal RV, combine KVReducer[K, V, RK, RV]) (RK, RV) {
	s(func(k K, v V) bool {
		initialKey, initialVal = combine(initialKey, initialVal, k, v)
		return true
	})
	return initialKey, initialVal
}

// TryFold is identical to [KVSeq.Fold], except it will recover any panic
// caused by [PanicHaltIteration] and return the wrapped error.
func (s KVSeq[K, V]) TryFold[RK, RV any](initialKey RK, initialVal RV, combine KVReducer[K, V, RK, RV]) (_ RK, _ RV, err error) {
	defer RecoverHaltIteration(&err)
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
	defer RecoverHaltIteration(&err)
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
	defer RecoverHaltIteration(&err)
	return s.Last()
}

// Any returns true if test returns true for at least one key/value pair in the
// iterator, and false otherwise. Returns false for an empty iterator.
func (s KVSeq[K, V]) Any(test KVYielder[K, V]) bool {
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
func (s KVSeq[K, V]) TryAny(test KVYielder[K, V]) (_ bool, err error) {
	defer RecoverHaltIteration(&err)
	return s.Any(test), nil
}

// None returns true if test returns false for every key/value pair of the
// iterator, and false otherwise. Returns true for an empty iterator.
func (s KVSeq[K, V]) None(test KVYielder[K, V]) bool {
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
func (s KVSeq[K, V]) TryNone(test KVYielder[K, V]) (_ bool, err error) {
	defer RecoverHaltIteration(&err)
	return s.None(test), nil
}

// Every returns true if test returns true for every key/value pair of the
// iterator, and false otherwise. Returns true for an empty iterator.
func (s KVSeq[K, V]) Every(test KVYielder[K, V]) bool {
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
func (s KVSeq[K, V]) TryEvery(test KVYielder[K, V]) (_ bool, err error) {
	defer RecoverHaltIteration(&err)
	return s.Every(test), nil
}

// Filter filters the iterator to only include only key/value pairs for which
// `filter` returns true.
func (s KVSeq[K, V]) Filter(filter KVYielder[K, V]) KVSeq[K, V] {
	return func(yield KVYielder[K, V]) {
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
	return func(yield KVYielder[K, V]) {
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
func (s KVSeq[K, V]) SkipWhile(test KVYielder[K, V]) KVSeq[K, V] {
	return func(yield KVYielder[K, V]) {
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
	return func(yield KVYielder[K, V]) {
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
func (s KVSeq[K, V]) TakeWhile(test KVYielder[K, V]) KVSeq[K, V] {
	return func(yield KVYielder[K, V]) {
		s(func(k K, v V) bool {
			return test(k, v) && yield(k, v)
		})
	}
}
