package loz

type Mapper1[V1, V2 any] Seq[V1]

// Map transforms the elements within the iterator using the provided mapper function.
func (s Mapper1[V1, V2]) Map(mapper mapper[V1, V2]) Seq[V2] {
	return func(yield yielder[V2]) {
		for v := range s {
			if !yield(mapper(v)) {
				break
			}
		}
	}
}

func (s Mapper1[V1, V2]) Expand(toElements mapper[V1, Seq[V2]]) Seq[V2] {
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

func IterSliceMap1[V1, V2 any](slice []V1) Mapper1[V1, V2] {
	return Mapper1[V1, V2](IterSlice(slice))
}

// See [Seq.Filter].
func (s Mapper1[V1, V2]) Filter(filter yielder[V1]) Mapper1[V1, V2] {
	return Mapper1[V1, V2](Seq[V1](s).Filter(filter))
}

// See [Seq.Skip].
func (s Mapper1[V1, V2]) Skip(toSkip int) Mapper1[V1, V2] {
	return Mapper1[V1, V2](Seq[V1](s).Skip(toSkip))
}

// See [Seq.SkipWhile].
func (s Mapper1[V1, V2]) SkipWhile(test yielder[V1]) Mapper1[V1, V2] {
	return Mapper1[V1, V2](Seq[V1](s).SkipWhile(test))
}

// See [Seq.Take].
func (s Mapper1[V1, V2]) Take(toTake int) Mapper1[V1, V2] {
	return Mapper1[V1, V2](Seq[V1](s).Take(toTake))
}

// See [Seq.TakeWhile].
func (s Mapper1[V1, V2]) TakeWhile(test yielder[V1]) Mapper1[V1, V2] {
	return Mapper1[V1, V2](Seq[V1](s).TakeWhile(test))
}

type KVMapper1[K1, V1, K2, V2 any] Seq2[K1, V1]

// Map transforms the keys and values within the iterator using the provided mapper function.
func (s KVMapper1[K1, V1, K2, V2]) Map(mapper mapper2[K1, V1, K2, V2]) Seq2[K2, V2] {
	return func(yield yielder2[K2, V2]) {
		for k, v := range s {
			if !yield(mapper(k, v)) {
				break
			}
		}
	}
}

func IterMapMap1[K1 comparable, V1, K2, V2 any](m map[K1]V1) KVMapper1[K1, V1, K2, V2] {
	return KVMapper1[K1, V1, K2, V2](IterMap(m))
}

// See [Seq2.Filter].
func (s KVMapper1[K1, V1, K2, V2]) Filter(filter yielder2[K1, V1]) KVMapper1[K1, V1, K2, V2] {
	return KVMapper1[K1, V1, K2, V2](Seq2[K1, V1](s).Filter(filter))
}

// See [Seq2.Skip].
func (s KVMapper1[K1, V1, K2, V2]) Skip(toSkip int) KVMapper1[K1, V1, K2, V2] {
	return KVMapper1[K1, V1, K2, V2](Seq2[K1, V1](s).Skip(toSkip))
}

// See [Seq2.SkipWhile].
func (s KVMapper1[K1, V1, K2, V2]) SkipWhile(test yielder2[K1, V1]) KVMapper1[K1, V1, K2, V2] {
	return KVMapper1[K1, V1, K2, V2](Seq2[K1, V1](s).SkipWhile(test))
}

// See [Seq2.Take].
func (s KVMapper1[K1, V1, K2, V2]) Take(toTake int) KVMapper1[K1, V1, K2, V2] {
	return KVMapper1[K1, V1, K2, V2](Seq2[K1, V1](s).Take(toTake))
}

// See [Seq2.TakeWhile].
func (s KVMapper1[K1, V1, K2, V2]) TakeWhile(test yielder2[K1, V1]) KVMapper1[K1, V1, K2, V2] {
	return KVMapper1[K1, V1, K2, V2](Seq2[K1, V1](s).TakeWhile(test))
}

type Mapper2[V1, V2, V3 any] Mapper1[V1, V2]

// Map transforms the elements within the iterator using the provided mapper function.
func (s Mapper2[V1, V2, V3]) Map(mapper mapper[V1, V2]) Mapper1[V2, V3] {
	return func(yield yielder[V2]) {
		for v := range s {
			if !yield(mapper(v)) {
				break
			}
		}
	}
}

func (s Mapper2[V1, V2, V3]) Expand(toElements mapper[V1, Seq[V2]]) Mapper1[V2, V3] {
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

func IterSliceMap2[V1, V2, V3 any](slice []V1) Mapper2[V1, V2, V3] {
	return Mapper2[V1, V2, V3](IterSlice(slice))
}

// See [Seq.Filter].
func (s Mapper2[V1, V2, V3]) Filter(filter yielder[V1]) Mapper2[V1, V2, V3] {
	return Mapper2[V1, V2, V3](Seq[V1](s).Filter(filter))
}

// See [Seq.Skip].
func (s Mapper2[V1, V2, V3]) Skip(toSkip int) Mapper2[V1, V2, V3] {
	return Mapper2[V1, V2, V3](Seq[V1](s).Skip(toSkip))
}

// See [Seq.SkipWhile].
func (s Mapper2[V1, V2, V3]) SkipWhile(test yielder[V1]) Mapper2[V1, V2, V3] {
	return Mapper2[V1, V2, V3](Seq[V1](s).SkipWhile(test))
}

// See [Seq.Take].
func (s Mapper2[V1, V2, V3]) Take(toTake int) Mapper2[V1, V2, V3] {
	return Mapper2[V1, V2, V3](Seq[V1](s).Take(toTake))
}

// See [Seq.TakeWhile].
func (s Mapper2[V1, V2, V3]) TakeWhile(test yielder[V1]) Mapper2[V1, V2, V3] {
	return Mapper2[V1, V2, V3](Seq[V1](s).TakeWhile(test))
}

type KVMapper2[K1, V1, K2, V2, K3, V3 any] KVMapper1[K1, V1, K2, V2]

// Map transforms the keys and values within the iterator using the provided mapper function.
func (s KVMapper2[K1, V1, K2, V2, K3, V3]) Map(mapper mapper2[K1, V1, K2, V2]) KVMapper1[K2, V2, K3, V3] {
	return func(yield yielder2[K2, V2]) {
		for k, v := range s {
			if !yield(mapper(k, v)) {
				break
			}
		}
	}
}

func IterMapMap2[K1 comparable, V1, K2, V2, K3, V3 any](m map[K1]V1) KVMapper2[K1, V1, K2, V2, K3, V3] {
	return KVMapper2[K1, V1, K2, V2, K3, V3](IterMap(m))
}

// See [Seq2.Filter].
func (s KVMapper2[K1, V1, K2, V2, K3, V3]) Filter(filter yielder2[K1, V1]) KVMapper2[K1, V1, K2, V2, K3, V3] {
	return KVMapper2[K1, V1, K2, V2, K3, V3](Seq2[K1, V1](s).Filter(filter))
}

// See [Seq2.Skip].
func (s KVMapper2[K1, V1, K2, V2, K3, V3]) Skip(toSkip int) KVMapper2[K1, V1, K2, V2, K3, V3] {
	return KVMapper2[K1, V1, K2, V2, K3, V3](Seq2[K1, V1](s).Skip(toSkip))
}

// See [Seq2.SkipWhile].
func (s KVMapper2[K1, V1, K2, V2, K3, V3]) SkipWhile(test yielder2[K1, V1]) KVMapper2[K1, V1, K2, V2, K3, V3] {
	return KVMapper2[K1, V1, K2, V2, K3, V3](Seq2[K1, V1](s).SkipWhile(test))
}

// See [Seq2.Take].
func (s KVMapper2[K1, V1, K2, V2, K3, V3]) Take(toTake int) KVMapper2[K1, V1, K2, V2, K3, V3] {
	return KVMapper2[K1, V1, K2, V2, K3, V3](Seq2[K1, V1](s).Take(toTake))
}

// See [Seq2.TakeWhile].
func (s KVMapper2[K1, V1, K2, V2, K3, V3]) TakeWhile(test yielder2[K1, V1]) KVMapper2[K1, V1, K2, V2, K3, V3] {
	return KVMapper2[K1, V1, K2, V2, K3, V3](Seq2[K1, V1](s).TakeWhile(test))
}

type Mapper3[V1, V2, V3, V4 any] Mapper2[V1, V2, V3]

// Map transforms the elements within the iterator using the provided mapper function.
func (s Mapper3[V1, V2, V3, V4]) Map(mapper mapper[V1, V2]) Mapper2[V2, V3, V4] {
	return func(yield yielder[V2]) {
		for v := range s {
			if !yield(mapper(v)) {
				break
			}
		}
	}
}

func (s Mapper3[V1, V2, V3, V4]) Expand(toElements mapper[V1, Seq[V2]]) Mapper2[V2, V3, V4] {
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

func IterSliceMap3[V1, V2, V3, V4 any](slice []V1) Mapper3[V1, V2, V3, V4] {
	return Mapper3[V1, V2, V3, V4](IterSlice(slice))
}

// See [Seq.Filter].
func (s Mapper3[V1, V2, V3, V4]) Filter(filter yielder[V1]) Mapper3[V1, V2, V3, V4] {
	return Mapper3[V1, V2, V3, V4](Seq[V1](s).Filter(filter))
}

// See [Seq.Skip].
func (s Mapper3[V1, V2, V3, V4]) Skip(toSkip int) Mapper3[V1, V2, V3, V4] {
	return Mapper3[V1, V2, V3, V4](Seq[V1](s).Skip(toSkip))
}

// See [Seq.SkipWhile].
func (s Mapper3[V1, V2, V3, V4]) SkipWhile(test yielder[V1]) Mapper3[V1, V2, V3, V4] {
	return Mapper3[V1, V2, V3, V4](Seq[V1](s).SkipWhile(test))
}

// See [Seq.Take].
func (s Mapper3[V1, V2, V3, V4]) Take(toTake int) Mapper3[V1, V2, V3, V4] {
	return Mapper3[V1, V2, V3, V4](Seq[V1](s).Take(toTake))
}

// See [Seq.TakeWhile].
func (s Mapper3[V1, V2, V3, V4]) TakeWhile(test yielder[V1]) Mapper3[V1, V2, V3, V4] {
	return Mapper3[V1, V2, V3, V4](Seq[V1](s).TakeWhile(test))
}

type KVMapper3[K1, V1, K2, V2, K3, V3, K4, V4 any] KVMapper2[K1, V1, K2, V2, K3, V3]

// Map transforms the keys and values within the iterator using the provided mapper function.
func (s KVMapper3[K1, V1, K2, V2, K3, V3, K4, V4]) Map(mapper mapper2[K1, V1, K2, V2]) KVMapper2[K2, V2, K3, V3, K4, V4] {
	return func(yield yielder2[K2, V2]) {
		for k, v := range s {
			if !yield(mapper(k, v)) {
				break
			}
		}
	}
}

func IterMapMap3[K1 comparable, V1, K2, V2, K3, V3, K4, V4 any](m map[K1]V1) KVMapper3[K1, V1, K2, V2, K3, V3, K4, V4] {
	return KVMapper3[K1, V1, K2, V2, K3, V3, K4, V4](IterMap(m))
}

// See [Seq2.Filter].
func (s KVMapper3[K1, V1, K2, V2, K3, V3, K4, V4]) Filter(filter yielder2[K1, V1]) KVMapper3[K1, V1, K2, V2, K3, V3, K4, V4] {
	return KVMapper3[K1, V1, K2, V2, K3, V3, K4, V4](Seq2[K1, V1](s).Filter(filter))
}

// See [Seq2.Skip].
func (s KVMapper3[K1, V1, K2, V2, K3, V3, K4, V4]) Skip(toSkip int) KVMapper3[K1, V1, K2, V2, K3, V3, K4, V4] {
	return KVMapper3[K1, V1, K2, V2, K3, V3, K4, V4](Seq2[K1, V1](s).Skip(toSkip))
}

// See [Seq2.SkipWhile].
func (s KVMapper3[K1, V1, K2, V2, K3, V3, K4, V4]) SkipWhile(test yielder2[K1, V1]) KVMapper3[K1, V1, K2, V2, K3, V3, K4, V4] {
	return KVMapper3[K1, V1, K2, V2, K3, V3, K4, V4](Seq2[K1, V1](s).SkipWhile(test))
}

// See [Seq2.Take].
func (s KVMapper3[K1, V1, K2, V2, K3, V3, K4, V4]) Take(toTake int) KVMapper3[K1, V1, K2, V2, K3, V3, K4, V4] {
	return KVMapper3[K1, V1, K2, V2, K3, V3, K4, V4](Seq2[K1, V1](s).Take(toTake))
}

// See [Seq2.TakeWhile].
func (s KVMapper3[K1, V1, K2, V2, K3, V3, K4, V4]) TakeWhile(test yielder2[K1, V1]) KVMapper3[K1, V1, K2, V2, K3, V3, K4, V4] {
	return KVMapper3[K1, V1, K2, V2, K3, V3, K4, V4](Seq2[K1, V1](s).TakeWhile(test))
}

type Mapper4[V1, V2, V3, V4, V5 any] Mapper3[V1, V2, V3, V4]

// Map transforms the elements within the iterator using the provided mapper function.
func (s Mapper4[V1, V2, V3, V4, V5]) Map(mapper mapper[V1, V2]) Mapper3[V2, V3, V4, V5] {
	return func(yield yielder[V2]) {
		for v := range s {
			if !yield(mapper(v)) {
				break
			}
		}
	}
}

func (s Mapper4[V1, V2, V3, V4, V5]) Expand(toElements mapper[V1, Seq[V2]]) Mapper3[V2, V3, V4, V5] {
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

func IterSliceMap4[V1, V2, V3, V4, V5 any](slice []V1) Mapper4[V1, V2, V3, V4, V5] {
	return Mapper4[V1, V2, V3, V4, V5](IterSlice(slice))
}

// See [Seq.Filter].
func (s Mapper4[V1, V2, V3, V4, V5]) Filter(filter yielder[V1]) Mapper4[V1, V2, V3, V4, V5] {
	return Mapper4[V1, V2, V3, V4, V5](Seq[V1](s).Filter(filter))
}

// See [Seq.Skip].
func (s Mapper4[V1, V2, V3, V4, V5]) Skip(toSkip int) Mapper4[V1, V2, V3, V4, V5] {
	return Mapper4[V1, V2, V3, V4, V5](Seq[V1](s).Skip(toSkip))
}

// See [Seq.SkipWhile].
func (s Mapper4[V1, V2, V3, V4, V5]) SkipWhile(test yielder[V1]) Mapper4[V1, V2, V3, V4, V5] {
	return Mapper4[V1, V2, V3, V4, V5](Seq[V1](s).SkipWhile(test))
}

// See [Seq.Take].
func (s Mapper4[V1, V2, V3, V4, V5]) Take(toTake int) Mapper4[V1, V2, V3, V4, V5] {
	return Mapper4[V1, V2, V3, V4, V5](Seq[V1](s).Take(toTake))
}

// See [Seq.TakeWhile].
func (s Mapper4[V1, V2, V3, V4, V5]) TakeWhile(test yielder[V1]) Mapper4[V1, V2, V3, V4, V5] {
	return Mapper4[V1, V2, V3, V4, V5](Seq[V1](s).TakeWhile(test))
}

type KVMapper4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5 any] KVMapper3[K1, V1, K2, V2, K3, V3, K4, V4]

// Map transforms the keys and values within the iterator using the provided mapper function.
func (s KVMapper4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5]) Map(mapper mapper2[K1, V1, K2, V2]) KVMapper3[K2, V2, K3, V3, K4, V4, K5, V5] {
	return func(yield yielder2[K2, V2]) {
		for k, v := range s {
			if !yield(mapper(k, v)) {
				break
			}
		}
	}
}

func IterMapMap4[K1 comparable, V1, K2, V2, K3, V3, K4, V4, K5, V5 any](m map[K1]V1) KVMapper4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5] {
	return KVMapper4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5](IterMap(m))
}

