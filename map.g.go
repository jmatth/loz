package loz

type Map1[V1, V2 any] Seq[V1]

// Map transforms the elements within the iterator using the provided mapper function.
func (s Map1[V1, V2]) Map(mapper mapper[V1, V2]) Seq[V2] {
	return func(yield yielder[V2]) {
		for v := range s {
			if !yield(mapper(v)) {
				break
			}
		}
	}
}

func (s Map1[V1, V2]) FilterMap(mapper filteringMapper[V1, V2]) Seq[V2] {
	return func(yield yielder[V2]) {
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

func (s Map1[V1, V2]) Expand(toElements mapper[V1, Seq[V2]]) Seq[V2] {
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

type KVMap1[K1, V1, K2, V2 any] KVSeq[K1, V1]

// Map transforms the keys and values within the iterator using the provided mapper function.
func (s KVMap1[K1, V1, K2, V2]) Map(mapper mapper2[K1, V1, K2, V2]) KVSeq[K2, V2] {
	return func(yield yielder2[K2, V2]) {
		for k, v := range s {
			if !yield(mapper(k, v)) {
				break
			}
		}
	}
}

// Map transforms the keys and values within the iterator using the provided mapper function.
func (s KVMap1[K1, V1, K2, V2]) FilterMap(mapper filteringMapper2[K1, V1, K2, V2]) KVSeq[K2, V2] {
	return func(yield yielder2[K2, V2]) {
		for k, v := range s {
			mk, mv, err := mapper(k, v)
			if err != nil {
				continue
			}
			if !yield(mk, mv) {
				break
			}
		}
	}
}

// See [KVSeq.Filter].
func (s KVMap1[K1, V1, K2, V2]) Filter(filter yielder2[K1, V1]) KVMap1[K1, V1, K2, V2] {
	return KVMap1[K1, V1, K2, V2](KVSeq[K1, V1](s).Filter(filter))
}

// See [KVSeq.Skip].
func (s KVMap1[K1, V1, K2, V2]) Skip(toSkip int) KVMap1[K1, V1, K2, V2] {
	return KVMap1[K1, V1, K2, V2](KVSeq[K1, V1](s).Skip(toSkip))
}

// See [KVSeq.SkipWhile].
func (s KVMap1[K1, V1, K2, V2]) SkipWhile(test yielder2[K1, V1]) KVMap1[K1, V1, K2, V2] {
	return KVMap1[K1, V1, K2, V2](KVSeq[K1, V1](s).SkipWhile(test))
}

// See [KVSeq.Take].
func (s KVMap1[K1, V1, K2, V2]) Take(toTake int) KVMap1[K1, V1, K2, V2] {
	return KVMap1[K1, V1, K2, V2](KVSeq[K1, V1](s).Take(toTake))
}

// See [KVSeq.TakeWhile].
func (s KVMap1[K1, V1, K2, V2]) TakeWhile(test yielder2[K1, V1]) KVMap1[K1, V1, K2, V2] {
	return KVMap1[K1, V1, K2, V2](KVSeq[K1, V1](s).TakeWhile(test))
}

type Map2[V1, V2, V3 any] Map1[V1, V2]

// Map transforms the elements within the iterator using the provided mapper function.
func (s Map2[V1, V2, V3]) Map(mapper mapper[V1, V2]) Map1[V2, V3] {
	return func(yield yielder[V2]) {
		for v := range s {
			if !yield(mapper(v)) {
				break
			}
		}
	}
}

func (s Map2[V1, V2, V3]) FilterMap(mapper filteringMapper[V1, V2]) Map1[V2, V3] {
	return func(yield yielder[V2]) {
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

func (s Map2[V1, V2, V3]) Expand(toElements mapper[V1, Seq[V2]]) Map1[V2, V3] {
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

type KVMap2[K1, V1, K2, V2, K3, V3 any] KVMap1[K1, V1, K2, V2]

// Map transforms the keys and values within the iterator using the provided mapper function.
func (s KVMap2[K1, V1, K2, V2, K3, V3]) Map(mapper mapper2[K1, V1, K2, V2]) KVMap1[K2, V2, K3, V3] {
	return func(yield yielder2[K2, V2]) {
		for k, v := range s {
			if !yield(mapper(k, v)) {
				break
			}
		}
	}
}

// Map transforms the keys and values within the iterator using the provided mapper function.
func (s KVMap2[K1, V1, K2, V2, K3, V3]) FilterMap(mapper filteringMapper2[K1, V1, K2, V2]) KVMap1[K2, V2, K3, V3] {
	return func(yield yielder2[K2, V2]) {
		for k, v := range s {
			mk, mv, err := mapper(k, v)
			if err != nil {
				continue
			}
			if !yield(mk, mv) {
				break
			}
		}
	}
}

// See [KVSeq.Filter].
func (s KVMap2[K1, V1, K2, V2, K3, V3]) Filter(filter yielder2[K1, V1]) KVMap2[K1, V1, K2, V2, K3, V3] {
	return KVMap2[K1, V1, K2, V2, K3, V3](KVSeq[K1, V1](s).Filter(filter))
}

// See [KVSeq.Skip].
func (s KVMap2[K1, V1, K2, V2, K3, V3]) Skip(toSkip int) KVMap2[K1, V1, K2, V2, K3, V3] {
	return KVMap2[K1, V1, K2, V2, K3, V3](KVSeq[K1, V1](s).Skip(toSkip))
}

// See [KVSeq.SkipWhile].
func (s KVMap2[K1, V1, K2, V2, K3, V3]) SkipWhile(test yielder2[K1, V1]) KVMap2[K1, V1, K2, V2, K3, V3] {
	return KVMap2[K1, V1, K2, V2, K3, V3](KVSeq[K1, V1](s).SkipWhile(test))
}

// See [KVSeq.Take].
func (s KVMap2[K1, V1, K2, V2, K3, V3]) Take(toTake int) KVMap2[K1, V1, K2, V2, K3, V3] {
	return KVMap2[K1, V1, K2, V2, K3, V3](KVSeq[K1, V1](s).Take(toTake))
}

// See [KVSeq.TakeWhile].
func (s KVMap2[K1, V1, K2, V2, K3, V3]) TakeWhile(test yielder2[K1, V1]) KVMap2[K1, V1, K2, V2, K3, V3] {
	return KVMap2[K1, V1, K2, V2, K3, V3](KVSeq[K1, V1](s).TakeWhile(test))
}

type Map3[V1, V2, V3, V4 any] Map2[V1, V2, V3]

// Map transforms the elements within the iterator using the provided mapper function.
func (s Map3[V1, V2, V3, V4]) Map(mapper mapper[V1, V2]) Map2[V2, V3, V4] {
	return func(yield yielder[V2]) {
		for v := range s {
			if !yield(mapper(v)) {
				break
			}
		}
	}
}

func (s Map3[V1, V2, V3, V4]) FilterMap(mapper filteringMapper[V1, V2]) Map2[V2, V3, V4] {
	return func(yield yielder[V2]) {
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

func (s Map3[V1, V2, V3, V4]) Expand(toElements mapper[V1, Seq[V2]]) Map2[V2, V3, V4] {
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

type KVMap3[K1, V1, K2, V2, K3, V3, K4, V4 any] KVMap2[K1, V1, K2, V2, K3, V3]

// Map transforms the keys and values within the iterator using the provided mapper function.
func (s KVMap3[K1, V1, K2, V2, K3, V3, K4, V4]) Map(mapper mapper2[K1, V1, K2, V2]) KVMap2[K2, V2, K3, V3, K4, V4] {
	return func(yield yielder2[K2, V2]) {
		for k, v := range s {
			if !yield(mapper(k, v)) {
				break
			}
		}
	}
}

// Map transforms the keys and values within the iterator using the provided mapper function.
func (s KVMap3[K1, V1, K2, V2, K3, V3, K4, V4]) FilterMap(mapper filteringMapper2[K1, V1, K2, V2]) KVMap2[K2, V2, K3, V3, K4, V4] {
	return func(yield yielder2[K2, V2]) {
		for k, v := range s {
			mk, mv, err := mapper(k, v)
			if err != nil {
				continue
			}
			if !yield(mk, mv) {
				break
			}
		}
	}
}

// See [KVSeq.Filter].
func (s KVMap3[K1, V1, K2, V2, K3, V3, K4, V4]) Filter(filter yielder2[K1, V1]) KVMap3[K1, V1, K2, V2, K3, V3, K4, V4] {
	return KVMap3[K1, V1, K2, V2, K3, V3, K4, V4](KVSeq[K1, V1](s).Filter(filter))
}

// See [KVSeq.Skip].
func (s KVMap3[K1, V1, K2, V2, K3, V3, K4, V4]) Skip(toSkip int) KVMap3[K1, V1, K2, V2, K3, V3, K4, V4] {
	return KVMap3[K1, V1, K2, V2, K3, V3, K4, V4](KVSeq[K1, V1](s).Skip(toSkip))
}

// See [KVSeq.SkipWhile].
func (s KVMap3[K1, V1, K2, V2, K3, V3, K4, V4]) SkipWhile(test yielder2[K1, V1]) KVMap3[K1, V1, K2, V2, K3, V3, K4, V4] {
	return KVMap3[K1, V1, K2, V2, K3, V3, K4, V4](KVSeq[K1, V1](s).SkipWhile(test))
}

// See [KVSeq.Take].
func (s KVMap3[K1, V1, K2, V2, K3, V3, K4, V4]) Take(toTake int) KVMap3[K1, V1, K2, V2, K3, V3, K4, V4] {
	return KVMap3[K1, V1, K2, V2, K3, V3, K4, V4](KVSeq[K1, V1](s).Take(toTake))
}

// See [KVSeq.TakeWhile].
func (s KVMap3[K1, V1, K2, V2, K3, V3, K4, V4]) TakeWhile(test yielder2[K1, V1]) KVMap3[K1, V1, K2, V2, K3, V3, K4, V4] {
	return KVMap3[K1, V1, K2, V2, K3, V3, K4, V4](KVSeq[K1, V1](s).TakeWhile(test))
}

type Map4[V1, V2, V3, V4, V5 any] Map3[V1, V2, V3, V4]

// Map transforms the elements within the iterator using the provided mapper function.
func (s Map4[V1, V2, V3, V4, V5]) Map(mapper mapper[V1, V2]) Map3[V2, V3, V4, V5] {
	return func(yield yielder[V2]) {
		for v := range s {
			if !yield(mapper(v)) {
				break
			}
		}
	}
}

func (s Map4[V1, V2, V3, V4, V5]) FilterMap(mapper filteringMapper[V1, V2]) Map3[V2, V3, V4, V5] {
	return func(yield yielder[V2]) {
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

func (s Map4[V1, V2, V3, V4, V5]) Expand(toElements mapper[V1, Seq[V2]]) Map3[V2, V3, V4, V5] {
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

type KVMap4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5 any] KVMap3[K1, V1, K2, V2, K3, V3, K4, V4]

// Map transforms the keys and values within the iterator using the provided mapper function.
func (s KVMap4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5]) Map(mapper mapper2[K1, V1, K2, V2]) KVMap3[K2, V2, K3, V3, K4, V4, K5, V5] {
	return func(yield yielder2[K2, V2]) {
		for k, v := range s {
			if !yield(mapper(k, v)) {
				break
			}
		}
	}
}

// Map transforms the keys and values within the iterator using the provided mapper function.
func (s KVMap4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5]) FilterMap(mapper filteringMapper2[K1, V1, K2, V2]) KVMap3[K2, V2, K3, V3, K4, V4, K5, V5] {
	return func(yield yielder2[K2, V2]) {
		for k, v := range s {
			mk, mv, err := mapper(k, v)
			if err != nil {
				continue
			}
			if !yield(mk, mv) {
				break
			}
		}
	}
}

// See [KVSeq.Filter].
func (s KVMap4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5]) Filter(filter yielder2[K1, V1]) KVMap4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5] {
	return KVMap4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5](KVSeq[K1, V1](s).Filter(filter))
}

// See [KVSeq.Skip].
func (s KVMap4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5]) Skip(toSkip int) KVMap4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5] {
	return KVMap4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5](KVSeq[K1, V1](s).Skip(toSkip))
}

// See [KVSeq.SkipWhile].
func (s KVMap4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5]) SkipWhile(test yielder2[K1, V1]) KVMap4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5] {
	return KVMap4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5](KVSeq[K1, V1](s).SkipWhile(test))
}

