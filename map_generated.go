package loz

// A Map1 is a wrapper around [Seq] that provides methods to map to 1 additional type.
type Map1[V1, V2 any] Seq[V1]

// Map transforms the elements within the iterator using the provided mapper function.
func (s Map1[V1, V2]) Map(mapper func(V1) V2) Seq[V2] {
	return func(yield yielder[V2]) {
		for v := range s {
			if !yield(mapper(v)) {
				break
			}
		}
	}
}

func (s Map1[V1, V2]) Expand(toElements func(V1) Seq[V2]) Seq[V2] {
	return func(yield yielder[V2]) {
		for v := range s {
			for e := range toElements(v) {
				if !yield(e) {
					break
				}
			}
		}
	}
}

// See [Seq.Filter].
func (s Map1[V1, V2]) Filter(filter yielder[V1]) Map1[V1, V2] {
	return Map1[V1, V2](Seq[V1](s).Filter(filter))
}

// See [Seq.Skip].
func (s Map1[V1, V2]) Skip(toSkip int) Map1[V1, V2] {
	return Map1[V1, V2](Seq[V1](s).Skip(toSkip))
}

// See [Seq.SkipWhile].
func (s Map1[V1, V2]) SkipWhile(test yielder[V1]) Map1[V1, V2] {
	return Map1[V1, V2](Seq[V1](s).SkipWhile(test))
}

// See [Seq.Take].
func (s Map1[V1, V2]) Take(toTake int) Map1[V1, V2] {
	return Map1[V1, V2](Seq[V1](s).Take(toTake))
}

// See [Seq.TakeWhile].
func (s Map1[V1, V2]) TakeWhile(test yielder[V1]) Map1[V1, V2] {
	return Map1[V1, V2](Seq[V1](s).TakeWhile(test))
}

// A Map2 is a wrapper around [Seq] that provides methods to map to 2 additional types.
type Map2[V1, V2, V3 any] Map1[V1, V2]

// Map transforms the elements within the iterator using the provided mapper function.
func (s Map2[V1, V2, V3]) Map(mapper func(V1) V2) Map1[V2, V3] {
	return func(yield yielder[V2]) {
		for v := range s {
			if !yield(mapper(v)) {
				break
			}
		}
	}
}

func (s Map2[V1, V2, V3]) Expand(toElements func(V1) Seq[V2]) Map1[V2, V3] {
	return func(yield yielder[V2]) {
		for v := range s {
			for e := range toElements(v) {
				if !yield(e) {
					break
				}
			}
		}
	}
}

// See [Seq.Filter].
func (s Map2[V1, V2, V3]) Filter(filter yielder[V1]) Map2[V1, V2, V3] {
	return Map2[V1, V2, V3](Seq[V1](s).Filter(filter))
}

// See [Seq.Skip].
func (s Map2[V1, V2, V3]) Skip(toSkip int) Map2[V1, V2, V3] {
	return Map2[V1, V2, V3](Seq[V1](s).Skip(toSkip))
}

// See [Seq.SkipWhile].
func (s Map2[V1, V2, V3]) SkipWhile(test yielder[V1]) Map2[V1, V2, V3] {
	return Map2[V1, V2, V3](Seq[V1](s).SkipWhile(test))
}

// See [Seq.Take].
func (s Map2[V1, V2, V3]) Take(toTake int) Map2[V1, V2, V3] {
	return Map2[V1, V2, V3](Seq[V1](s).Take(toTake))
}

// See [Seq.TakeWhile].
func (s Map2[V1, V2, V3]) TakeWhile(test yielder[V1]) Map2[V1, V2, V3] {
	return Map2[V1, V2, V3](Seq[V1](s).TakeWhile(test))
}

// A Map3 is a wrapper around [Seq] that provides methods to map to 3 additional types.
type Map3[V1, V2, V3, V4 any] Map2[V1, V2, V3]

// Map transforms the elements within the iterator using the provided mapper function.
func (s Map3[V1, V2, V3, V4]) Map(mapper func(V1) V2) Map2[V2, V3, V4] {
	return func(yield yielder[V2]) {
		for v := range s {
			if !yield(mapper(v)) {
				break
			}
		}
	}
}