// See [Seq2.Filter].
func (s KVMapper4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5]) Filter(filter yielder2[K1, V1]) KVMapper4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5] {
	return KVMapper4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5](Seq2[K1, V1](s).Filter(filter))
}

// See [Seq2.Skip].
func (s KVMapper4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5]) Skip(toSkip int) KVMapper4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5] {
	return KVMapper4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5](Seq2[K1, V1](s).Skip(toSkip))
}

// See [Seq2.SkipWhile].
func (s KVMapper4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5]) SkipWhile(test yielder2[K1, V1]) KVMapper4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5] {
	return KVMapper4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5](Seq2[K1, V1](s).SkipWhile(test))
}

// See [Seq2.Take].
func (s KVMapper4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5]) Take(toTake int) KVMapper4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5] {
	return KVMapper4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5](Seq2[K1, V1](s).Take(toTake))
}

// See [Seq2.TakeWhile].
func (s KVMapper4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5]) TakeWhile(test yielder2[K1, V1]) KVMapper4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5] {
	return KVMapper4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5](Seq2[K1, V1](s).TakeWhile(test))
}

type Mapper5[V1, V2, V3, V4, V5, V6 any] Mapper4[V1, V2, V3, V4, V5]

// Map transforms the elements within the iterator using the provided mapper function.
func (s Mapper5[V1, V2, V3, V4, V5, V6]) Map(mapper mapper[V1, V2]) Mapper4[V2, V3, V4, V5, V6] {
	return func(yield yielder[V2]) {
		for v := range s {
			if !yield(mapper(v)) {
				break
			}
		}
	}
}

