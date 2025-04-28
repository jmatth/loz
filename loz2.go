package loz

import (
	"errors"
	"iter"
	"slices"
)

type yielder2[K, V any] = func(K, V) bool

type Seq2[K, V any] iter.Seq2[K, V]

func All[V any](slice []V) Seq2[int, V] {
	return Seq2[int, V](slices.All(slice))
}

func (s Seq2[K, V]) IterKeys() Seq[K] {
	return func(yield yielder[K]) {
		for k, _ := range s {
			if !yield(k) {
				break
			}
		}
	}
}

func (s Seq2[K, V]) IterValues() Seq[V] {
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

type reducer22[K, V any] = func(K, V, K, V) (K, V)
func (s Seq2[K, V]) Reduce(combine reducer22[K, V]) (K, V, error) {
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
		return keyResult, valResult, errors.New("Reduce called on empty Seq2")
	}
	return keyResult, valResult, nil
}

func (s Seq2[K, V]) Fold(initialKey K, initialVal V, combine reducer22[K, V]) (K, V) {
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
		return key, val, errors.New("First called on empty Seq2")
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
		return key, val, errors.New("Last called on empty Seq")
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
		skipping := false
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
	return func(yield yielder2[K, V])  {
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
			if !test(k, v) && yield(k, v) {
				break
			}
		}
	}
}