func (s Map3[V1, V2, V3, V4]) Expand(toElements func(V1) Seq[V2]) Map2[V2, V3, V4] {
	return func(yield yielder[V2]) {
		for v := range s {
			for e := range toElements(v) {
				if !yield(e) {
					break
				}
			}
		}
	}
}

// See [Seq.Filter].
func (s Map3[V1, V2, V3, V4]) Filter(filter yielder[V1]) Map3[V1, V2, V3, V4] {
	return Map3[V1, V2, V3, V4](Seq[V1](s).Filter(filter))
}

// See [Seq.Skip].
func (s Map3[V1, V2, V3, V4]) Skip(toSkip int) Map3[V1, V2, V3, V4] {
	return Map3[V1, V2, V3, V4](Seq[V1](s).Skip(toSkip))
}

// See [Seq.SkipWhile].
func (s Map3[V1, V2, V3, V4]) SkipWhile(test yielder[V1]) Map3[V1, V2, V3, V4] {
	return Map3[V1, V2, V3, V4](Seq[V1](s).SkipWhile(test))
}

// See [Seq.Take].
func (s Map3[V1, V2, V3, V4]) Take(toTake int) Map3[V1, V2, V3, V4] {
	return Map3[V1, V2, V3, V4](Seq[V1](s).Take(toTake))
}

// See [Seq.TakeWhile].
func (s Map3[V1, V2, V3, V4]) TakeWhile(test yielder[V1]) Map3[V1, V2, V3, V4] {
	return Map3[V1, V2, V3, V4](Seq[V1](s).TakeWhile(test))
}

// A Map4 is a wrapper around [Seq] that provides methods to map to 4 additional types.
type Map4[V1, V2, V3, V4, V5 any] Map3[V1, V2, V3, V4]

// Map transforms the elements within the iterator using the provided mapper function.
func (s Map4[V1, V2, V3, V4, V5]) Map(mapper func(V1) V2) Map3[V2, V3, V4, V5] {
	return func(yield yielder[V2]) {
		for v := range s {
			if !yield(mapper(v)) {
				break
			}
		}
	}
}

func (s Map4[V1, V2, V3, V4, V5]) Expand(toElements func(V1) Seq[V2]) Map3[V2, V3, V4, V5] {
	return func(yield yielder[V2]) {
		for v := range s {
			for e := range toElements(v) {
				if !yield(e) {
					break
				}
			}
		}
	}
}

// See [Seq.Filter].
func (s Map4[V1, V2, V3, V4, V5]) Filter(filter yielder[V1]) Map4[V1, V2, V3, V4, V5] {
	return Map4[V1, V2, V3, V4, V5](Seq[V1](s).Filter(filter))
}

// See [Seq.Skip].
func (s Map4[V1, V2, V3, V4, V5]) Skip(toSkip int) Map4[V1, V2, V3, V4, V5] {
	return Map4[V1, V2, V3, V4, V5](Seq[V1](s).Skip(toSkip))
}

// See [Seq.SkipWhile].
func (s Map4[V1, V2, V3, V4, V5]) SkipWhile(test yielder[V1]) Map4[V1, V2, V3, V4, V5] {
	return Map4[V1, V2, V3, V4, V5](Seq[V1](s).SkipWhile(test))
}

// See [Seq.Take].
func (s Map4[V1, V2, V3, V4, V5]) Take(toTake int) Map4[V1, V2, V3, V4, V5] {
	return Map4[V1, V2, V3, V4, V5](Seq[V1](s).Take(toTake))
}

// See [Seq.TakeWhile].
func (s Map4[V1, V2, V3, V4, V5]) TakeWhile(test yielder[V1]) Map4[V1, V2, V3, V4, V5] {
	return Map4[V1, V2, V3, V4, V5](Seq[V1](s).TakeWhile(test))
}

// A Map5 is a wrapper around [Seq] that provides methods to map to 5 additional types.
type Map5[V1, V2, V3, V4, V5, V6 any] Map4[V1, V2, V3, V4, V5]