func (s Mapper5[V1, V2, V3, V4, V5, V6]) Expand(toElements mapper[V1, Seq[V2]]) Mapper4[V2, V3, V4, V5, V6] {
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

func IterSliceMap5[V1, V2, V3, V4, V5, V6 any](slice []V1) Mapper5[V1, V2, V3, V4, V5, V6] {
	return Mapper5[V1, V2, V3, V4, V5, V6](IterSlice(slice))
}

// See [Seq.Filter].
func (s Mapper5[V1, V2, V3, V4, V5, V6]) Filter(filter yielder[V1]) Mapper5[V1, V2, V3, V4, V5, V6] {
	return Mapper5[V1, V2, V3, V4, V5, V6](Seq[V1](s).Filter(filter))
}

// See [Seq.Skip].
func (s Mapper5[V1, V2, V3, V4, V5, V6]) Skip(toSkip int) Mapper5[V1, V2, V3, V4, V5, V6] {
	return Mapper5[V1, V2, V3, V4, V5, V6](Seq[V1](s).Skip(toSkip))
}

// See [Seq.SkipWhile].
func (s Mapper5[V1, V2, V3, V4, V5, V6]) SkipWhile(test yielder[V1]) Mapper5[V1, V2, V3, V4, V5, V6] {
	return Mapper5[V1, V2, V3, V4, V5, V6](Seq[V1](s).SkipWhile(test))
}

// See [Seq.Take].
func (s Mapper5[V1, V2, V3, V4, V5, V6]) Take(toTake int) Mapper5[V1, V2, V3, V4, V5, V6] {
	return Mapper5[V1, V2, V3, V4, V5, V6](Seq[V1](s).Take(toTake))
}

// See [Seq.TakeWhile].
func (s Mapper5[V1, V2, V3, V4, V5, V6]) TakeWhile(test yielder[V1]) Mapper5[V1, V2, V3, V4, V5, V6] {
	return Mapper5[V1, V2, V3, V4, V5, V6](Seq[V1](s).TakeWhile(test))
}

type KVMapper5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6 any] KVMapper4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5]

