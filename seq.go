package loz

import (
	"iter"
	"slices"
)

// Seq is an alias to [iter.Seq] that provides additional methods for filtering,
// transforming, and collecting the elements.
type Seq[V any] iter.Seq[V]

// IterSlice creates a Seq over the contents of a slice.
func IterSlice[V any](slice []V) Seq[V] {
	return Seq[V](slices.Values(slice))
}

// CollectSlice collects all the elements within the iterator into a slice by
// calling [slices.Collect].
func (s Seq[V]) CollectSlice() []V {
	return slices.Collect(iter.Seq[V](s))
}

// TryCollectSlice is identical to [Seq.CollectSlice], except it will recover
// any panic caused by [PanicHaltIteration] and return the wrapped error.
func (s Seq[V]) TryCollectSlice() (result []V, err error) {
	defer recoverHaltIteration(&err)
	return s.CollectSlice(), nil
}

// ForEach consumes the iterator and calls the provided function with each of
// the elements.
func (s Seq[V]) ForEach(process processor[V]) {
	for v := range s {
		process(v)
	}
}

// TryForEach is identical to [Seq.ForEach], except it will recover any panic
// caused by [PanicHaltIteration] and return the wrapped error.
func (s Seq[V]) TryForEach(process processor[V]) (err error) {
	defer recoverHaltIteration(&err)
	s.ForEach(process)
	return nil
}

// Map transforms the elements within the iterator using the provided mapper
// function. Due to limitations of the Go type system, the mapped value must be
// the same type as the input. To perform mapping operations that change type,
// see [Map1], [Map2], etc.
func (s Seq[V]) Map(mapper mapper[V, V]) Seq[V] {
	return func(yield yielder[V]) {
		for v := range s {
			if !yield(mapper(v)) {
				break
			}
		}
	}
}

// FilterMap is a combination of [Seq.Filter] and [Seq.Map]. If the provided
// mapper function returns an error, then the current element of the iteration
// will be skipped. If no error is returned, then the mapped value is passed to
// the next iteration stage as normal.
func (s Seq[V]) FilterMap(mapper filteringMapper[V, V]) Seq[V] {
	return func(yield yielder[V]) {
		for v := range s {
			mapped, err := mapper(v)
			if err != nil {
				continue
			}
			if !yield(mapped) {
				break
			}
		}
	}
}

// Reduce reduces the iterator to a single value by iteratively combining its
// elements using the provided function. If the iterator is empty a zero value
// will be returned along with an error.
func (s Seq[V]) Reduce(combine reducer[V, V]) (V, error) {
	isFirst := true
	var result V
	for v := range s {
		if isFirst {
			result = v
			isFirst = false
			continue
		}
		result = combine(result, v)
	}
	if isFirst {
		return result, EmptySeqErr
	}
	return result, nil
}

// TryReduce is identical to [Seq.Reduce], except it will recover any panic
// caused by [PanicHaltIteration] and return the wrapped error.
func (s Seq[V]) TryReduce(combine reducer[V, V]) (result V, err error) {
	defer recoverHaltIteration(&err)
	return s.Reduce(combine)
}

// Fold reduces the iterator to a single value by iteratively combining its
// elements with an initial value using the provided function. If the iterator
// is empty the initial value will be returned unmodified.
func (s Seq[V]) Fold(initial V, combine reducer[V, V]) V {
	for v := range s {
		initial = combine(initial, v)
	}
	return initial
}

// TryFold is identical to [Seq.Fold], except it will recover any panic caused
// by [PanicHaltIteration] and return the wrapped error.
func (s Seq[V]) TryFold(initial V, combine reducer[V, V]) (result V, err error) {
	defer recoverHaltIteration(&err)
	return s.Fold(initial, combine), nil
}

// First consumes the iterator and returns its first element. If the iterator
// is empty a zero value will be returned with an error.
func (s Seq[V]) First() (V, error) {
	isEmpty := true
	var result V
	for result = range s {
		isEmpty = false
		break
	}
	if isEmpty {
		return result, EmptySeqErr
	}
	return result, nil
}

// TryFirst is identical to [Seq.First], except it will recover any panic
// caused by [PanicHaltIteration] and return the wrapped error.
func (s Seq[V]) TryFirst() (result V, err error) {
	defer recoverHaltIteration(&err)
	return s.First()
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
		return result, EmptySeqErr
	}
	return result, nil
}

// TryLast is identical to [Seq.Last], except it will recover any panic caused
// by [PanicHaltIteration] and return the wrapped error.
func (s Seq[V]) TryLast() (result V, err error) {
	defer recoverHaltIteration(&err)
	return s.Last()
}

// Any returns true if test returns true for at least one element in the
// iterator, and false otherwise. Returns false for an empty iterator.
func (s Seq[V]) Any(test yielder[V]) bool {
	for v := range s {
		if test(v) {
			return true
		}
	}
	return false
}

// TryAny is identical to [Seq.Any], except it will recover any panic caused by
// [PanicHaltIteration] and return the wrapped error.
func (s Seq[V]) TryAny(test yielder[V]) (result bool, err error) {
	defer recoverHaltIteration(&err)
	return s.Any(test), nil
}

// Every returns true if test returns false for every element of the iterator,
// and false otherwise. Returns true for an empty iterator.
func (s Seq[V]) None(test yielder[V]) bool {
	for v := range s {
		if test(v) {
			return false
		}
	}
	return true
}

// TryNone is identical to [Seq.None], except it will recover any panic caused
// by [PanicHaltIteration] and return the wrapped error.
func (s Seq[V]) TryNone(test yielder[V]) (result bool, err error) {
	defer recoverHaltIteration(&err)
	return s.None(test), err
}

// Every returns true if test returns true for every element of the iterator,
// and false otherwise. Returns true for an empty iterator.
func (s Seq[V]) Every(test yielder[V]) bool {
	for v := range s {
		if !test(v) {
			return false
		}
	}
	return true
}

// TryEvery is identical to [Seq.Every], except it will recover any panic
// caused by [PanicHaltIteration] and return the wrapped error.
func (s Seq[V]) TryEvery(test yielder[V]) (result bool, err error) {
	defer recoverHaltIteration(&err)
	return s.Every(test), nil
}

// Filter filters the iterator to only include only elements for which filter
// returns true.
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

// Skip skips the first toSkip elements of the iterator. If toSkip is greater
// than or equal to the number of elements in the iterator the result will be
// an empty iterator.
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

// TakeWhile restricts the iterator to the leading elements for which test
// returns true.
func (s Seq[V]) TakeWhile(test yielder[V]) Seq[V] {
	return func(yield yielder[V]) {
		for v := range s {
			if !test(v) || !yield(v) {
				break
			}
		}
	}
}

func (s Seq[V]) Indexed() KVSeq[int, V] {
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

func (s Seq[V]) Expand(toElements mapper[V, Seq[V]]) Seq[V] {
	return func(yield yielder[V]) {
	outer:
		for v := range s {
			for e := range toElements(v) {
				if !yield(e) {
					break outer
				}
			}
		}
	}
}
