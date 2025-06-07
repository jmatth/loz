package loz

type yielder[V any] = func(V) bool
type processor[V any] = func(V)
type mapper[V, O any] = func(V) O
type filteringMapper[V, O any] = func(V) (O, error)
type reducer[V, O any] = func(O, V) O

// Switch back to this when the go team fixes their compiler.
// https://github.com/golang/go/issues/63285
// type expander[V, O any] = func(V) Seq[O]

type yielder2[K, V any] = func(K, V) bool
type reducer2[K, V any] = func(K, V, K, V) (K, V)
type mapper2[K1, V1, K2, V2 any] = func(K1, V1) (K2, V2)
type filteringMapper2[K1, V1, K2, V2 any] = func(K1, V1) (K2, V2, error)