// See [KVSeq.Take].
func (s KVMap4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5]) Take(toTake int) KVMap4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5] {
	return KVMap4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5](KVSeq[K1, V1](s).Take(toTake))
}

// See [KVSeq.TakeWhile].
func (s KVMap4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5]) TakeWhile(test yielder2[K1, V1]) KVMap4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5] {
	return KVMap4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5](KVSeq[K1, V1](s).TakeWhile(test))
}

type Map5[V1, V2, V3, V4, V5, V6 any] Map4[V1, V2, V3, V4, V5]

// Map transforms the elements within the iterator using the provided mapper function.
func (s Map5[V1, V2, V3, V4, V5, V6]) Map(mapper mapper[V1, V2]) Map4[V2, V3, V4, V5, V6] {
	return func(yield yielder[V2]) {
		for v := range s {
			if !yield(mapper(v)) {
				break
			}
		}
	}
}

func (s Map5[V1, V2, V3, V4, V5, V6]) FilterMap(mapper filteringMapper[V1, V2]) Map4[V2, V3, V4, V5, V6] {
	return func(yield yielder[V2]) {
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

func (s Map5[V1, V2, V3, V4, V5, V6]) Expand(toElements mapper[V1, Seq[V2]]) Map4[V2, V3, V4, V5, V6] {
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

type KVMap5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6 any] KVMap4[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5]

// Map transforms the keys and values within the iterator using the provided mapper function.
func (s KVMap5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6]) Map(mapper mapper2[K1, V1, K2, V2]) KVMap4[K2, V2, K3, V3, K4, V4, K5, V5, K6, V6] {
	return func(yield yielder2[K2, V2]) {
		for k, v := range s {
			if !yield(mapper(k, v)) {
				break
			}
		}
	}
}

// Map transforms the keys and values within the iterator using the provided mapper function.
func (s KVMap5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6]) FilterMap(mapper filteringMapper2[K1, V1, K2, V2]) KVMap4[K2, V2, K3, V3, K4, V4, K5, V5, K6, V6] {
	return func(yield yielder2[K2, V2]) {
		for k, v := range s {
			mk, mv, err := mapper(k, v)
			if err != nil {
				continue
			}
			if !yield(mk, mv) {
				break
			}
		}
	}
}

