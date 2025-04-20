package loz

// A Map1 is a wrapper around [Seq] that provides methods to map to 1 additional type.
type Map1[T1, T2 any] Seq[T1]

// Map transforms the elements within the iterator using the provided mapper function.
func (s Map1[T1, T2]) Map(mapper func(T1) T2) Seq[T2] {
	return func(yield yielder[T2]) {
		for v := range s {
				if !yield(mapper(v)) {
					break
				}
		}
	}
}

// See [Seq.Where].
func (s Map1[T1, T2]) Where(filter yielder[T1]) Map1[T1, T2] {
	return Map1[T1, T2](Seq[T1](s).Where(filter))
}

// See [Seq.Skip].
func (s Map1[T1, T2]) Skip(toSkip int) Map1[T1, T2] {
	return Map1[T1, T2](Seq[T1](s).Skip(toSkip))
}

// See [Seq.SkipWhile].
func (s Map1[T1, T2]) SkipWhile(test yielder[T1]) Map1[T1, T2] {
	return Map1[T1, T2](Seq[T1](s).SkipWhile(test))
}

// See [Seq.Take].
func (s Map1[T1, T2]) Take(toTake int) Map1[T1, T2] {
	return Map1[T1, T2](Seq[T1](s).Take(toTake))
}

// See [Seq.TakeWhile].
func (s Map1[T1, T2]) TakeWhile(test yielder[T1]) Map1[T1, T2] {
	return Map1[T1, T2](Seq[T1](s).TakeWhile(test))
}
// A Map2 is a wrapper around [Seq] that provides methods to map to 2 additional types.
type Map2[T1, T2, T3 any] Map1[T1, T2]

// Map transforms the elements within the iterator using the provided mapper function.
func (s Map2[T1, T2, T3]) Map(mapper func(T1) T2) Map1[T2, T3] {
	return func(yield yielder[T2]) {
		for v := range s {
				if !yield(mapper(v)) {
					break
				}
		}
	}
}

// See [Seq.Where].
func (s Map2[T1, T2, T3]) Where(filter yielder[T1]) Map2[T1, T2, T3] {
	return Map2[T1, T2, T3](Seq[T1](s).Where(filter))
}

// See [Seq.Skip].
func (s Map2[T1, T2, T3]) Skip(toSkip int) Map2[T1, T2, T3] {
	return Map2[T1, T2, T3](Seq[T1](s).Skip(toSkip))
}

// See [Seq.SkipWhile].
func (s Map2[T1, T2, T3]) SkipWhile(test yielder[T1]) Map2[T1, T2, T3] {
	return Map2[T1, T2, T3](Seq[T1](s).SkipWhile(test))
}

// See [Seq.Take].
func (s Map2[T1, T2, T3]) Take(toTake int) Map2[T1, T2, T3] {
	return Map2[T1, T2, T3](Seq[T1](s).Take(toTake))
}

// See [Seq.TakeWhile].
func (s Map2[T1, T2, T3]) TakeWhile(test yielder[T1]) Map2[T1, T2, T3] {
	return Map2[T1, T2, T3](Seq[T1](s).TakeWhile(test))
}
// A Map3 is a wrapper around [Seq] that provides methods to map to 3 additional types.
type Map3[T1, T2, T3, T4 any] Map2[T1, T2, T3]

// Map transforms the elements within the iterator using the provided mapper function.
func (s Map3[T1, T2, T3, T4]) Map(mapper func(T1) T2) Map2[T2, T3, T4] {
	return func(yield yielder[T2]) {
		for v := range s {
				if !yield(mapper(v)) {
					break
				}
		}
	}
}

// See [Seq.Where].
func (s Map3[T1, T2, T3, T4]) Where(filter yielder[T1]) Map3[T1, T2, T3, T4] {
	return Map3[T1, T2, T3, T4](Seq[T1](s).Where(filter))
}

// See [Seq.Skip].
func (s Map3[T1, T2, T3, T4]) Skip(toSkip int) Map3[T1, T2, T3, T4] {
	return Map3[T1, T2, T3, T4](Seq[T1](s).Skip(toSkip))
}

// See [Seq.SkipWhile].
func (s Map3[T1, T2, T3, T4]) SkipWhile(test yielder[T1]) Map3[T1, T2, T3, T4] {
	return Map3[T1, T2, T3, T4](Seq[T1](s).SkipWhile(test))
}

