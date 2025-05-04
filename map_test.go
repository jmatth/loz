package loz

import (
	"fmt"
	"maps"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	mapper := Mapper1[int, string](IterSlice(nums))
	mapped := mapper.Map(func(n int) string { return fmt.Sprintf("%v", n) })
	assert.Equal(t, []string{"1", "2", "3", "4", "5"}, Seq[string](mapped).ToSlice())
}

func TestMultiMap(t *testing.T) {
	nums := []string{"1", "200", "3", "42", "55"}
	mapper := Mapper3[string, byte, int, float64](IterSlice(nums))
	mapped := mapper.
		Map(func(s string) byte { return s[0] }).
		Map(func(b byte) int { return int(b - 0x30) }).
		Map(func(n int) float64 { return float64(n) * 11 / 10 }).
		ToSlice()
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
	iterator := KVMapper1[
		int, string,
		string, byte](maps.All(m)).
		Map(func(k int, v string) (string, byte) {
			return fmt.Sprintf("%v", k), v[0]
		})
	assert.ElementsMatch(t, iterator.Keys().ToSlice(), []string{"1", "2", "3", "4", "5"})
	assert.ElementsMatch(t, iterator.Values().ToSlice(), []byte{'o', 't', 't', 'f', 'f'})
}