// See [KVSeq.Filter].
func (s KVMap5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6]) Filter(filter yielder2[K1, V1]) KVMap5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6] {
	return KVMap5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6](KVSeq[K1, V1](s).Filter(filter))
}

// See [KVSeq.Skip].
func (s KVMap5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6]) Skip(toSkip int) KVMap5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6] {
	return KVMap5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6](KVSeq[K1, V1](s).Skip(toSkip))
}

// See [KVSeq.SkipWhile].
func (s KVMap5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6]) SkipWhile(test yielder2[K1, V1]) KVMap5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6] {
	return KVMap5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6](KVSeq[K1, V1](s).SkipWhile(test))
}

// See [KVSeq.Take].
func (s KVMap5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6]) Take(toTake int) KVMap5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6] {
	return KVMap5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6](KVSeq[K1, V1](s).Take(toTake))
}

// See [KVSeq.TakeWhile].
func (s KVMap5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6]) TakeWhile(test yielder2[K1, V1]) KVMap5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6] {
	return KVMap5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6](KVSeq[K1, V1](s).TakeWhile(test))
}

type Map6[V1, V2, V3, V4, V5, V6, V7 any] Map5[V1, V2, V3, V4, V5, V6]

// Map transforms the elements within the iterator using the provided mapper function.
func (s Map6[V1, V2, V3, V4, V5, V6, V7]) Map(mapper mapper[V1, V2]) Map5[V2, V3, V4, V5, V6, V7] {
	return func(yield yielder[V2]) {
		for v := range s {
			if !yield(mapper(v)) {
				break
			}
		}
	}
}