// Map transforms the keys and values within the iterator using the provided mapper function.
func (s KVMapper5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6]) Map(mapper mapper2[K1, V1, K2, V2]) KVMapper4[K2, V2, K3, V3, K4, V4, K5, V5, K6, V6] {
	return func(yield yielder2[K2, V2]) {
		for k, v := range s {
			if !yield(mapper(k, v)) {
				break
			}
		}
	}
}

func IterMapMap5[K1 comparable, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6 any](m map[K1]V1) KVMapper5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6] {
	return KVMapper5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6](IterMap(m))
}

// See [Seq2.Filter].
func (s KVMapper5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6]) Filter(filter yielder2[K1, V1]) KVMapper5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6] {
	return KVMapper5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6](Seq2[K1, V1](s).Filter(filter))
}

// See [Seq2.Skip].
func (s KVMapper5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6]) Skip(toSkip int) KVMapper5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6] {
	return KVMapper5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6](Seq2[K1, V1](s).Skip(toSkip))
}

// See [Seq2.SkipWhile].
func (s KVMapper5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6]) SkipWhile(test yielder2[K1, V1]) KVMapper5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6] {
	return KVMapper5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6](Seq2[K1, V1](s).SkipWhile(test))
}

// See [Seq2.Take].
func (s KVMapper5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6]) Take(toTake int) KVMapper5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6] {
	return KVMapper5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6](Seq2[K1, V1](s).Take(toTake))
}