// Map transforms the elements within the iterator using the provided mapper function.
func (s Map5[V1, V2, V3, V4, V5, V6]) Map(mapper func(V1) V2) Map4[V2, V3, V4, V5, V6] {
	return func(yield yielder[V2]) {
		for v := range s {
			if !yield(mapper(v)) {
				break
			}
		}
	}
}

func (s Map5[V1, V2, V3, V4, V5, V6]) Expand(toElements func(V1) Seq[V2]) Map4[V2, V3, V4, V5, V6] {
	return func(yield yielder[V2]) {
		for v := range s {
			for e := range toElements(v) {
				if !yield(e) {
					break
				}
			}
		}
	}
}

// See [Seq.Filter].
func (s Map5[V1, V2, V3, V4, V5, V6]) Filter(filter yielder[V1]) Map5[V1, V2, V3, V4, V5, V6] {
	return Map5[V1, V2, V3, V4, V5, V6](Seq[V1](s).Filter(filter))
}

// See [Seq.Skip].
func (s Map5[V1, V2, V3, V4, V5, V6]) Skip(toSkip int) Map5[V1, V2, V3, V4, V5, V6] {
	return Map5[V1, V2, V3, V4, V5, V6](Seq[V1](s).Skip(toSkip))
}

// See [Seq.SkipWhile].
func (s Map5[V1, V2, V3, V4, V5, V6]) SkipWhile(test yielder[V1]) Map5[V1, V2, V3, V4, V5, V6] {
	return Map5[V1, V2, V3, V4, V5, V6](Seq[V1](s).SkipWhile(test))
}

// See [Seq.Take].
func (s Map5[V1, V2, V3, V4, V5, V6]) Take(toTake int) Map5[V1, V2, V3, V4, V5, V6] {
	return Map5[V1, V2, V3, V4, V5, V6](Seq[V1](s).Take(toTake))
}

// See [Seq.TakeWhile].
func (s Map5[V1, V2, V3, V4, V5, V6]) TakeWhile(test yielder[V1]) Map5[V1, V2, V3, V4, V5, V6] {
	return Map5[V1, V2, V3, V4, V5, V6](Seq[V1](s).TakeWhile(test))
}

// A Map6 is a wrapper around [Seq] that provides methods to map to 6 additional types.
type Map6[V1, V2, V3, V4, V5, V6, V7 any] Map5[V1, V2, V3, V4, V5, V6]

// Map transforms the elements within the iterator using the provided mapper function.
func (s Map6[V1, V2, V3, V4, V5, V6, V7]) Map(mapper func(V1) V2) Map5[V2, V3, V4, V5, V6, V7] {
	return func(yield yielder[V2]) {
		for v := range s {
			if !yield(mapper(v)) {
				break
			}
		}
	}
}

func (s Map6[V1, V2, V3, V4, V5, V6, V7]) Expand(toElements func(V1) Seq[V2]) Map5[V2, V3, V4, V5, V6, V7] {
	return func(yield yielder[V2]) {
		for v := range s {
			for e := range toElements(v) {
				if !yield(e) {
					break
				}
			}
		}
	}
}

// See [Seq.Filter].
func (s Map6[V1, V2, V3, V4, V5, V6, V7]) Filter(filter yielder[V1]) Map6[V1, V2, V3, V4, V5, V6, V7] {
	return Map6[V1, V2, V3, V4, V5, V6, V7](Seq[V1](s).Filter(filter))
}

// See [Seq.Skip].
func (s Map6[V1, V2, V3, V4, V5, V6, V7]) Skip(toSkip int) Map6[V1, V2, V3, V4, V5, V6, V7] {
	return Map6[V1, V2, V3, V4, V5, V6, V7](Seq[V1](s).Skip(toSkip))
}

// See [Seq.SkipWhile].
func (s Map6[V1, V2, V3, V4, V5, V6, V7]) SkipWhile(test yielder[V1]) Map6[V1, V2, V3, V4, V5, V6, V7] {
	return Map6[V1, V2, V3, V4, V5, V6, V7](Seq[V1](s).SkipWhile(test))
}