func (s Map6[V1, V2, V3, V4, V5, V6, V7]) FilterMap(mapper filteringMapper[V1, V2]) Map5[V2, V3, V4, V5, V6, V7] {
	return func(yield yielder[V2]) {
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

func (s Map6[V1, V2, V3, V4, V5, V6, V7]) Expand(toElements mapper[V1, Seq[V2]]) Map5[V2, V3, V4, V5, V6, V7] {
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

type KVMap6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7 any] KVMap5[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6]

// Map transforms the keys and values within the iterator using the provided mapper function.
func (s KVMap6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7]) Map(mapper mapper2[K1, V1, K2, V2]) KVMap5[K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7] {
	return func(yield yielder2[K2, V2]) {
		for k, v := range s {
			if !yield(mapper(k, v)) {
				break
			}
		}
	}
}

// Map transforms the keys and values within the iterator using the provided mapper function.
func (s KVMap6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7]) FilterMap(mapper filteringMapper2[K1, V1, K2, V2]) KVMap5[K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7] {
	return func(yield yielder2[K2, V2]) {
		for k, v := range s {
			mk, mv, err := mapper(k, v)
			if err != nil {
				continue
			}
			if !yield(mk, mv) {
				break
			}
		}
	}
}

// See [KVSeq.Filter].
func (s KVMap6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7]) Filter(filter yielder2[K1, V1]) KVMap6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7] {
	return KVMap6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7](KVSeq[K1, V1](s).Filter(filter))
}

