package loz

import (
	"errors"
	"iter"
	"slices"
)

type yielder[E any] = func(E) bool

type Seq[T any] iter.Seq[T]

func Values[E any](slice []E) Seq[E] {
	return Seq[E](slices.Values(slice))
}

func (s Seq[E]) Where(filter yielder[E]) Seq[E] {
	return func(yield yielder[E]) {
		for v := range s {
			if filter(v) {
				if !yield(v) {
					break
				}
			}
		}
	}
}

func (s Seq[E]) ToSlice() []E {
	return slices.Collect(iter.Seq[E](s))
}

func (s Seq[E]) Skip(toSkip int) Seq[E] {
	return func(yield yielder[E]) {
		skipped := 0
		for v := range s {
			if skipped < toSkip {
				skipped++
				continue
			}
			if !yield(v) {
				break
			}
		}
	}
}

func (s Seq[E]) SkipWhile(test yielder[E]) Seq[E] {
	return func(yield yielder[E]) {
		skipping := true
		for v := range s {
			if skipping {
				if test(v) {
					continue
				} else {
					skipping = false
				}
			}
			if !yield(v) {
				break
			}
		}
	}
}

func (s Seq[E]) Take(toTake int) Seq[E] {
	return func(yield yielder[E]) {
		took := 0
		for v := range s {
			if took >= toTake {
				break
			}
			took++
			if !yield(v) {
				break
			}
		}
	}
}

func (s Seq[E]) TakeWhile(test func(E) bool) Seq[E] {
	return func(yield yielder[E]) {
		for v := range s {
			if !test(v) {
				break
			}
			if !yield(v) {
				break
			}
		}
	}
}

func (s Seq[E]) Reduce(combine func(a, b E) E) (E, error) {
	var result E
	isFirst := true
	for v := range s {
		if isFirst {
			result = v
			isFirst = false
			continue
		}
		result = combine(result, v)
	}
	if isFirst {
		return result, errors.New("Reduce called on a Seq with no elements")
	}
	return result, nil
}

func (s Seq[E]) Fold(initial E, combine func(a, b E) E) E {
	for v := range s {
		initial = combine(initial, v)
	}
	return initial
}

func (s Seq[E]) First() (E, error) {
	var result E
	isEmpty := true
	for result = range s { 
		isEmpty = false
		break
	}
	if isEmpty {
		return result, errors.New("First called on empty Seq")
	}
	return result, nil
}

func (s Seq[E]) Last() (E, error) {
	isEmpty := true
	var result E
	for result = range s {
		if isEmpty {
			isEmpty = false
		}
	}
	if isEmpty {
		return result, errors.New("First called on empty Seq")
	}
	return result, nil
}

func (s Seq[E]) Any(test yielder[E]) bool {
	for v := range s {
		if test(v) {
			return true
		}
	}
	return false
}

func (s Seq[E]) None(test yielder[E]) bool {
	return !s.Any(test)
}

func (s Seq[E]) Every(test yielder[E]) bool {
	for v := range s {
		if !test(v) {
			return false
		}
	}
	return true
}

func (s Seq[E]) ForEach(process func(E)) {
	for v := range s {
		process(v)
	}
}
