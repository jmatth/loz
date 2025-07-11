package loz_test

import (
	"errors"
	"fmt"
	"iter"
	"maps"
	"slices"
	"strings"
	"testing"

	"github.com/jmatth/loz"
	"github.com/stretchr/testify/assert"
)

func toMap[K comparable, V any](seq loz.KVSeq[K, V]) map[K]V {
	return maps.Collect(iter.Seq2[K, V](seq))
}

func ExampleKVSeq_Keys() {
	keys := loz.IterMap(map[int]string{1: "one", 2: "two", 3: "three"}).
		Keys().
		CollectSlice()
	slices.Sort(keys)
	fmt.Printf("%v", keys)
	// Output: [1 2 3]
}

func ExampleKVSeq_Values() {
	vals := loz.IterMap(map[int]string{1: "one", 2: "two", 3: "three"}).
		Values().
		CollectSlice()
	slices.Sort(vals)
	fmt.Printf("%v", vals)
	// Output: [one three two]
}

func iterKVPairs[K, V any](kvs ...any) loz.KVSeq[K, V] {
	return func(yield func(K, V) bool) {
		for i := 0; i < len(kvs); i += 2 {
			if !yield(kvs[i].(K), kvs[i+1].(V)) {
				break
			}
		}
	}
}

// func ExampleCompKKVSeq_CollectMap() {
// 	result := CompKKVSeq[int, string](iterKVPairs[int, string](1, "one", 2, "two", 3, "three")).
// 		CollectMap()
// 	fmt.Printf("%v", result)
// 	// Output: map[1:one 2:two 3:three]
// }

func ExampleKVSeq_ForEach() {
	iterKVPairs[int, string](1, "one", 2, "two", 3, "three").
		ForEach(func(k int, v string) {
			fmt.Printf("%v: %v\n", k, v)
		})
	// Output: 1: one
	// 2: two
	// 3: three
}

func ExampleKVSeq_Map() {
	seq := iterKVPairs[int, string](1, "one", 2, "two", 3, "three").
		Map(func(k int, v string) (int, string) {
			return k * 2, v + "+" + v
		})
	result := maps.Collect(iter.Seq2[int, string](seq))
	fmt.Printf("%v", result)
	// Output: map[2:one+one 4:two+two 6:three+three]
}

func ExampleKVSeq_Take() {
	seq := loz.IterSlice([]string{"zero", "one", "two", "three", "four"}).
		Indexed().
		Take(2)
	result := toMap(seq)
	fmt.Printf("%v", result)
	// Output: map[0:zero 1:one]
}

func ExampleKVSeq_TakeWhile() {
	seq := loz.IterSlice([]string{"zero", "one", "two", "three", "four"}).
		Indexed().
		TakeWhile(func(k int, v string) bool {
			return k < 3
		})
	result := toMap(seq)
	fmt.Printf("%v", result)
	// Output: map[0:zero 1:one 2:two]
}

func ExampleKVSeq_Skip() {
	seq := loz.IterSlice([]string{"zero", "one", "two", "three", "four"}).
		Indexed().
		Skip(3)
	result := toMap(seq)
	fmt.Printf("%v", result)
	// Output: map[3:three 4:four]
}

func ExampleKVSeq_SkipWhile() {
	seq := loz.IterSlice([]string{"zero", "one", "two", "three", "four"}).
		Indexed().
		SkipWhile(func(k int, v string) bool {
			return k < 3
		})
	result := toMap(seq)
	fmt.Printf("%v", result)
	// Output: map[3:three 4:four]
}

func ExampleKVSeq_Filter() {
	seq := loz.IterSlice([]string{"zero", "one", "two", "three", "four"}).
		Indexed().
		Filter(func(k int, v string) bool {
			return k%2 != 0 || len(v) == 3
		})
	result := toMap(seq)
	fmt.Printf("%v", result)
	// Output: map[1:one 2:two 3:three]
}

func ExampleKVSeq_Any() {
	seq := loz.IterMap(map[string]string{
		"greeting": "Hello there!",
		"response": "General Kenobi!",
	})
	valsHaveHello := seq.
		Any(func(k, v string) bool {
			return strings.HasPrefix(v, "Hello")
		})
	keysAreShort := seq.Any(func(k, v string) bool {
		return len(k) < 4
	})
	fmt.Printf("%v, %v", valsHaveHello, keysAreShort)
	// Output: true, false
}

func ExampleKVSeq_None() {
	seq := loz.IterMap(map[string]string{
		"greeting": "Hello there!",
		"response": "General Kenobi!",
		"followUp": "You are a bold one.",
	})
	noBoredValues := seq.
		None(func(k, v string) bool {
			return v[len(v)-1] != '!'
		})
	noUpperCaseKeys := seq.None(func(k, v string) bool {
		return strings.ToLower(k)[0] != k[0]
	})
	fmt.Printf("%v, %v", noBoredValues, noUpperCaseKeys)
	// Output: false, true
}