// See [KVSeq.Skip].
func (s KVMap6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7]) Skip(toSkip int) KVMap6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7] {
	return KVMap6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7](KVSeq[K1, V1](s).Skip(toSkip))
}

// See [KVSeq.SkipWhile].
func (s KVMap6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7]) SkipWhile(test yielder2[K1, V1]) KVMap6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7] {
	return KVMap6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7](KVSeq[K1, V1](s).SkipWhile(test))
}

// See [KVSeq.Take].
func (s KVMap6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7]) Take(toTake int) KVMap6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7] {
	return KVMap6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7](KVSeq[K1, V1](s).Take(toTake))
}

// See [KVSeq.TakeWhile].
func (s KVMap6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7]) TakeWhile(test yielder2[K1, V1]) KVMap6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7] {
	return KVMap6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7](KVSeq[K1, V1](s).TakeWhile(test))
}

type Map7[V1, V2, V3, V4, V5, V6, V7, V8 any] Map6[V1, V2, V3, V4, V5, V6, V7]

// Map transforms the elements within the iterator using the provided mapper function.
func (s Map7[V1, V2, V3, V4, V5, V6, V7, V8]) Map(mapper mapper[V1, V2]) Map6[V2, V3, V4, V5, V6, V7, V8] {
	return func(yield yielder[V2]) {
		for v := range s {
			if !yield(mapper(v)) {
				break
			}
		}
	}
}