// See [Seq2.TakeWhile].
func (s KVMapper5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6]) TakeWhile(test yielder2[K1, V1]) KVMapper5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6] {
	return KVMapper5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6](Seq2[K1, V1](s).TakeWhile(test))
}

type Mapper6[V1, V2, V3, V4, V5, V6, V7 any] Mapper5[V1, V2, V3, V4, V5, V6]

// Map transforms the elements within the iterator using the provided mapper function.
func (s Mapper6[V1, V2, V3, V4, V5, V6, V7]) Map(mapper mapper[V1, V2]) Mapper5[V2, V3, V4, V5, V6, V7] {
	return func(yield yielder[V2]) {
		for v := range s {
			if !yield(mapper(v)) {
				break
			}
		}
	}
}

func (s Mapper6[V1, V2, V3, V4, V5, V6, V7]) Expand(toElements mapper[V1, Seq[V2]]) Mapper5[V2, V3, V4, V5, V6, V7] {
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

func IterSliceMap6[V1, V2, V3, V4, V5, V6, V7 any](slice []V1) Mapper6[V1, V2, V3, V4, V5, V6, V7] {
	return Mapper6[V1, V2, V3, V4, V5, V6, V7](IterSlice(slice))
}

// See [Seq.Filter].
func (s Mapper6[V1, V2, V3, V4, V5, V6, V7]) Filter(filter yielder[V1]) Mapper6[V1, V2, V3, V4, V5, V6, V7] {
	return Mapper6[V1, V2, V3, V4, V5, V6, V7](Seq[V1](s).Filter(filter))
}

// See [Seq.Skip].
func (s Mapper6[V1, V2, V3, V4, V5, V6, V7]) Skip(toSkip int) Mapper6[V1, V2, V3, V4, V5, V6, V7] {
	return Mapper6[V1, V2, V3, V4, V5, V6, V7](Seq[V1](s).Skip(toSkip))
}

// See [Seq.SkipWhile].
func (s Mapper6[V1, V2, V3, V4, V5, V6, V7]) SkipWhile(test yielder[V1]) Mapper6[V1, V2, V3, V4, V5, V6, V7] {
	return Mapper6[V1, V2, V3, V4, V5, V6, V7](Seq[V1](s).SkipWhile(test))
}

// See [Seq.Take].
func (s Mapper6[V1, V2, V3, V4, V5, V6, V7]) Take(toTake int) Mapper6[V1, V2, V3, V4, V5, V6, V7] {
	return Mapper6[V1, V2, V3, V4, V5, V6, V7](Seq[V1](s).Take(toTake))
}

// See [Seq.TakeWhile].
func (s Mapper6[V1, V2, V3, V4, V5, V6, V7]) TakeWhile(test yielder[V1]) Mapper6[V1, V2, V3, V4, V5, V6, V7] {
	return Mapper6[V1, V2, V3, V4, V5, V6, V7](Seq[V1](s).TakeWhile(test))
}

type KVMapper6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7 any] KVMapper5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6]