func ExampleKVSeq_Every() {
	seq := loz.IterMap(map[string]string{
		"greeting": "Hello there!",
		"response": "General Kenobi!",
		"followUp": "You are a bold one.",
	})
	allExcitedValues := seq.
		Every(func(k, v string) bool {
			return v[len(v)-1] == '!'
		})
	allLowercaseKeys := seq.Every(func(k, v string) bool {
		return strings.ToLower(k)[0] == k[0]
	})
	fmt.Printf("%v, %v", allExcitedValues, allLowercaseKeys)
	// Output: false, true
}

func ExampleKVSeq_First() {
	k, v, err := loz.IterMap(map[int]bool{}).First()
	fmt.Printf("%v, %v, %v\n", k, v, err == nil)
	k, v, err = iterKVPairs[int, bool](1, true, 2, false, 3, true).First()
	fmt.Printf("%v, %v, %v", k, v, err == nil)
	// Output: 0, false, false
	// 1, true, true
}

func ExampleKVSeq_Last() {
	k, v, err := loz.IterMap(map[int]bool{}).Last()
	fmt.Printf("%v, %v, %v\n", k, v, err != nil)
	k, v, err = iterKVPairs[int, bool](1, true, 2, false, 3, true).Last()
	fmt.Printf("%v, %v, %v", k, v, err)
	// Output: 0, false, true
	// 3, true, <nil>
}

func ExampleKVSeq_Fold() {
	addKMultV := func(k1, v1, k2, v2 int) (int, int) {
		return k1 + k2, v1 * v2
	}
	foldEmptyKey, foldEmptyVal := loz.IterSlice([]int{}).Indexed().Fold(100, 42, addKMultV)
	fmt.Printf("%v, %v\n", foldEmptyKey, foldEmptyVal)
	foldKey, foldVal := loz.IterSlice([]int{2, 4, 2}).Indexed().Fold(0, 1, addKMultV)
	fmt.Printf("%v, %v", foldKey, foldVal)
	// Output: 100, 42
	// 3, 16
}

func ExampleKVSeq_Reduce() {
	addKMultV := func(k1, v1, k2, v2 int) (int, int) {
		return k1 + k2, v1 * v2
	}
	reducedK, reducedV, err := loz.IterSlice([]int{}).Indexed().Reduce(addKMultV)
	fmt.Printf("%v, %v, %v\n", reducedK, reducedV, err != nil)
	reducedK, reducedV, err = loz.IterSlice([]int{2, 4, 2}).Indexed().Reduce(addKMultV)
	fmt.Printf("%v, %v, %v", reducedK, reducedV, err)
	// Output: 0, 0, true
	// 3, 16, <nil>
}

func ExampleKVSeq_FilterMap() {
	matching := loz.IterSlice([]int{8, 1, 5, 3}).
		Indexed().
		FilterMap(func(i1, i2 int) (int, int, error) {
			if i1 != i2 {
				return 0, 0, errors.New("Index and val not equal")
			}
			return i1, i2, nil
		}).
		Values().
		CollectSlice()
	fmt.Print(matching)
	// Output: [1 3]
}

func TestKVSeqTryMethods(t *testing.T) {
	seq := loz.IterMap(map[int]string{1: "one", 2: "two", 3: "three"})
	haltingErr := errors.New("Testing error")
	haltingSeq := seq.Map(func(i int, s string) (int, string) {
		loz.PanicHaltIteration(haltingErr)
		return i, s
	})
	table := []struct {
		name string
		run  func(loz.KVSeq[int, string]) error
	}{
		{
			"TryForEach",
			func(s loz.KVSeq[int, string]) error {
				return s.TryForEach(func(i int, s string) {})
			},
		},
		{
			"TryReduce",
			func(s loz.KVSeq[int, string]) error {
				_, _, err := s.TryReduce(func(i1 int, s1 string, i2 int, s2 string) (int, string) {
					return i1 + i2, s1 + s2
				})
				return err
			},
		},
		{
			"TryFold",
			func(s loz.KVSeq[int, string]) error {
				_, _, err := s.TryFold(0, "", func(i1 int, s1 string, i2 int, s2 string) (int, string) {
					return i1 + i2, s1 + s2
				})
				return err
			},
		},
		{
			"TryFirst",
			func(s loz.KVSeq[int, string]) error {
				_, _, err := s.TryFirst()
				return err
			},
		},
		{
			"TryLast",
			func(s loz.KVSeq[int, string]) error {
				_, _, err := s.TryLast()
				return err
			},
		},
		{
			"TryAny",
			func(s loz.KVSeq[int, string]) error {
				_, err := s.TryAny(func(i int, s string) bool {
					return i%2 == 0
				})
				return err
			},
		},
		{
			"TryNone",
			func(s loz.KVSeq[int, string]) error {
				_, err := s.TryNone(func(i int, s string) bool {
					return i%2 == 0
				})
				return err
			},
		},
		{
			"TryEvery",
			func(s loz.KVSeq[int, string]) error {
				_, err := s.TryEvery(func(i int, s string) bool {
					return i%2 == 0
				})
				return err
			},
		},
	}

	for _, row := range table {
		t.Run(row.name, func(t *testing.T) {
			err := row.run(seq)
			assert.Nil(t, err)
			err = row.run(haltingSeq)
			assert.Equal(t, err, haltingErr)
		})
	}
}
