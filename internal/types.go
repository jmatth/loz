//nolint:godoclint
package internal

type (
	Yielder[V any]               = func(V) bool
	Processor[V any]             = func(V)
	Mapper[V, O any]             = func(V) O
	FilteringMapper[V, O any]    = func(V) (O, bool)
	FilteringMapperErr[V, O any] = func(V) (O, error)
	Reducer[V, O any]            = func(O, V) O
)

// Switch back to this when the go team fixes their compiler.
// https://github.com/golang/go/issues/63285
// type expander[V, O any] = func(V) Seq[O]

type (
	Yielder2[K, V any]                      = func(K, V) bool
	Reducer2[K, V, OK, OV any]              = func(OK, OV, K, V) (OK, OV)
	Mapper2[K1, V1, K2, V2 any]             = func(K1, V1) (K2, V2)
	FilteringMapper2[K1, V1, K2, V2 any]    = func(K1, V1) (K2, V2, bool)
	FilteringMapperErr2[K1, V1, K2, V2 any] = func(K1, V1) (K2, V2, error)
)