// Map transforms the keys and values within the iterator using the provided mapper function.
func (s KVMapper6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7]) Map(mapper mapper2[K1, V1, K2, V2]) KVMapper5[K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7] {
	return func(yield yielder2[K2, V2]) {
		for k, v := range s {
			if !yield(mapper(k, v)) {
				break
			}
		}
	}
}

func IterMapMap6[K1 comparable, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7 any](m map[K1]V1) KVMapper6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7] {
	return KVMapper6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7](IterMap(m))
}

// See [Seq2.Filter].
func (s KVMapper6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7]) Filter(filter yielder2[K1, V1]) KVMapper6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7] {
	return KVMapper6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7](Seq2[K1, V1](s).Filter(filter))
}

// See [Seq2.Skip].
func (s KVMapper6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7]) Skip(toSkip int) KVMapper6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7] {
	return KVMapper6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7](Seq2[K1, V1](s).Skip(toSkip))
}

// See [Seq2.SkipWhile].
func (s KVMapper6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7]) SkipWhile(test yielder2[K1, V1]) KVMapper6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7] {
	return KVMapper6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7](Seq2[K1, V1](s).SkipWhile(test))
}

// See [Seq2.Take].
func (s KVMapper6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7]) Take(toTake int) KVMapper6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7] {
	return KVMapper6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7](Seq2[K1, V1](s).Take(toTake))
}

// See [Seq2.TakeWhile].
func (s KVMapper6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7]) TakeWhile(test yielder2[K1, V1]) KVMapper6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7] {
	return KVMapper6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7](Seq2[K1, V1](s).TakeWhile(test))
}

type Mapper7[V1, V2, V3, V4, V5, V6, V7, V8 any] Mapper6[V1, V2, V3, V4, V5, V6, V7]

// Map transforms the elements within the iterator using the provided mapper function.
func (s Mapper7[V1, V2, V3, V4, V5, V6, V7, V8]) Map(mapper mapper[V1, V2]) Mapper6[V2, V3, V4, V5, V6, V7, V8] {
	return func(yield yielder[V2]) {
		for v := range s {
			if !yield(mapper(v)) {
				break
			}
		}
	}
}

func (s Mapper7[V1, V2, V3, V4, V5, V6, V7, V8]) Expand(toElements mapper[V1, Seq[V2]]) Mapper6[V2, V3, V4, V5, V6, V7, V8] {
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

func IterSliceMap7[V1, V2, V3, V4, V5, V6, V7, V8 any](slice []V1) Mapper7[V1, V2, V3, V4, V5, V6, V7, V8] {
	return Mapper7[V1, V2, V3, V4, V5, V6, V7, V8](IterSlice(slice))
}

// See [Seq.Filter].
func (s Mapper7[V1, V2, V3, V4, V5, V6, V7, V8]) Filter(filter yielder[V1]) Mapper7[V1, V2, V3, V4, V5, V6, V7, V8] {
	return Mapper7[V1, V2, V3, V4, V5, V6, V7, V8](Seq[V1](s).Filter(filter))
}

// See [Seq.Skip].
func (s Mapper7[V1, V2, V3, V4, V5, V6, V7, V8]) Skip(toSkip int) Mapper7[V1, V2, V3, V4, V5, V6, V7, V8] {
	return Mapper7[V1, V2, V3, V4, V5, V6, V7, V8](Seq[V1](s).Skip(toSkip))
}

// See [Seq.SkipWhile].
func (s Mapper7[V1, V2, V3, V4, V5, V6, V7, V8]) SkipWhile(test yielder[V1]) Mapper7[V1, V2, V3, V4, V5, V6, V7, V8] {
	return Mapper7[V1, V2, V3, V4, V5, V6, V7, V8](Seq[V1](s).SkipWhile(test))
}

// See [Seq.Take].
func (s Mapper7[V1, V2, V3, V4, V5, V6, V7, V8]) Take(toTake int) Mapper7[V1, V2, V3, V4, V5, V6, V7, V8] {
	return Mapper7[V1, V2, V3, V4, V5, V6, V7, V8](Seq[V1](s).Take(toTake))
}

// See [Seq.TakeWhile].
func (s Mapper7[V1, V2, V3, V4, V5, V6, V7, V8]) TakeWhile(test yielder[V1]) Mapper7[V1, V2, V3, V4, V5, V6, V7, V8] {
	return Mapper7[V1, V2, V3, V4, V5, V6, V7, V8](Seq[V1](s).TakeWhile(test))
}

type KVMapper7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8 any] KVMapper6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7]