func (s Map7[V1, V2, V3, V4, V5, V6, V7, V8]) FilterMap(mapper filteringMapper[V1, V2]) Map6[V2, V3, V4, V5, V6, V7, V8] {
	return func(yield yielder[V2]) {
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

func (s Map7[V1, V2, V3, V4, V5, V6, V7, V8]) Expand(toElements mapper[V1, Seq[V2]]) Map6[V2, V3, V4, V5, V6, V7, V8] {
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

type KVMap7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8 any] KVMap6[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7]

// Map transforms the keys and values within the iterator using the provided mapper function.
func (s KVMap7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8]) Map(mapper mapper2[K1, V1, K2, V2]) KVMap6[K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8] {
	return func(yield yielder2[K2, V2]) {
		for k, v := range s {
			if !yield(mapper(k, v)) {
				break
			}
		}
	}
}

// Map transforms the keys and values within the iterator using the provided mapper function.
func (s KVMap7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8]) FilterMap(mapper filteringMapper2[K1, V1, K2, V2]) KVMap6[K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8] {
	return func(yield yielder2[K2, V2]) {
		for k, v := range s {
			mk, mv, err := mapper(k, v)
			if err != nil {
				continue
			}
			if !yield(mk, mv) {
				break
			}
		}
	}
}

// See [KVSeq.Filter].
func (s KVMap7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8]) Filter(filter yielder2[K1, V1]) KVMap7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8] {
	return KVMap7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8](KVSeq[K1, V1](s).Filter(filter))
}

// See [KVSeq.Skip].
func (s KVMap7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8]) Skip(toSkip int) KVMap7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8] {
	return KVMap7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8](KVSeq[K1, V1](s).Skip(toSkip))
}

// See [KVSeq.SkipWhile].
func (s KVMap7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8]) SkipWhile(test yielder2[K1, V1]) KVMap7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8] {
	return KVMap7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8](KVSeq[K1, V1](s).SkipWhile(test))
}

// See [KVSeq.Take].
func (s KVMap7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8]) Take(toTake int) KVMap7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8] {
	return KVMap7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8](KVSeq[K1, V1](s).Take(toTake))
}

// See [KVSeq.TakeWhile].
func (s KVMap7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8]) TakeWhile(test yielder2[K1, V1]) KVMap7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8] {
	return KVMap7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8](KVSeq[K1, V1](s).TakeWhile(test))
}

type Map8[V1, V2, V3, V4, V5, V6, V7, V8, V9 any] Map7[V1, V2, V3, V4, V5, V6, V7, V8]

// Map transforms the elements within the iterator using the provided mapper function.
func (s Map8[V1, V2, V3, V4, V5, V6, V7, V8, V9]) Map(mapper mapper[V1, V2]) Map7[V2, V3, V4, V5, V6, V7, V8, V9] {
	return func(yield yielder[V2]) {
		for v := range s {
			if !yield(mapper(v)) {
				break
			}
		}
	}
}

func (s Map8[V1, V2, V3, V4, V5, V6, V7, V8, V9]) FilterMap(mapper filteringMapper[V1, V2]) Map7[V2, V3, V4, V5, V6, V7, V8, V9] {
	return func(yield yielder[V2]) {
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

func (s Map8[V1, V2, V3, V4, V5, V6, V7, V8, V9]) Expand(toElements mapper[V1, Seq[V2]]) Map7[V2, V3, V4, V5, V6, V7, V8, V9] {
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

type KVMap8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9 any] KVMap7[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8]

// Map transforms the keys and values within the iterator using the provided mapper function.
func (s KVMap8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9]) Map(mapper mapper2[K1, V1, K2, V2]) KVMap7[K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9] {
	return func(yield yielder2[K2, V2]) {
		for k, v := range s {
			if !yield(mapper(k, v)) {
				break
			}
		}
	}
}

// Map transforms the keys and values within the iterator using the provided mapper function.
func (s KVMap8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9]) FilterMap(mapper filteringMapper2[K1, V1, K2, V2]) KVMap7[K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9] {
	return func(yield yielder2[K2, V2]) {
		for k, v := range s {
			mk, mv, err := mapper(k, v)
			if err != nil {
				continue
			}
			if !yield(mk, mv) {
				break
			}
		}
	}
}

// See [KVSeq.Filter].
func (s KVMap8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9]) Filter(filter yielder2[K1, V1]) KVMap8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9] {
	return KVMap8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9](KVSeq[K1, V1](s).Filter(filter))
}