// See [Seq.Take].
func (s Map3[T1, T2, T3, T4]) Take(toTake int) Map3[T1, T2, T3, T4] {
	return Map3[T1, T2, T3, T4](Seq[T1](s).Take(toTake))
}

// See [Seq.TakeWhile].
func (s Map3[T1, T2, T3, T4]) TakeWhile(test yielder[T1]) Map3[T1, T2, T3, T4] {
	return Map3[T1, T2, T3, T4](Seq[T1](s).TakeWhile(test))
}
// A Map4 is a wrapper around [Seq] that provides methods to map to 4 additional types.
type Map4[T1, T2, T3, T4, T5 any] Map3[T1, T2, T3, T4]

// Map transforms the elements within the iterator using the provided mapper function.
func (s Map4[T1, T2, T3, T4, T5]) Map(mapper func(T1) T2) Map3[T2, T3, T4, T5] {
	return func(yield yielder[T2]) {
		for v := range s {
				if !yield(mapper(v)) {
					break
				}
		}
	}
}

// See [Seq.Where].
func (s Map4[T1, T2, T3, T4, T5]) Where(filter yielder[T1]) Map4[T1, T2, T3, T4, T5] {
	return Map4[T1, T2, T3, T4, T5](Seq[T1](s).Where(filter))
}

// See [Seq.Skip].
func (s Map4[T1, T2, T3, T4, T5]) Skip(toSkip int) Map4[T1, T2, T3, T4, T5] {
	return Map4[T1, T2, T3, T4, T5](Seq[T1](s).Skip(toSkip))
}

// See [Seq.SkipWhile].
func (s Map4[T1, T2, T3, T4, T5]) SkipWhile(test yielder[T1]) Map4[T1, T2, T3, T4, T5] {
	return Map4[T1, T2, T3, T4, T5](Seq[T1](s).SkipWhile(test))
}

// See [Seq.Take].
func (s Map4[T1, T2, T3, T4, T5]) Take(toTake int) Map4[T1, T2, T3, T4, T5] {
	return Map4[T1, T2, T3, T4, T5](Seq[T1](s).Take(toTake))
}

// See [Seq.TakeWhile].
func (s Map4[T1, T2, T3, T4, T5]) TakeWhile(test yielder[T1]) Map4[T1, T2, T3, T4, T5] {
	return Map4[T1, T2, T3, T4, T5](Seq[T1](s).TakeWhile(test))
}
// A Map5 is a wrapper around [Seq] that provides methods to map to 5 additional types.
type Map5[T1, T2, T3, T4, T5, T6 any] Map4[T1, T2, T3, T4, T5]

// Map transforms the elements within the iterator using the provided mapper function.
func (s Map5[T1, T2, T3, T4, T5, T6]) Map(mapper func(T1) T2) Map4[T2, T3, T4, T5, T6] {
	return func(yield yielder[T2]) {
		for v := range s {
				if !yield(mapper(v)) {
					break
				}
		}
	}
}

// See [Seq.Where].
func (s Map5[T1, T2, T3, T4, T5, T6]) Where(filter yielder[T1]) Map5[T1, T2, T3, T4, T5, T6] {
	return Map5[T1, T2, T3, T4, T5, T6](Seq[T1](s).Where(filter))
}

// See [Seq.Skip].
func (s Map5[T1, T2, T3, T4, T5, T6]) Skip(toSkip int) Map5[T1, T2, T3, T4, T5, T6] {
	return Map5[T1, T2, T3, T4, T5, T6](Seq[T1](s).Skip(toSkip))
}

// See [Seq.SkipWhile].
func (s Map5[T1, T2, T3, T4, T5, T6]) SkipWhile(test yielder[T1]) Map5[T1, T2, T3, T4, T5, T6] {
	return Map5[T1, T2, T3, T4, T5, T6](Seq[T1](s).SkipWhile(test))
}

// See [Seq.Take].
func (s Map5[T1, T2, T3, T4, T5, T6]) Take(toTake int) Map5[T1, T2, T3, T4, T5, T6] {
	return Map5[T1, T2, T3, T4, T5, T6](Seq[T1](s).Take(toTake))
}