// See [Seq.Take].
func (s Map6[V1, V2, V3, V4, V5, V6, V7]) Take(toTake int) Map6[V1, V2, V3, V4, V5, V6, V7] {
	return Map6[V1, V2, V3, V4, V5, V6, V7](Seq[V1](s).Take(toTake))
}

// See [Seq.TakeWhile].
func (s Map6[V1, V2, V3, V4, V5, V6, V7]) TakeWhile(test yielder[V1]) Map6[V1, V2, V3, V4, V5, V6, V7] {
	return Map6[V1, V2, V3, V4, V5, V6, V7](Seq[V1](s).TakeWhile(test))
}

// A Map7 is a wrapper around [Seq] that provides methods to map to 7 additional types.
type Map7[V1, V2, V3, V4, V5, V6, V7, V8 any] Map6[V1, V2, V3, V4, V5, V6, V7]

// Map transforms the elements within the iterator using the provided mapper function.
func (s Map7[V1, V2, V3, V4, V5, V6, V7, V8]) Map(mapper func(V1) V2) Map6[V2, V3, V4, V5, V6, V7, V8] {
	return func(yield yielder[V2]) {
		for v := range s {
			if !yield(mapper(v)) {
				break
			}
		}
	}
}

func (s Map7[V1, V2, V3, V4, V5, V6, V7, V8]) Expand(toElements func(V1) Seq[V2]) Map6[V2, V3, V4, V5, V6, V7, V8] {
	return func(yield yielder[V2]) {
		for v := range s {
			for e := range toElements(v) {
				if !yield(e) {
					break
				}
			}
		}
	}
}

// See [Seq.Filter].
func (s Map7[V1, V2, V3, V4, V5, V6, V7, V8]) Filter(filter yielder[V1]) Map7[V1, V2, V3, V4, V5, V6, V7, V8] {
	return Map7[V1, V2, V3, V4, V5, V6, V7, V8](Seq[V1](s).Filter(filter))
}

// See [Seq.Skip].
func (s Map7[V1, V2, V3, V4, V5, V6, V7, V8]) Skip(toSkip int) Map7[V1, V2, V3, V4, V5, V6, V7, V8] {
	return Map7[V1, V2, V3, V4, V5, V6, V7, V8](Seq[V1](s).Skip(toSkip))
}

// See [Seq.SkipWhile].
func (s Map7[V1, V2, V3, V4, V5, V6, V7, V8]) SkipWhile(test yielder[V1]) Map7[V1, V2, V3, V4, V5, V6, V7, V8] {
	return Map7[V1, V2, V3, V4, V5, V6, V7, V8](Seq[V1](s).SkipWhile(test))
}

// See [Seq.Take].
func (s Map7[V1, V2, V3, V4, V5, V6, V7, V8]) Take(toTake int) Map7[V1, V2, V3, V4, V5, V6, V7, V8] {
	return Map7[V1, V2, V3, V4, V5, V6, V7, V8](Seq[V1](s).Take(toTake))
}

// See [Seq.TakeWhile].
func (s Map7[V1, V2, V3, V4, V5, V6, V7, V8]) TakeWhile(test yielder[V1]) Map7[V1, V2, V3, V4, V5, V6, V7, V8] {
	return Map7[V1, V2, V3, V4, V5, V6, V7, V8](Seq[V1](s).TakeWhile(test))
}

// A Map8 is a wrapper around [Seq] that provides methods to map to 8 additional types.
type Map8[V1, V2, V3, V4, V5, V6, V7, V8, V9 any] Map7[V1, V2, V3, V4, V5, V6, V7, V8]

// Map transforms the elements within the iterator using the provided mapper function.
func (s Map8[V1, V2, V3, V4, V5, V6, V7, V8, V9]) Map(mapper func(V1) V2) Map7[V2, V3, V4, V5, V6, V7, V8, V9] {
	return func(yield yielder[V2]) {
		for v := range s {
			if !yield(mapper(v)) {
				break
			}
		}
	}
}

