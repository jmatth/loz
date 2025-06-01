package loz

import (
	"iter"
	"maps"
)

// Seq2 is an alias to [iter.Seq2] that provides additional methods for
// filtering, transforming, and collecting the elements.
type Seq2[K, V any] iter.Seq2[K, V]

// IterMap creates a Seq over the key/value pairs of a map.
func IterMap[K comparable, V any](input map[K]V) Seq2[K, V] {
	return Seq2[K, V](maps.All(input))
}

// ToKeys converts a Seq2[K, V] to a Seq[K], continuing the iteration with only
// the keys.
func (s Seq2[K, V]) Keys() Seq[K] {
	return func(yield yielder[K]) {
		for k := range s {
			if !yield(k) {
				break
			}
		}
	}
}

// Values converts a Seq2[K, V] to a Seq[V], continuing the iteration with only
// the values.
func (s Seq2[K, V]) Values() Seq[V] {
	return func(yield yielder[V]) {
		for _, v := range s {
			if !yield(v) {
				break
			}
		}
	}
}

// ForEach consumes the iterator and calls the provided function with each of
// the key/value pairs.
func (s Seq2[K, V]) ForEach(process func(K, V)) {
	for k, v := range s {
		process(k, v)
	}
}

// Map transforms the key/value pairs within the iterator using the provided
// mapper function. Due to limitations of the Go type system, the mapped keys
// and values must be the same types as the input. To perform mapping
// operations that change types, see [KVMapper1], [KVMapper2], etc.
func (s Seq2[K, V]) Map(mapper func(K, V) (K, V)) Seq2[K, V] {
	return func(yield yielder2[K, V]) {
		for k, v := range s {
			if !yield(mapper(k, v)) {
				return
			}
		}
	}
}

// Reduce reduces the iterator to a single key/value pair by iteratively
// combining its elements using the provided function. If the iterator is empty
// then zero values will be returned along with an error.
func (s Seq2[K, V]) Reduce(combine reducer2[K, V]) (K, V, error) {
	var keyResult K
	var valResult V
	isFirst := true
	for k, v := range s {
		if isFirst {
			keyResult = k
			valResult = v
			isFirst = false
			continue
		}
		keyResult, valResult = combine(keyResult, valResult, k, v)
	}
	if isFirst {
		return keyResult, valResult, EmptySeqErr
	}
	return keyResult, valResult, nil
}

// Fold reduces the iterator to a single key/value pair by iteratively
// combining its elements with initial values using the provided function. If
// the iterator is empty the initial values will be returned unmodified.
func (s Seq2[K, V]) Fold(initialKey K, initialVal V, combine reducer2[K, V]) (K, V) {
	for k, v := range s {
		initialKey, initialVal = combine(initialKey, initialVal, k, v)
	}
	return initialKey, initialVal
}

// First consumes the iterator and returns its first key/value pair. If the
// iterator is empty then zero values will be returned along with an error.
func (s Seq2[K, V]) First() (K, V, error) {
	var key K
	var val V
	isEmpty := true
	for key, val = range s {
		isEmpty = false
		break
	}
	if isEmpty {
		return key, val, EmptySeqErr
	}
	return key, val, nil
}

// Last consumes the iterator and returns its last key/value pair. If the
// iterator is empty then zero values will be returned along with an error.
func (s Seq2[K, V]) Last() (K, V, error) {
	var key K
	var val V
	isEmpty := true
	for key, val = range s {
		if isEmpty {
			isEmpty = false
		}
	}
	if isEmpty {
		return key, val, EmptySeqErr
	}
	return key, val, nil
}

// Any returns true if test returns true for at least one key/value pair in the
// iterator, and false otherwise. Returns false for an empty iterator.
func (s Seq2[K, V]) Any(test yielder2[K, V]) bool {
	for k, v := range s {
		if test(k, v) {
			return true
		}
	}
	return false
}

// Every returns true if test returns false for every key/value pair of the
// iterator, and false otherwise. Returns true for an empty iterator.
func (s Seq2[K, V]) None(test yielder2[K, V]) bool {
	for k, v := range s {
		if test(k, v) {
			return false
		}
	}
	return true
}

// Every returns true if test returns true for every key/value pair of the
// iterator, and false otherwise. Returns true for an empty iterator.
func (s Seq2[K, V]) Every(test yielder2[K, V]) bool {
	for k, v := range s {
		if !test(k, v) {
			return false
		}
	}
	return true
}

func (s Seq2[K, V]) Filter(filter yielder2[K, V]) Seq2[K, V] {
	return func(yield yielder2[K, V]) {
		for k, v := range s {
			if filter(k, v) {
				if !yield(k, v) {
					break
				}
			}
		}
	}
}

// Skip skips the first toSkip key/value pairs of the iterator. If toSkip is
// greater than or equal to the number of elements in the iterator the result
// will be an empty iterator.
func (s Seq2[K, V]) Skip(toSkip int) Seq2[K, V] {
	return func(yield yielder2[K, V]) {
		var skipped int
		for k, v := range s {
			if skipped < toSkip {
				skipped++
				continue
			}
			if !yield(k, v) {
				break
			}
		}
	}
}

// SkipWhile skips the leading key/value pairs for which test returns true.
func (s Seq2[K, V]) SkipWhile(test yielder2[K, V]) Seq2[K, V] {
	return func(yield yielder2[K, V]) {
		skipping := true
		for k, v := range s {
			if skipping {
				if test(k, v) {
					continue
				}
				skipping = false
			}
			if !yield(k, v) {
				break
			}
		}
	}
}

// Take restricts the iterator to at most the first toTake key/value pairs.
func (s Seq2[K, V]) Take(toTake int) Seq2[K, V] {
	return func(yield yielder2[K, V]) {
		var took int
		for k, v := range s {
			if took >= toTake {
				break
			}
			took++
			if !yield(k, v) {
				break
			}
		}
	}
}

// TakeWhile restricts the iterator to the leading key/value pairs for which
// test returns true.
func (s Seq2[K, V]) TakeWhile(test yielder2[K, V]) Seq2[K, V] {
	return func(yield yielder2[K, V]) {
		for k, v := range s {
			if !test(k, v) || !yield(k, v) {
				break
			}
		}
	}
}
