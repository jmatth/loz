package loz_test

import (
	"fmt"
	"maps"
	"testing"

	"github.com/jmatth/loz"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	mapper := loz.Map1[int, string](loz.IterSlice(nums))
	mapped := mapper.Map(func(n int) string { return fmt.Sprintf("%v", n) })
	assert.Equal(t, []string{"1", "2", "3", "4", "5"}, mapped.CollectSlice())
}

func TestMultiMap(t *testing.T) {
	nums := []string{"1", "200", "3", "42", "55"}
	mapper := loz.Map3[string, byte, int, float64](loz.IterSlice(nums))
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
	iterator := loz.KVMap1[
		int, string,
		string, byte](maps.All(m)).
		Map(func(k int, v string) (string, byte) {
			return fmt.Sprintf("%v", k), v[0]
		})
	assert.ElementsMatch(t, iterator.Keys().CollectSlice(), []string{"1", "2", "3", "4", "5"})
	assert.ElementsMatch(t, iterator.Values().CollectSlice(), []byte{'o', 't', 't', 'f', 'f'})
}

func ExampleMap1_Fold() {
	result := loz.Map1[int, string](loz.RangeFrom(1, 6)).
		Fold("", func(acc string, n int) string {
			if acc == "" {
				return fmt.Sprintf("%d", n)
			}
			return fmt.Sprintf("%s, %d", acc, n)
		})
	fmt.Printf("%v", result)
	// Output: 1, 2, 3, 4, 5
}