func (s Map8[V1, V2, V3, V4, V5, V6, V7, V8, V9]) Expand(toElements func(V1) Seq[V2]) Map7[V2, V3, V4, V5, V6, V7, V8, V9] {
	return func(yield yielder[V2]) {
		for v := range s {
			for e := range toElements(v) {
				if !yield(e) {
					break
				}
			}
		}
	}
}

// See [Seq.Filter].
func (s Map8[V1, V2, V3, V4, V5, V6, V7, V8, V9]) Filter(filter yielder[V1]) Map8[V1, V2, V3, V4, V5, V6, V7, V8, V9] {
	return Map8[V1, V2, V3, V4, V5, V6, V7, V8, V9](Seq[V1](s).Filter(filter))
}

// See [Seq.Skip].
func (s Map8[V1, V2, V3, V4, V5, V6, V7, V8, V9]) Skip(toSkip int) Map8[V1, V2, V3, V4, V5, V6, V7, V8, V9] {
	return Map8[V1, V2, V3, V4, V5, V6, V7, V8, V9](Seq[V1](s).Skip(toSkip))
}

// See [Seq.SkipWhile].
func (s Map8[V1, V2, V3, V4, V5, V6, V7, V8, V9]) SkipWhile(test yielder[V1]) Map8[V1, V2, V3, V4, V5, V6, V7, V8, V9] {
	return Map8[V1, V2, V3, V4, V5, V6, V7, V8, V9](Seq[V1](s).SkipWhile(test))
}

// See [Seq.Take].
func (s Map8[V1, V2, V3, V4, V5, V6, V7, V8, V9]) Take(toTake int) Map8[V1, V2, V3, V4, V5, V6, V7, V8, V9] {
	return Map8[V1, V2, V3, V4, V5, V6, V7, V8, V9](Seq[V1](s).Take(toTake))
}

// See [Seq.TakeWhile].
func (s Map8[V1, V2, V3, V4, V5, V6, V7, V8, V9]) TakeWhile(test yielder[V1]) Map8[V1, V2, V3, V4, V5, V6, V7, V8, V9] {
	return Map8[V1, V2, V3, V4, V5, V6, V7, V8, V9](Seq[V1](s).TakeWhile(test))
}

// A Map9 is a wrapper around [Seq] that provides methods to map to 9 additional types.
type Map9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10 any] Map8[V1, V2, V3, V4, V5, V6, V7, V8, V9]

// Map transforms the elements within the iterator using the provided mapper function.
func (s Map9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10]) Map(mapper func(V1) V2) Map8[V2, V3, V4, V5, V6, V7, V8, V9, V10] {
	return func(yield yielder[V2]) {
		for v := range s {
			if !yield(mapper(v)) {
				break
			}
		}
	}
}

func (s Map9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10]) Expand(toElements func(V1) Seq[V2]) Map8[V2, V3, V4, V5, V6, V7, V8, V9, V10] {
	return func(yield yielder[V2]) {
		for v := range s {
			for e := range toElements(v) {
				if !yield(e) {
					break
				}
			}
		}
	}
}

// See [Seq.Filter].
func (s Map9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10]) Filter(filter yielder[V1]) Map9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10] {
	return Map9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10](Seq[V1](s).Filter(filter))
}

// See [Seq.Skip].
func (s Map9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10]) Skip(toSkip int) Map9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10] {
	return Map9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10](Seq[V1](s).Skip(toSkip))
}

// See [Seq.SkipWhile].
func (s Map9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10]) SkipWhile(test yielder[V1]) Map9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10] {
	return Map9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10](Seq[V1](s).SkipWhile(test))
}

// See [Seq.Take].
func (s Map9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10]) Take(toTake int) Map9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10] {
	return Map9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10](Seq[V1](s).Take(toTake))
}

// See [Seq.TakeWhile].
func (s Map9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10]) TakeWhile(test yielder[V1]) Map9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10] {
	return Map9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10](Seq[V1](s).TakeWhile(test))
}
