package loz

import (
	"iter"
	"slices"

	. "github.com/jmatth/loz/internal"
)

// Seq is an alias to [iter.Seq] that provides additional methods for filtering,
// transforming, and collecting the elements.
type Seq[V any] iter.Seq[V]

// Generate creates a [Seq] by calling the provided generator with [0, count).
// A count < 1 yields an empty iterator.
func Generate[V any](count int, generator func(idx int) V) Seq[V] {
	return func(yield Yielder[V]) {
		for i := range count {
			if !yield(generator(i)) {
				break
			}
		}
	}
}

// IterSlice creates a Seq over the contents of a slice.
func IterSlice[V any](slice []V) Seq[V] {
	return Seq[V](slices.Values(slice))
}

// Unwrap casts a [Seq] back to an [iter.Seq].
func (s Seq[V]) Unwrap() iter.Seq[V] {
	return iter.Seq[V](s)
}

// ForEach consumes the iterator and calls the provided function with each of
// the elements.
func (s Seq[V]) ForEach(process Processor[V]) {
	s(func(v V) bool {
		process(v)
		return true
	})
}

// TryForEach is identical to [Seq.ForEach], except it will recover any panic
// caused by [PanicHaltIteration] and return the wrapped error.
func (s Seq[V]) TryForEach(process Processor[V]) (err error) {
	defer RecoverHaltIteration(&err)
	s.ForEach(process)
	return nil
}

// Map transforms the elements within the iterator using the provided mapper
// function. Due to limitations of the Go type system, the mapped value must be
// the same type as the input.
func (s Seq[V]) Map[R any](mapper Mapper[V, R]) Seq[R] {
	return func(yield Yielder[R]) {
		s(func(v V) bool {
			return yield(mapper(v))
		})
	}
}

// FilterMap is a combination of [Seq.Filter] and [Seq.Map]. If the provided
// mapper function returns false, then the current element of the iteration
// will be skipped. If true is returned, then the mapped value is passed to the
// next iteration stage.
func (s Seq[V]) FilterMap[R any](mapper FilteringMapper[V, R]) Seq[R] {
	return func(yield Yielder[R]) {
		s(func(v V) bool {
			mapped, ok := mapper(v)
			if !ok {
				return true
			}
			return yield(mapped)
		})
	}
}

// Reduce reduces the iterator to a single value by iteratively combining its
// elements using the provided function. If the iterator is empty a zero value
// will be returned along with an error. To reduce to a different type and/or
// provide a starting value, see [Seq.Fold].
func (s Seq[V]) Reduce(combine Reducer[V, V]) (V, error) {
	isFirst := true
	var result V
	s(func(v V) bool {
		if isFirst {
			result = v
			isFirst = false
			return true
		}
		result = combine(result, v)
		return true
	})
	if isFirst {
		return result, EmptySeqErr
	}
	return result, nil
}

// TryReduce is identical to [Seq.Reduce], except it will recover any panic
// caused by [PanicHaltIteration] and return the wrapped error.
func (s Seq[V]) TryReduce(combine Reducer[V, V]) (result V, err error) {
	defer RecoverHaltIteration(&err)
	return s.Reduce(combine)
}

// Fold reduces the iterator to a single value by iteratively combining its
// elements with an initial value using the provided function. If the iterator
// is empty the initial value will be returned unmodified.
func (s Seq[V]) Fold[R any](initial R, combine Reducer[V, R]) R {
	s(func(v V) bool {
		initial = combine(initial, v)
		return true
	})
	return initial
}

// TryFold is identical to [Seq.Fold], except it will recover any panic caused
// by [PanicHaltIteration] and return the wrapped error.
func (s Seq[V]) TryFold[R any](initial R, combine Reducer[V, R]) (result R, err error) {
	defer RecoverHaltIteration(&err)
	return s.Fold(initial, combine), nil
}

func (s Seq[V]) Collect[R any](collector Collector[R, V]) R {
	initial := collector.Initial()
	return s.Fold(initial, collector.Collect)
}

func (s Seq[V]) TryCollect[R any](collector Collector[R, V]) (R, error) {
	initial := collector.Initial()
	return s.TryFold(initial, collector.Collect)
}

// First consumes the iterator and returns its first element. If the iterator
// is empty a zero value will be returned with an error.
func (s Seq[V]) First() (V, error) {
	isEmpty := true
	var result V
	s(func(v V) bool {
		result, isEmpty = v, false
		return false
	})
	if isEmpty {
		return result, EmptySeqErr
	}
	return result, nil
}

// TryFirst is identical to [Seq.First], except it will recover any panic
// caused by [PanicHaltIteration] and return the wrapped error.
func (s Seq[V]) TryFirst() (result V, err error) {
	defer RecoverHaltIteration(&err)
	return s.First()
}

