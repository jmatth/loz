package internal

type Yielder[V any] = func(V) bool
type Processor[V any] = func(V)
type Mapper[V, O any] = func(V) O
type FilteringMapper[V, O any] = func(V) (O, bool)
type FilteringMapperErr[V, O any] = func(V) (O, error)
type Reducer[V, O any] = func(O, V) O

// Switch back to this when the go team fixes their compiler.
// https://github.com/golang/go/issues/63285
// type expander[V, O any] = func(V) Seq[O]

type Yielder2[K, V any] = func(K, V) bool
type Reducer2[K, V any] = func(K, V, K, V) (K, V)
type Mapper2[K1, V1, K2, V2 any] = func(K1, V1) (K2, V2)
type FilteringMapper2[K1, V1, K2, V2 any] = func(K1, V1) (K2, V2, bool)
type FilteringMapperErr2[K1, V1, K2, V2 any] = func(K1, V1) (K2, V2, error)
