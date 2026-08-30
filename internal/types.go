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

type (
	KVYielder[K, V any]                      = func(K, V) bool
	KVReducer[K, V, OK, OV any]              = func(OK, OV, K, V) (OK, OV)
	KVMapper[K1, V1, K2, V2 any]             = func(K1, V1) (K2, V2)
	KVFilteringMapper[K1, V1, K2, V2 any]    = func(K1, V1) (K2, V2, bool)
	KVFilteringMapperErr[K1, V1, K2, V2 any] = func(K1, V1) (K2, V2, error)
)