// See [Seq.TakeWhile].
func (s Map5[T1, T2, T3, T4, T5, T6]) TakeWhile(test yielder[T1]) Map5[T1, T2, T3, T4, T5, T6] {
	return Map5[T1, T2, T3, T4, T5, T6](Seq[T1](s).TakeWhile(test))
}
// A Map6 is a wrapper around [Seq] that provides methods to map to 6 additional types.
type Map6[T1, T2, T3, T4, T5, T6, T7 any] Map5[T1, T2, T3, T4, T5, T6]

// Map transforms the elements within the iterator using the provided mapper function.
func (s Map6[T1, T2, T3, T4, T5, T6, T7]) Map(mapper func(T1) T2) Map5[T2, T3, T4, T5, T6, T7] {
	return func(yield yielder[T2]) {
		for v := range s {
				if !yield(mapper(v)) {
					break
				}
		}
	}
}

// See [Seq.Where].
func (s Map6[T1, T2, T3, T4, T5, T6, T7]) Where(filter yielder[T1]) Map6[T1, T2, T3, T4, T5, T6, T7] {
	return Map6[T1, T2, T3, T4, T5, T6, T7](Seq[T1](s).Where(filter))
}

// See [Seq.Skip].
func (s Map6[T1, T2, T3, T4, T5, T6, T7]) Skip(toSkip int) Map6[T1, T2, T3, T4, T5, T6, T7] {
	return Map6[T1, T2, T3, T4, T5, T6, T7](Seq[T1](s).Skip(toSkip))
}

// See [Seq.SkipWhile].
func (s Map6[T1, T2, T3, T4, T5, T6, T7]) SkipWhile(test yielder[T1]) Map6[T1, T2, T3, T4, T5, T6, T7] {
	return Map6[T1, T2, T3, T4, T5, T6, T7](Seq[T1](s).SkipWhile(test))
}

// See [Seq.Take].
func (s Map6[T1, T2, T3, T4, T5, T6, T7]) Take(toTake int) Map6[T1, T2, T3, T4, T5, T6, T7] {
	return Map6[T1, T2, T3, T4, T5, T6, T7](Seq[T1](s).Take(toTake))
}

// See [Seq.TakeWhile].
func (s Map6[T1, T2, T3, T4, T5, T6, T7]) TakeWhile(test yielder[T1]) Map6[T1, T2, T3, T4, T5, T6, T7] {
	return Map6[T1, T2, T3, T4, T5, T6, T7](Seq[T1](s).TakeWhile(test))
}
// A Map7 is a wrapper around [Seq] that provides methods to map to 7 additional types.
type Map7[T1, T2, T3, T4, T5, T6, T7, T8 any] Map6[T1, T2, T3, T4, T5, T6, T7]

// Map transforms the elements within the iterator using the provided mapper function.
func (s Map7[T1, T2, T3, T4, T5, T6, T7, T8]) Map(mapper func(T1) T2) Map6[T2, T3, T4, T5, T6, T7, T8] {
	return func(yield yielder[T2]) {
		for v := range s {
				if !yield(mapper(v)) {
					break
				}
		}
	}
}

// See [Seq.Where].
func (s Map7[T1, T2, T3, T4, T5, T6, T7, T8]) Where(filter yielder[T1]) Map7[T1, T2, T3, T4, T5, T6, T7, T8] {
	return Map7[T1, T2, T3, T4, T5, T6, T7, T8](Seq[T1](s).Where(filter))
}

// See [Seq.Skip].
func (s Map7[T1, T2, T3, T4, T5, T6, T7, T8]) Skip(toSkip int) Map7[T1, T2, T3, T4, T5, T6, T7, T8] {
	return Map7[T1, T2, T3, T4, T5, T6, T7, T8](Seq[T1](s).Skip(toSkip))
}

// See [Seq.SkipWhile].
func (s Map7[T1, T2, T3, T4, T5, T6, T7, T8]) SkipWhile(test yielder[T1]) Map7[T1, T2, T3, T4, T5, T6, T7, T8] {
	return Map7[T1, T2, T3, T4, T5, T6, T7, T8](Seq[T1](s).SkipWhile(test))
}

