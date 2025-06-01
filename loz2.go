package loz

import (
	"iter"
	"maps"
)

type Seq2[K, V any] iter.Seq2[K, V]

// IterMap creates a Seq over the key/value pairs of a map.
func IterMap[K comparable, V any](input map[K]V) Seq2[K, V] {
	return Seq2[K, V](maps.All(input))
}

// ToKeys converts a Seq2[K, V] to a Seq[K], continuing the iteration with only the keys.
func (s Seq2[K, V]) Keys() Seq[K] {
	return func(yield yielder[K]) {
		for k := range s {
			if !yield(k) {
				break
			}
		}
	}
}

// Values converts a Seq2[K, V] to a Seq[V], continuing the iteration with only the values.
func (s Seq2[K, V]) Values() Seq[V] {
	return func(yield yielder[V]) {
		for _, v := range s {
			if !yield(v) {
				break
			}
		}
	}
}

func (s Seq2[K, V]) ForEach(process func(K, V)) {
	for k, v := range s {
		process(k, v)
	}
}

func (s Seq2[K, V]) Map(mapper func(K, V) (K, V)) Seq2[K, V] {
	return func(yield yielder2[K, V]) {
		for k, v := range s {
			if !yield(mapper(k, v)) {
				return
			}
		}
	}
}

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

func (s Seq2[K, V]) Fold(initialKey K, initialVal V, combine reducer2[K, V]) (K, V) {
	for k, v := range s {
		initialKey, initialVal = combine(initialKey, initialVal, k, v)
	}
	return initialKey, initialVal
}

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

func (s Seq2[K, V]) Any(test yielder2[K, V]) bool {
	for k, v := range s {
		if test(k, v) {
			return true
		}
	}
	return false
}

func (s Seq2[K, V]) None(test yielder2[K, V]) bool {
	for k, v := range s {
		if test(k, v) {
			return false
		}
	}
	return true
}

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

func (s Seq2[K, V]) TakeWhile(test yielder2[K, V]) Seq2[K, V] {
	return func(yield yielder2[K, V]) {
		for k, v := range s {
			if !test(k, v) || !yield(k, v) {
				break
			}
		}
	}
}
