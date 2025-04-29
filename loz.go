package loz

import (
	"errors"
	"iter"
	"slices"
)

type yielder[V any] = func(V) bool

// Seq is an alias to [iter.Seq] that provides additional methods for filtering and transforming the elements.
type Seq[V any] iter.Seq[V]

// IterSlice creates a Seq over the contents of a slice.
func IterSlice[V any](slice []V) Seq[V] {
	return Seq[V](slices.Values(slice))
}

// ToSlice collects all the elements within the iterator into a slice by calling [slices.Collect].
func (s Seq[V]) ToSlice() []V {
	return slices.Collect(iter.Seq[V](s))
}

// ForEach consumes the iterator and calls the provided function with each of the elements.
func (s Seq[V]) ForEach(process func(V)) {
	for v := range s {
		process(v)
	}
}

// Map transforms the elements within the iterator using the provided mapper function.
// Due to limitations of the Go type system, the mapped value must be the same type as the input.
// To perform mapping operations that change type, see [Map1], [Map2], etc.
func (s Seq[V]) Map(mapper func(V) V) Seq[V] {
	return func(yield yielder[V]) {
		for v := range s {
			if !yield(mapper(v)) {
				break
			}
		}
	}
}

type reducer[V, O any] = func(O, V) O

// Reduce reduces the iterator to a single value by iteratively combining its elements using the provided function.
// If the iterator is empty a zero value will be returned along with an error.
func (s Seq[V]) Reduce(combine reducer[V, V]) (V, error) {
	var result V
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
		return result, errors.New("Reduce called on empty Seq")
	}
	return result, nil
}

// Fold reduces the iterator to a single value by iteratively combining its elements with an initial value using the provided function.
// If the iterator is empty the initial value will be returned unmodified.
func (s Seq[V]) Fold(initial V, combine reducer[V, V]) V {
	for v := range s {
		initial = combine(initial, v)
	}
	return initial
}

// First consumes the iterator and returns its first element.
// If the iterator is empty a zero value will be returned with an error.
func (s Seq[V]) First() (V, error) {
	var result V
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
func (s Seq[V]) Last() (V, error) {
	isEmpty := true
	var result V
	for result = range s {
		if isEmpty {
			isEmpty = false
		}
	}
	if isEmpty {
		return result, errors.New("Last called on empty Seq")
	}
	return result, nil
}

// Any returns true if test returns true for at least one element in the iterator, and false otherwise.
// Returns false for an empty iterator.
func (s Seq[V]) Any(test yielder[V]) bool {
	for v := range s {
		if test(v) {
			return true
		}
	}
	return false
}

// Every returns true if test returns false for every element of the iterator, and false otherwise.
// Returns true for an empty iterator.
func (s Seq[V]) None(test yielder[V]) bool {
	for v := range s {
		if test(v) {
			return false
		}
	}
	return true
}

// Every returns true if test returns true for every element of the iterator, and false otherwise.
// Returns true for an empty iterator.
func (s Seq[V]) Every(test yielder[V]) bool {
	for v := range s {
		if !test(v) {
			return false
		}
	}
	return true
}

// Filter filters the iterator to only include only elements for which filter returns true.
func (s Seq[V]) Filter(filter yielder[V]) Seq[V] {
	return func(yield yielder[V]) {
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
func (s Seq[V]) Skip(toSkip int) Seq[V] {
	return func(yield yielder[V]) {
		var skipped int
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
func (s Seq[V]) SkipWhile(test yielder[V]) Seq[V] {
	return func(yield yielder[V]) {
		skipping := true
		for v := range s {
			if skipping {
				if test(v) {
					continue
				}
				skipping = false
			}
			if !yield(v) {
				break
			}
		}
	}
}

// Take restricts the iterator to at most the first toTake elements.
func (s Seq[V]) Take(toTake int) Seq[V] {
	return func(yield yielder[V]) {
		var took int
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
func (s Seq[V]) TakeWhile(test yielder[V]) Seq[V] {
	return func(yield yielder[V]) {
		for v := range s {
			if !test(v) || !yield(v) {
				break
			}
		}
	}
}

func (s Seq[V]) Indexed() Seq2[int, V] {
	return func(yield yielder2[int, V]) {
		var i int
		for v := range s {
			if !yield(i, v) {
				break
			}
			i++
		}
	}
}

func (s Seq[V]) Expand(toElements func(V) Seq[V]) Seq[V] {
	return func(yield yielder[V]) {
		for v := range s {
			for e := range toElements(v) {
				if !yield(e) {
					break
				}
			}
		}
	}
}