// Last consumes the iterator and returns its last element.
// If the iterator is empty a zero value will be returned with an error.
func (s Seq[V]) Last() (V, error) {
	isEmpty := true
	var result V
	s(func(v V) bool {
		if isEmpty {
			isEmpty = false
		}
		result = v
		return true
	})
	if isEmpty {
		return result, EmptySeqErr
	}
	return result, nil
}

// TryLast is identical to [Seq.Last], except it will recover any panic caused
// by [PanicHaltIteration] and return the wrapped error.
func (s Seq[V]) TryLast() (result V, err error) {
	defer RecoverHaltIteration(&err)
	return s.Last()
}

// Any returns true if test returns true for at least one element in the
// iterator, and false otherwise. Returns false for an empty iterator.
func (s Seq[V]) Any(test Yielder[V]) bool {
	result := false
	s(func(v V) bool {
		if test(v) {
			result = true
			return false
		}
		return true
	})
	return result
}

// TryAny is identical to [Seq.Any], except it will recover any panic caused by
// [PanicHaltIteration] and return the wrapped error.
func (s Seq[V]) TryAny(test Yielder[V]) (result bool, err error) {
	defer RecoverHaltIteration(&err)
	return s.Any(test), nil
}

// None returns true if test returns false for every element of the iterator,
// and false otherwise. Returns true for an empty iterator.
func (s Seq[V]) None(test Yielder[V]) bool {
	result := true
	s(func(v V) bool {
		if test(v) {
			result = false
			return false
		}
		return true
	})
	return result
}

// TryNone is identical to [Seq.None], except it will recover any panic caused
// by [PanicHaltIteration] and return the wrapped error.
func (s Seq[V]) TryNone(test Yielder[V]) (result bool, err error) {
	defer RecoverHaltIteration(&err)
	return s.None(test), err
}

// Every returns true if test returns true for every element of the iterator,
// and false otherwise. Returns true for an empty iterator.
func (s Seq[V]) Every(test Yielder[V]) bool {
	result := true
	s(func(v V) bool {
		if !test(v) {
			result = false
			return false
		}
		return true
	})
	return result
}

// TryEvery is identical to [Seq.Every], except it will recover any panic
// caused by [PanicHaltIteration] and return the wrapped error.
func (s Seq[V]) TryEvery(test Yielder[V]) (result bool, err error) {
	defer RecoverHaltIteration(&err)
	return s.Every(test), nil
}

// Filter filters the iterator to only include only elements for which filter
// returns true.
func (s Seq[V]) Filter(filter Yielder[V]) Seq[V] {
	return func(yield Yielder[V]) {
		s(func(v V) bool {
			if !filter(v) {
				return true
			}
			return yield(v)
		})
	}
}

// Skip skips the first toSkip elements of the iterator. If toSkip is greater
// than or equal to the number of elements in the iterator the result will be
// an empty iterator.
func (s Seq[V]) Skip(toSkip int) Seq[V] {
	return func(yield Yielder[V]) {
		var skipped int
		s(func(v V) bool {
			if skipped < toSkip {
				skipped++
				return true
			}
			return yield(v)
		})
	}
}

// SkipWhile skips the leading elements for which test returns true.
func (s Seq[V]) SkipWhile(test Yielder[V]) Seq[V] {
	return func(yield Yielder[V]) {
		skipping := true
		s(func(v V) bool {
			if skipping {
				if test(v) {
					return true
				}
				skipping = false
			}
			return yield(v)
		})
	}
}

// Take restricts the iterator to at most the first toTake elements.
func (s Seq[V]) Take(toTake int) Seq[V] {
	return func(yield Yielder[V]) {
		var took int
		s(func(v V) bool {
			if took >= toTake {
				return false
			}
			took++
			return yield(v)
		})
	}
}

// TakeWhile restricts the iterator to the leading elements for which test
// returns true.
func (s Seq[V]) TakeWhile(test Yielder[V]) Seq[V] {
	return func(yield Yielder[V]) {
		s(func(v V) bool {
			return test(v) && yield(v)
		})
	}
}

func (s Seq[V]) Indexed() KVSeq[int, V] {
	return func(yield Yielder2[int, V]) {
		var i int
		s(func(v V) bool {
			result := yield(i, v)
			i++
			return result
		})
	}
}

func (s Seq[V]) Expand[R any](toElements Mapper[V, Seq[R]]) Seq[R] {
	return func(yield Yielder[R]) {
		s(func(v V) bool {
			for e := range toElements(v) {
				if !yield(e) {
					return false
				}
			}
			return true
		})
	}
}