// Map transforms the keys and values within the iterator using the provided mapper function.
func (s KVMapper7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8]) Map(mapper mapper2[K1, V1, K2, V2]) KVMapper6[K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8] {
	return func(yield yielder2[K2, V2]) {
		for k, v := range s {
			if !yield(mapper(k, v)) {
				break
			}
		}
	}
}

func IterMapMap7[K1 comparable, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8 any](m map[K1]V1) KVMapper7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8] {
	return KVMapper7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8](IterMap(m))
}

// See [Seq2.Filter].
func (s KVMapper7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8]) Filter(filter yielder2[K1, V1]) KVMapper7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8] {
	return KVMapper7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8](Seq2[K1, V1](s).Filter(filter))
}

// See [Seq2.Skip].
func (s KVMapper7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8]) Skip(toSkip int) KVMapper7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8] {
	return KVMapper7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8](Seq2[K1, V1](s).Skip(toSkip))
}

// See [Seq2.SkipWhile].
func (s KVMapper7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8]) SkipWhile(test yielder2[K1, V1]) KVMapper7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8] {
	return KVMapper7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8](Seq2[K1, V1](s).SkipWhile(test))
}

// See [Seq2.Take].
func (s KVMapper7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8]) Take(toTake int) KVMapper7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8] {
	return KVMapper7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8](Seq2[K1, V1](s).Take(toTake))
}

// See [Seq2.TakeWhile].
func (s KVMapper7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8]) TakeWhile(test yielder2[K1, V1]) KVMapper7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8] {
	return KVMapper7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8](Seq2[K1, V1](s).TakeWhile(test))
}

type Mapper8[V1, V2, V3, V4, V5, V6, V7, V8, V9 any] Mapper7[V1, V2, V3, V4, V5, V6, V7, V8]

// Map transforms the elements within the iterator using the provided mapper function.
func (s Mapper8[V1, V2, V3, V4, V5, V6, V7, V8, V9]) Map(mapper mapper[V1, V2]) Mapper7[V2, V3, V4, V5, V6, V7, V8, V9] {
	return func(yield yielder[V2]) {
		for v := range s {
			if !yield(mapper(v)) {
				break
			}
		}
	}
}

func (s Mapper8[V1, V2, V3, V4, V5, V6, V7, V8, V9]) Expand(toElements mapper[V1, Seq[V2]]) Mapper7[V2, V3, V4, V5, V6, V7, V8, V9] {
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

func IterSliceMap8[V1, V2, V3, V4, V5, V6, V7, V8, V9 any](slice []V1) Mapper8[V1, V2, V3, V4, V5, V6, V7, V8, V9] {
	return Mapper8[V1, V2, V3, V4, V5, V6, V7, V8, V9](IterSlice(slice))
}

// See [Seq.Filter].
func (s Mapper8[V1, V2, V3, V4, V5, V6, V7, V8, V9]) Filter(filter yielder[V1]) Mapper8[V1, V2, V3, V4, V5, V6, V7, V8, V9] {
	return Mapper8[V1, V2, V3, V4, V5, V6, V7, V8, V9](Seq[V1](s).Filter(filter))
}

// See [Seq.Skip].
func (s Mapper8[V1, V2, V3, V4, V5, V6, V7, V8, V9]) Skip(toSkip int) Mapper8[V1, V2, V3, V4, V5, V6, V7, V8, V9] {
	return Mapper8[V1, V2, V3, V4, V5, V6, V7, V8, V9](Seq[V1](s).Skip(toSkip))
}

// See [Seq.SkipWhile].
func (s Mapper8[V1, V2, V3, V4, V5, V6, V7, V8, V9]) SkipWhile(test yielder[V1]) Mapper8[V1, V2, V3, V4, V5, V6, V7, V8, V9] {
	return Mapper8[V1, V2, V3, V4, V5, V6, V7, V8, V9](Seq[V1](s).SkipWhile(test))
}

// See [Seq.Take].
func (s Mapper8[V1, V2, V3, V4, V5, V6, V7, V8, V9]) Take(toTake int) Mapper8[V1, V2, V3, V4, V5, V6, V7, V8, V9] {
	return Mapper8[V1, V2, V3, V4, V5, V6, V7, V8, V9](Seq[V1](s).Take(toTake))
}

// See [Seq.TakeWhile].
func (s Mapper8[V1, V2, V3, V4, V5, V6, V7, V8, V9]) TakeWhile(test yielder[V1]) Mapper8[V1, V2, V3, V4, V5, V6, V7, V8, V9] {
	return Mapper8[V1, V2, V3, V4, V5, V6, V7, V8, V9](Seq[V1](s).TakeWhile(test))
}

type KVMapper8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9 any] KVMapper7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8]

// Map transforms the keys and values within the iterator using the provided mapper function.
func (s KVMapper8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9]) Map(mapper mapper2[K1, V1, K2, V2]) KVMapper7[K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9] {
	return func(yield yielder2[K2, V2]) {
		for k, v := range s {
			if !yield(mapper(k, v)) {
				break
			}
		}
	}
}