// See [KVSeq.Skip].
func (s KVMap8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9]) Skip(toSkip int) KVMap8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9] {
	return KVMap8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9](KVSeq[K1, V1](s).Skip(toSkip))
}

// See [KVSeq.SkipWhile].
func (s KVMap8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9]) SkipWhile(test yielder2[K1, V1]) KVMap8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9] {
	return KVMap8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9](KVSeq[K1, V1](s).SkipWhile(test))
}

// See [KVSeq.Take].
func (s KVMap8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9]) Take(toTake int) KVMap8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9] {
	return KVMap8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9](KVSeq[K1, V1](s).Take(toTake))
}

// See [KVSeq.TakeWhile].
func (s KVMap8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9]) TakeWhile(test yielder2[K1, V1]) KVMap8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9] {
	return KVMap8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9](KVSeq[K1, V1](s).TakeWhile(test))
}

type Map9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10 any] Map8[V1, V2, V3, V4, V5, V6, V7, V8, V9]

// Map transforms the elements within the iterator using the provided mapper function.
func (s Map9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10]) Map(mapper mapper[V1, V2]) Map8[V2, V3, V4, V5, V6, V7, V8, V9, V10] {
	return func(yield yielder[V2]) {
		for v := range s {
			if !yield(mapper(v)) {
				break
			}
		}
	}
}

func (s Map9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10]) FilterMap(mapper filteringMapper[V1, V2]) Map8[V2, V3, V4, V5, V6, V7, V8, V9, V10] {
	return func(yield yielder[V2]) {
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

func (s Map9[V1, V2, V3, V4, V5, V6, V7, V8, V9, V10]) Expand(toElements mapper[V1, Seq[V2]]) Map8[V2, V3, V4, V5, V6, V7, V8, V9, V10] {
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

type KVMap9[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10 any] KVMap8[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9]

// Map transforms the keys and values within the iterator using the provided mapper function.
func (s KVMap9[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10]) Map(mapper mapper2[K1, V1, K2, V2]) KVMap8[K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10] {
	return func(yield yielder2[K2, V2]) {
		for k, v := range s {
			if !yield(mapper(k, v)) {
				break
			}
		}
	}
}

// Map transforms the keys and values within the iterator using the provided mapper function.
func (s KVMap9[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10]) FilterMap(mapper filteringMapper2[K1, V1, K2, V2]) KVMap8[K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10] {
	return func(yield yielder2[K2, V2]) {
		for k, v := range s {
			mk, mv, err := mapper(k, v)
			if err != nil {
				continue
			}
			if !yield(mk, mv) {
				break
			}
		}
	}
}

// See [KVSeq.Filter].
func (s KVMap9[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10]) Filter(filter yielder2[K1, V1]) KVMap9[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10] {
	return KVMap9[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10](KVSeq[K1, V1](s).Filter(filter))
}

// See [KVSeq.Skip].
func (s KVMap9[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10]) Skip(toSkip int) KVMap9[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10] {
	return KVMap9[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10](KVSeq[K1, V1](s).Skip(toSkip))
}

// See [KVSeq.SkipWhile].
func (s KVMap9[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10]) SkipWhile(test yielder2[K1, V1]) KVMap9[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10] {
	return KVMap9[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10](KVSeq[K1, V1](s).SkipWhile(test))
}

// See [KVSeq.Take].
func (s KVMap9[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10]) Take(toTake int) KVMap9[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10] {
	return KVMap9[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10](KVSeq[K1, V1](s).Take(toTake))
}

// See [KVSeq.TakeWhile].
func (s KVMap9[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10]) TakeWhile(test yielder2[K1, V1]) KVMap9[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10] {
	return KVMap9[K1, V1, K2, V2, K3, V3, K4, V4, K5, V5, K6, V6, K7, V7, K8, V8, K9, V9, K10, V10](KVSeq[K1, V1](s).TakeWhile(test))
}
