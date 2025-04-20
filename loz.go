package loz

import (
	"errors"
	"iter"
	"slices"
)

type yielder[T any] = func(T) bool

// Seq is an alias to [iter.Seq] that provides additional methods for filtering and transforming the elements.
type Seq[T any] iter.Seq[T]

func Values[T any](slice []T) Seq[T] {
	return Seq[T](slices.Values(slice))
}

// ToSlice collects all the elements within the iterator into a slice by calling [slices.Collect].
func (s Seq[T]) ToSlice() []T {
	return slices.Collect(iter.Seq[T](s))
}

// ForEach consumes the iterator and calls the provided function with each of the elements.
func (s Seq[T]) ForEach(process func(T)) {
	for v := range s {
		process(v)
	}
}

// Map transforms the elements within the iterator using the provided mapper function.
// Due to limitations of the Go type system, the mapped value must be the same type as the input.
// To perform mapping operations that change type, see [Map1], [Map2], etc.
func (s Seq[T]) Map(mapper func(T) T) Seq[T] {
	return func(yield yielder[T]) {
		for v := range s {
			if !yield(mapper(v)) {
				break
			}
		}
	}
}

type reducer[E, O any] = func(a O, b E) O

// Reduce reduces the iterator to a single value by iteratively combining its elements using the provided function.
// If the iterator is empty a zero value will be returned along with an error.
func (s Seq[T]) Reduce(combine reducer[T, T]) (T, error) {
	var result T
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

// Fold reduces the iterator to a single value by iteratively combining its elements with an initial value using the provided function.
// If the iterator is empty the initial value will be returned unmodified.
func (s Seq[T]) Fold(initial T, combine reducer[T, T]) T {
	for v := range s {
		initial = combine(initial, v)
	}
	return initial
}

// First consumes the iterator and returns its first element.
// If the iterator is empty a zero value will be returned with an error.
func (s Seq[T]) First() (T, error) {
	var result T
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

// Last consumes the iterator and returns its last element.
// If the iterator is empty a zero value will be returned with an error.
func (s Seq[T]) Last() (T, error) {
	isEmpty := true
	var result T
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

// Any returns true if test returns true for at least one element in the iterator, and false otherwise.
// Returns false for an empty iterator.
func (s Seq[T]) Any(test yielder[T]) bool {
	for v := range s {
		if test(v) {
			return true
		}
	}
	return false
}

// Every returns true if test returns false for every element of the iterator, and false otherwise.
// Returns true for an empty iterator.
func (s Seq[T]) None(test yielder[T]) bool {
	for v := range s {
		if test(v) {
			return false
		}
	}
	return true
}

// Every returns true if test returns true for every element of the iterator, and false otherwise.
// Returns true for an empty iterator.
func (s Seq[T]) Every(test yielder[T]) bool {
	for v := range s {
		if !test(v) {
			return false
		}
	}
	return true
}

// Where filters the iterator to only include only elements for which filter returns true.
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

// Skip skips the first toSkip elements of the iterator.
// If toSkip is greater than or equal to the number of elements in the iterator the result will be an empty iterator.
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

// SkipWhile skips the leading elements for which test returns true.
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

// Take restricts the iterator to at most the first toTake elements.
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

// TakeWhile restricts the iterator to the leading elements for which test returns true.
func (s Seq[E]) TakeWhile(test yielder[E]) Seq[E] {
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