func IterMapMap8[K1 comparable, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9 any](m map[K1]V1) KVMapper8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9] {
	return KVMapper8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9](IterMap(m))
}

// See [Seq2.Filter].
func (s KVMapper8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9]) Filter(filter yielder2[K1, V1]) KVMapper8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9] {
	return KVMapper8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9](Seq2[K1, V1](s).Filter(filter))
}

// See [Seq2.Skip].
func (s KVMapper8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9]) Skip(toSkip int) KVMapper8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9] {
	return KVMapper8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9](Seq2[K1, V1](s).Skip(toSkip))
}

// See [Seq2.SkipWhile].
func (s KVMapper8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9]) SkipWhile(test yielder2[K1, V1]) KVMapper8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9] {
	return KVMapper8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9](Seq2[K1, V1](s).SkipWhile(test))
}

// See [Seq2.Take].
func (s KVMapper8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9]) Take(toTake int) KVMapper8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9] {
	return KVMapper8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9](Seq2[K1, V1](s).Take(toTake))
}

// See [Seq2.TakeWhile].
func (s KVMapper8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9]) TakeWhile(test yielder2[K1, V1]) KVMapper8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9] {
	return KVMapper8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9](Seq2[K1, V1](s).TakeWhile(test))
}

type Mapper9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10 any] Mapper8[V1, V2, V3, V4, V5, V6, V7, V8, V9]

// Map transforms the elements within the iterator using the provided mapper function.
func (s Mapper9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10]) Map(mapper mapper[V1, V2]) Mapper8[V2, V3, V4, V5, V6, V7, V8, V9, V10] {
	return func(yield yielder[V2]) {
		for v := range s {
			if !yield(mapper(v)) {
				break
			}
		}
	}
}

func (s Mapper9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10]) Expand(toElements mapper[V1, Seq[V2]]) Mapper8[V2, V3, V4, V5, V6, V7, V8, V9, V10] {
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

func IterSliceMap9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10 any](slice []V1) Mapper9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10] {
	return Mapper9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10](IterSlice(slice))
}

// See [Seq.Filter].
func (s Mapper9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10]) Filter(filter yielder[V1]) Mapper9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10] {
	return Mapper9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10](Seq[V1](s).Filter(filter))
}

// See [Seq.Skip].
func (s Mapper9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10]) Skip(toSkip int) Mapper9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10] {
	return Mapper9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10](Seq[V1](s).Skip(toSkip))
}

// See [Seq.SkipWhile].
func (s Mapper9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10]) SkipWhile(test yielder[V1]) Mapper9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10] {
	return Mapper9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10](Seq[V1](s).SkipWhile(test))
}

// See [Seq.Take].
func (s Mapper9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10]) Take(toTake int) Mapper9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10] {
	return Mapper9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10](Seq[V1](s).Take(toTake))
}

// See [Seq.TakeWhile].
func (s Mapper9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10]) TakeWhile(test yielder[V1]) Mapper9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10] {
	return Mapper9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10](Seq[V1](s).TakeWhile(test))
}

type KVMapper9[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10 any] KVMapper8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9]

// Map transforms the keys and values within the iterator using the provided mapper function.
func (s KVMapper9[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10]) Map(mapper mapper2[K1, V1, K2, V2]) KVMapper8[K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10] {
	return func(yield yielder2[K2, V2]) {
		for k, v := range s {
			if !yield(mapper(k, v)) {
				break
			}
		}
	}
}

func IterMapMap9[K1 comparable, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10 any](m map[K1]V1) KVMapper9[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10] {
	return KVMapper9[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10](IterMap(m))
}

// See [Seq2.Filter].
func (s KVMapper9[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10]) Filter(filter yielder2[K1, V1]) KVMapper9[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10] {
	return KVMapper9[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10](Seq2[K1, V1](s).Filter(filter))
}

// See [Seq2.Skip].
func (s KVMapper9[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10]) Skip(toSkip int) KVMapper9[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10] {
	return KVMapper9[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10](Seq2[K1, V1](s).Skip(toSkip))
}

// See [Seq2.SkipWhile].
func (s KVMapper9[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10]) SkipWhile(test yielder2[K1, V1]) KVMapper9[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10] {
	return KVMapper9[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10](Seq2[K1, V1](s).SkipWhile(test))
}

// See [Seq2.Take].
func (s KVMapper9[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10]) Take(toTake int) KVMapper9[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10] {
	return KVMapper9[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10](Seq2[K1, V1](s).Take(toTake))
}

// See [Seq2.TakeWhile].
func (s KVMapper9[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10]) TakeWhile(test yielder2[K1, V1]) KVMapper9[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10] {
	return KVMapper9[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10](Seq2[K1, V1](s).TakeWhile(test))
}
