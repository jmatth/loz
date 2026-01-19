package mapping_test

import (
	"fmt"
	"maps"
	"reflect"
	"slices"
	"testing"

	"github.com/jmatth/loz"
	lom "github.com/jmatth/loz/mapping"
	"github.com/stretchr/testify/assert"
)

func TestMapHasAllSeqMethods(t *testing.T) {
	seq := loz.IterSlice([]string{})
	m := lom.Map1[string, int](seq)
	seqType := reflect.TypeOf(seq)
	mapType := reflect.TypeOf(m)
	blocked := []string{
		"Any",
		"TryAny",
		"None",
		"TryNone",
		"Every",
		"TryEvery",
		"First",
		"TryFirst",
		"Last",
		"TryLast",
		"ForEach",
		"TryForEach",
		"CollectSlice",
		"TryCollectSlice",
		"AppendSlice",
		"Reduce",
		"TryReduce",
		"Indexed",
	}
	for i := range seqType.NumMethod() {
		seqMethod := seqType.Method(i)
		if slices.Contains(blocked, seqMethod.Name) {
			continue
		}
		_, ok := mapType.MethodByName(seqMethod.Name)
		assert.Truef(t, ok, "Method %s should exist on Map types", seqMethod.Name)
	}
}

func TestMap(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	mapper := lom.Map1[int, string](loz.IterSlice(nums))
	mapped := mapper.Map(func(n int) string { return fmt.Sprintf("%v", n) })
	assert.Equal(t, []string{"1", "2", "3", "4", "5"}, mapped.CollectSlice())
}

func TestMultiMap(t *testing.T) {
	nums := []string{"1", "200", "3", "42", "55"}
	mapper := lom.Map3[string, byte, int, float64](loz.IterSlice(nums))
	mapped := mapper.
		Map(func(s string) byte { return s[0] }).
		Map(func(b byte) int { return int(b - 0x30) }).
		Map(func(n int) float64 { return float64(n) * 11 / 10 }).
		CollectSlice()
	assert.Equal(t, []float64{1.1, 2.2, 3.3, 4.4, 5.5}, mapped)
}

func TestKVMap1(t *testing.T) {
	m := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
	}
	iterator := lom.KVMap1[
		int, string,
		string, byte](maps.All(m)).
		Map(func(k int, v string) (string, byte) {
			return fmt.Sprintf("%v", k), v[0]
		})
	assert.ElementsMatch(t, iterator.Keys().CollectSlice(), []string{"1", "2", "3", "4", "5"})
	assert.ElementsMatch(t, iterator.Values().CollectSlice(), []byte{'o', 't', 't', 'f', 'f'})
}