// See [Seq.Take].
func (s Map7[T1, T2, T3, T4, T5, T6, T7, T8]) Take(toTake int) Map7[T1, T2, T3, T4, T5, T6, T7, T8] {
	return Map7[T1, T2, T3, T4, T5, T6, T7, T8](Seq[T1](s).Take(toTake))
}

// See [Seq.TakeWhile].
func (s Map7[T1, T2, T3, T4, T5, T6, T7, T8]) TakeWhile(test yielder[T1]) Map7[T1, T2, T3, T4, T5, T6, T7, T8] {
	return Map7[T1, T2, T3, T4, T5, T6, T7, T8](Seq[T1](s).TakeWhile(test))
}
// A Map8 is a wrapper around [Seq] that provides methods to map to 8 additional types.
type Map8[T1, T2, T3, T4, T5, T6, T7, T8, T9 any] Map7[T1, T2, T3, T4, T5, T6, T7, T8]

// Map transforms the elements within the iterator using the provided mapper function.
func (s Map8[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Map(mapper func(T1) T2) Map7[T2, T3, T4, T5, T6, T7, T8, T9] {
	return func(yield yielder[T2]) {
		for v := range s {
				if !yield(mapper(v)) {
					break
				}
		}
	}
}

// See [Seq.Where].
func (s Map8[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Where(filter yielder[T1]) Map8[T1, T2, T3, T4, T5, T6, T7, T8, T9] {
	return Map8[T1, T2, T3, T4, T5, T6, T7, T8, T9](Seq[T1](s).Where(filter))
}

// See [Seq.Skip].
func (s Map8[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Skip(toSkip int) Map8[T1, T2, T3, T4, T5, T6, T7, T8, T9] {
	return Map8[T1, T2, T3, T4, T5, T6, T7, T8, T9](Seq[T1](s).Skip(toSkip))
}

// See [Seq.SkipWhile].
func (s Map8[T1, T2, T3, T4, T5, T6, T7, T8, T9]) SkipWhile(test yielder[T1]) Map8[T1, T2, T3, T4, T5, T6, T7, T8, T9] {
	return Map8[T1, T2, T3, T4, T5, T6, T7, T8, T9](Seq[T1](s).SkipWhile(test))
}

// See [Seq.Take].
func (s Map8[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Take(toTake int) Map8[T1, T2, T3, T4, T5, T6, T7, T8, T9] {
	return Map8[T1, T2, T3, T4, T5, T6, T7, T8, T9](Seq[T1](s).Take(toTake))
}

// See [Seq.TakeWhile].
func (s Map8[T1, T2, T3, T4, T5, T6, T7, T8, T9]) TakeWhile(test yielder[T1]) Map8[T1, T2, T3, T4, T5, T6, T7, T8, T9] {
	return Map8[T1, T2, T3, T4, T5, T6, T7, T8, T9](Seq[T1](s).TakeWhile(test))
}
// A Map9 is a wrapper around [Seq] that provides methods to map to 9 additional types.
type Map9[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any] Map8[T1, T2, T3, T4, T5, T6, T7, T8, T9]

// Map transforms the elements within the iterator using the provided mapper function.
func (s Map9[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Map(mapper func(T1) T2) Map8[T2, T3, T4, T5, T6, T7, T8, T9, T10] {
	return func(yield yielder[T2]) {
		for v := range s {
				if !yield(mapper(v)) {
					break
				}
		}
	}
}

// See [Seq.Where].
func (s Map9[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Where(filter yielder[T1]) Map9[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10] {
	return Map9[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10](Seq[T1](s).Where(filter))
}

// See [Seq.Skip].
func (s Map9[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Skip(toSkip int) Map9[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10] {
	return Map9[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10](Seq[T1](s).Skip(toSkip))
}

// See [Seq.SkipWhile].
func (s Map9[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) SkipWhile(test yielder[T1]) Map9[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10] {
	return Map9[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10](Seq[T1](s).SkipWhile(test))
}

// See [Seq.Take].
func (s Map9[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Take(toTake int) Map9[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10] {
	return Map9[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10](Seq[T1](s).Take(toTake))
}

// See [Seq.TakeWhile].
func (s Map9[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) TakeWhile(test yielder[T1]) Map9[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10] {
	return Map9[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10](Seq[T1](s).TakeWhile(test))
}

